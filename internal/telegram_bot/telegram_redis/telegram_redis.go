package telegram_redis

import (
	"fmt"
	"ollama-bot/pkg/core/redis"
	"ollama-bot/pkg/core_telegram"
	"strconv"
)

const USERS_KEY = "users"

const (
	USER_OUTPUT_MESSAGES     = "user_output_messages"
	USER_INPUT_MESSAGES      = "user_input_messages"
	GROUP_CALLBACK_MESSAGES  = "group_callback_messages"
	GROUP_EDIT_MESSAGES      = "group_edit_messages"
	DELETE_KEYBOARD_MESSAGES = "delete_keyboard_messages"
	TELEGRAM_OFFSET_KEY      = "telegram_offset"

	AI_QUEUE = "ai_queue"
)

type TelegramRedisApi interface {
	UserApi

	RPushInputMessage(message *core_telegram.TelegramMessage) bool
	RPushGroupCallbackMessage(message *core_telegram.TelegramCallback) bool
	RPushOutputMessage(redisKey string, message *core_telegram.OutputMessage) bool
	LPushOutputMessage(redisKey string, message *core_telegram.OutputMessage) bool
	RPushEditMessage(message *core_telegram.EditMessage) bool
	RPushEditKeyboard(message *core_telegram.EditKeyboard) bool

	PopCallbackMessage() *core_telegram.TelegramCallback
	PopInputMessage() *core_telegram.TelegramMessage
	PopOutputMessage(redisKey string) *core_telegram.OutputMessage
	PopEditMessage() *core_telegram.EditMessage
	PopEditKeyboardMessage() *core_telegram.EditKeyboard

	SaveOffset(value int64) bool
	GetOffset() int64

	EnqueueAiRequest(text string, chatId int64, messageId int64)
	LPopAiRequest() *core_telegram.EditMessage
}

type TelegramRedisImpl struct {
	redis redis.Client
}

func New(coreRedis redis.Client) TelegramRedisApi {
	return &TelegramRedisImpl{
		redis: coreRedis,
	}
}

func (adapter *TelegramRedisImpl) LPushOutputMessage(redisKey string, message *core_telegram.OutputMessage) bool {
	var msgJson = core_telegram.MapOutputMessageToJson(message)
	err := adapter.redis.LPush(redisKey, msgJson)
	return err == nil
}

func (adapter *TelegramRedisImpl) RPushOutputMessage(redisKey string, message *core_telegram.OutputMessage) bool {
	var msgJson = core_telegram.MapOutputMessageToJson(message)
	err := adapter.redis.RPush(redisKey, msgJson)
	return err == nil
}

func (adapter *TelegramRedisImpl) RPushEditMessage(message *core_telegram.EditMessage) bool {
	var msgJson = core_telegram.MapEditMessageToJson(message)
	err := adapter.redis.RPush(GROUP_EDIT_MESSAGES, msgJson)
	return err == nil
}

func (adapter *TelegramRedisImpl) RPushEditKeyboard(message *core_telegram.EditKeyboard) bool {
	var msgJson = core_telegram.MapEditKeyboardMessageToJson(message)
	err := adapter.redis.RPush(DELETE_KEYBOARD_MESSAGES, msgJson)
	return err == nil
}

func (adapter *TelegramRedisImpl) RPushInputMessage(message *core_telegram.TelegramMessage) bool {
	var msgJson = core_telegram.MapInputMessageToJson(message)
	err := adapter.redis.RPush(USER_INPUT_MESSAGES, msgJson)
	return err == nil
}

func (adapter *TelegramRedisImpl) RPushGroupCallbackMessage(message *core_telegram.TelegramCallback) bool {
	var msgJson = core_telegram.MapCallbackMessageToJson(message)
	err := adapter.redis.RPush(GROUP_CALLBACK_MESSAGES, msgJson)
	return err == nil
}

func (adapter *TelegramRedisImpl) PopCallbackMessage() *core_telegram.TelegramCallback {
	result := adapter.redis.LPop(GROUP_CALLBACK_MESSAGES)
	if len(result) == 0 {
		return nil
	}
	return core_telegram.MapJsonToCallbackMessage(result)
}

func (adapter *TelegramRedisImpl) PopEditMessage() *core_telegram.EditMessage {
	result := adapter.redis.LPop(GROUP_EDIT_MESSAGES)
	if len(result) == 0 {
		return nil
	}
	return core_telegram.MapJsonToEditMessage(result)
}

func (adapter *TelegramRedisImpl) PopEditKeyboardMessage() *core_telegram.EditKeyboard {
	result := adapter.redis.LPop(DELETE_KEYBOARD_MESSAGES)
	if len(result) == 0 {
		return nil
	}
	return core_telegram.MapJsonToEditKeyboardMessage(result)
}

func (adapter *TelegramRedisImpl) PopInputMessage() *core_telegram.TelegramMessage {
	result := adapter.redis.LPop(USER_INPUT_MESSAGES)
	if len(result) == 0 {
		return nil
	}
	return core_telegram.MapJsonToInputMessage(result)
}

func (adapter *TelegramRedisImpl) PopOutputMessage(redisKey string) *core_telegram.OutputMessage {
	result := adapter.redis.LPop(redisKey)
	if len(result) == 0 {
		return nil
	}
	return core_telegram.MapJsonToOutputMessage(result)
}

func (adapter *TelegramRedisImpl) SaveOffset(value int64) bool {
	err := adapter.redis.SetData(TELEGRAM_OFFSET_KEY, fmt.Sprint(value))
	return err == nil
}

func (adapter *TelegramRedisImpl) GetOffset() int64 {
	result := adapter.redis.GetData(TELEGRAM_OFFSET_KEY)
	i, err := strconv.ParseInt(result, 10, 64)
	if err != nil {
		return 0
	}
	return i
}

func (adapter *TelegramRedisImpl) EnqueueAiRequest(text string, chatId int64, messageId int64) {
	var editMsg = core_telegram.EditMessage{
		Text:      text,
		ChatId:    chatId,
		MessageId: messageId,
	}

	err := adapter.redis.RPush(AI_QUEUE, editMsg.ToJson())
	if err != nil {
		fmt.Println(err.Error())
	}
}

func (adapter *TelegramRedisImpl) LPopAiRequest() *core_telegram.EditMessage {
	var result = adapter.redis.LPop(AI_QUEUE)
	return core_telegram.EditMessageFromJson(result)
}
