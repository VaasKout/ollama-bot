package telegram_logger

import (
	"ollama-bot/configs"
	"ollama-bot/internal/telegram_bot/telegram_redis"
	"ollama-bot/pkg/core_telegram"
	"ollama-bot/pkg/logger"
)

const (
	BOT_LOGS_KEY = "bot_logs"
)

type LoggerApi interface {
	EnqueueMsg(message string)
}

type LoggerImpl struct {
	botLogsChatId int64
	telegramRedis telegram_redis.TelegramRedisApi
	logger        *logger.Logger
}

func New(
	props *configs.BotProps,
	telegramRedis telegram_redis.TelegramRedisApi,
) LoggerApi {
	if props == nil {
		panic("BOT PROPS ARE NIL")
	}
	return &LoggerImpl{
		botLogsChatId: props.BotLogChatId,
		telegramRedis: telegramRedis,
	}
}

func (tgLogger *LoggerImpl) EnqueueMsg(message string) {
	if message == "" {
		return
	}
	text := GetBotAnswerMessage(message)
	botMsg := core_telegram.InitOutputMessage(
		tgLogger.botLogsChatId,
		text,
		nil,
		nil,
	)
	tgLogger.telegramRedis.RPushOutputMessage(BOT_LOGS_KEY, botMsg)
}
