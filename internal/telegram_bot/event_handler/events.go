package event_handler

import (
	"ollama-bot/internal/telegram_bot/telegram_redis"
)

type EventApi interface {
	HandleInputMessage(chatId int64, message string)
	HandleOutputMessage(chatId int64, messageId int64)
}

type EventImpl struct {
	user    *telegram_redis.UserData
	tgRedis telegram_redis.TelegramRedisApi
}

func (e *EventImpl) HandleInputMessage(chatId int64, message string) {
	var user = e.tgRedis.GetUserData(chatId)
	if user == nil || message == "" {
		return
	}
	e.user = user
	e.user.LastInputText = message

	e.tgRedis.SaveUserData(e.user)

	if action, ok := InputMessageMap[user.LastInputText]; ok {
		action(e)
		return
	}
	if action, ok := InputStateMap[user.CurrentState]; ok {
		action(e)
		return
	}
}

func (e *EventImpl) HandleOutputMessage(chatId int64, messageId int64) {
	var user = e.tgRedis.GetUserData(chatId)
	if user == nil || messageId == 0 {
		return
	}
	user.LastMessageID = messageId
	e.user = user
	e.tgRedis.SaveUserData(e.user)

	if action, ok := OutputStateMap[user.CurrentState]; ok {
		action(e)
		return
	}
}

func New(tgRedis telegram_redis.TelegramRedisApi) EventApi {
	return &EventImpl{
		tgRedis: tgRedis,
	}
}
