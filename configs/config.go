package configs

import (
	"os"
	"strconv"
)

type Config struct {
	ModelProps *ModelProps
	RedisProps *RedisProps
	BotProps   *BotProps
}

type ModelProps struct {
	Model string
}

type RedisProps struct {
	RedisUser     string
	RedisPassword string
	RedisAddress  string
}

type BotProps struct {
	Token        string
	BotLogChatId int64
}

func New() *Config {
	redisUser := os.Getenv("REDIS_USER")
	redisPassword := os.Getenv("REDIS_PASSWORD")
	redisAddress := os.Getenv("REDIS_ADDRESS")

	botToken := os.Getenv("BOT_KEY")
	if botToken == "" {
		panic("Bot key is not found")
	}

	botLogsChatId, err := strconv.ParseInt(os.Getenv("BOT_LOGS_CHAT_ID"), 10, 64)
	if err != nil || botLogsChatId == 0 {
		panic(err)
	}

	model := os.Getenv("MODEL")
	if model == "" {
		panic("Model property is not found")
	}

	return &Config{
		RedisProps: &RedisProps{
			RedisUser:     redisUser,
			RedisPassword: redisPassword,
			RedisAddress:  redisAddress,
		},
		BotProps: &BotProps{
			Token:        botToken,
			BotLogChatId: botLogsChatId,
		},
		ModelProps: &ModelProps{
			Model: model,
		},
	}
}
