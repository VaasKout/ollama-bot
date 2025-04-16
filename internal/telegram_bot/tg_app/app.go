package tg_app

import (
	"ollama-bot/configs"
	"ollama-bot/internal/ollama_api"
	"ollama-bot/internal/telegram_bot/controller"
	"ollama-bot/internal/telegram_bot/telegram_redis"
	"ollama-bot/internal/telegram_logger"
	"ollama-bot/pkg/core/redis"
	"ollama-bot/pkg/logger"
	"time"
)

func Run() {
	config := configs.New()
	logApi := logger.New(logger.INFO, true)
	coreRedis := redis.New(config)

	ollamaApi := ollama_api.New(config.ModelProps.Model)
	tgBot := tg_controller.New(config.BotProps, coreRedis, ollamaApi, logApi)

	go tgBot.HandleBotMessages()
	go tgBot.HandleAiMessages()
	go tgBot.SendMessages(telegram_redis.USER_OUTPUT_MESSAGES, time.Millisecond*500)
	go tgBot.SendMessages(telegram_logger.BOT_LOGS_KEY, time.Second*1)
	tgBot.ListenUpdates()
}
