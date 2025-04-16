package event_handler

import (
	"ollama-bot/internal/telegram_bot/telegram_redis"
	"ollama-bot/pkg/core_telegram"
)

var InputMessageMap = map[string]func(e *EventImpl){
	"/start": func(e *EventImpl) {
		var outputText = "Welcome to the DeepSeek Bot"

		e.user.UpdateUserState(telegram_redis.WRITING_TO_AI_STATE)
		e.tgRedis.SaveUserData(e.user)

		var outputMessage = core_telegram.InitOutputMessage(e.user.ChatId, outputText, nil, nil)
		e.tgRedis.RPushOutputMessage(telegram_redis.USER_OUTPUT_MESSAGES, outputMessage)
	},
}

var InputStateMap = map[string]func(e *EventImpl){
	telegram_redis.WRITING_TO_AI_STATE:       writeToAi,
	telegram_redis.AWAITING_RESPONSE_FROM_AI: writeToAi,
}

var writeToAi = func(e *EventImpl) {
	e.user.UpdateUserState(telegram_redis.AWAITING_RESPONSE_FROM_AI)
	e.tgRedis.SaveUserData(e.user)
	var outputText = "Loading..."
	var outputMessage = core_telegram.InitOutputMessage(e.user.ChatId, outputText, nil, nil)
	e.tgRedis.RPushOutputMessage(telegram_redis.USER_OUTPUT_MESSAGES, outputMessage)
}

var OutputStateMap = map[string]func(e *EventImpl){
	telegram_redis.AWAITING_RESPONSE_FROM_AI: func(e *EventImpl) {
		e.tgRedis.EnqueueAiRequest(e.user.LastInputText, e.user.ChatId, e.user.LastMessageID)
	},
}
