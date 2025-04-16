package tg_controller

import (
	"fmt"
	"ollama-bot/configs"
	"ollama-bot/internal/ollama_api"

	"ollama-bot/internal/telegram_bot/event_handler"
	"ollama-bot/internal/telegram_bot/telegram_redis"
	"ollama-bot/internal/telegram_logger"
	"ollama-bot/pkg/core/redis"
	"ollama-bot/pkg/core_telegram"
	"ollama-bot/pkg/logger"
	"time"
)

type TelegramControllerApi interface {
	HandleBotMessages()
	HandleAiMessages()
	SendMessages(redisKey string, delay time.Duration)
	ListenUpdates()
}

type TelegramControllerImpl struct {
	telegramRedis  telegram_redis.TelegramRedisApi
	tgNetwork      core_telegram.TelegramNetworkApi
	eventHandler   event_handler.EventApi
	telegramLogger telegram_logger.LoggerApi
	ollamaApi      ollama_api.OllamaApi
}

func New(
	props *configs.BotProps,
	coreRedis redis.Client,
	ollamaApi ollama_api.OllamaApi,
	logger *logger.Logger,
) TelegramControllerApi {
	if props == nil {
		panic("BOT PROPS ARE NIL")
	}
	var telegramRedis = telegram_redis.New(coreRedis)
	var telegramNetwork = core_telegram.New(logger, props.Token)
	var telegramLogger = telegram_logger.New(props, telegramRedis)
	var eventHandler = event_handler.New(telegramRedis)
	return &TelegramControllerImpl{
		telegramRedis:  telegramRedis,
		tgNetwork:      telegramNetwork,
		telegramLogger: telegramLogger,
		eventHandler:   eventHandler,
		ollamaApi:      ollamaApi,
	}
}

func (controller *TelegramControllerImpl) ListenUpdates() {
	for {
		offset := controller.telegramRedis.GetOffset()
		telegramResponse, err := controller.tgNetwork.GetUpdate(offset)

		if err != nil {
			fmt.Println("GET UPDATES ERROR")
			fmt.Println(err)
			continue
		}

		for _, item := range telegramResponse.Result {
			if item.Message != nil {
				controller.telegramRedis.RPushInputMessage(item.Message)
			}
			if item.Callback != nil {
				controller.telegramRedis.RPushGroupCallbackMessage(item.Callback)
			}
			controller.telegramRedis.SaveOffset(item.UpdateId + 1)
		}
		time.Sleep(500 * time.Millisecond)
	}
}

func (controller *TelegramControllerImpl) SendMessages(redisKey string, delay time.Duration) {
	for {
		time.Sleep(delay)
		message := controller.telegramRedis.PopOutputMessage(redisKey)

		if message != nil && message.ChatId != 0 {
			result, err := controller.tgNetwork.SendMessage(message)
			if err != nil {
				fmt.Println("SEND MESSAGES ERROR " + err.Error())
				controller.telegramRedis.LPushOutputMessage(redisKey, message)
				time.Sleep(time.Second * 5)
				continue
			}
			if result != nil && result.Ok {
				controller.eventHandler.HandleOutputMessage(result.Result.Chat.Id, result.Result.MessageId)
			}

		} else {
			time.Sleep(time.Millisecond * 500)
		}

		editMessage := controller.telegramRedis.PopEditMessage()
		if editMessage != nil && editMessage.ChatId != 0 {
			fmt.Println(editMessage)
			result := controller.tgNetwork.ProcessCallback(editMessage)

			switch result.StatusCode {
			case 429:
				controller.telegramRedis.RPushEditMessage(editMessage)
				time.Sleep(time.Second * 5)
				continue
			}
		}

		editKeyboard := controller.telegramRedis.PopEditKeyboardMessage()
		if editKeyboard != nil && editKeyboard.ChatId != 0 {
			result := controller.tgNetwork.ProcessCallback(editKeyboard)
			switch result.StatusCode {
			case 429:
				controller.telegramRedis.RPushEditKeyboard(editKeyboard)
				time.Sleep(time.Second * 5)
				continue
			}
		}
	}
}

func (controller *TelegramControllerImpl) HandleBotMessages() {
	for {
		message := controller.telegramRedis.PopInputMessage()
		if message != nil && message.From.Id != 0 {
			controller.eventHandler.HandleInputMessage(message.Chat.Id, message.Text)
		} else {
			time.Sleep(time.Millisecond * 500)
		}
	}
}

func (controller *TelegramControllerImpl) HandleAiMessages() {
	for {
		var promptData = controller.telegramRedis.LPopAiRequest()
		if promptData == nil || promptData.Text == "" {
			time.Sleep(time.Millisecond * 500)
		} else {
			fullMsg := ""
			var bodyCh = make(chan ollama_api.Answer)
			var errCh = make(chan error)
			var ticker = time.NewTicker(time.Second)
			var user = controller.telegramRedis.GetUserData(promptData.ChatId)
			go controller.ollamaApi.GetResponse(promptData.Text, bodyCh, errCh, user.Context)
		loop:
			for {
				select {
				case answer, ok := <-bodyCh:
					fullMsg += answer.Response
					if answer.Done {
						controller.telegramRedis.SaveContext(promptData.ChatId, answer.Context)
					}
					if !ok {
						ticker.Stop()
						promptData.Text = fullMsg
						controller.telegramRedis.RPushEditMessage(promptData)
						break loop
					}

				case err := <-errCh:
					if err != nil {
						controller.telegramLogger.EnqueueMsg(err.Error())
						break loop
					}
				case <-ticker.C:
					promptData.Text = fullMsg
					controller.telegramRedis.RPushEditMessage(promptData)
				}
			}
		}
	}
}
