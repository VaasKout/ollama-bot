package telegram_redis

import (
	"fmt"
	"ollama-bot/pkg/core_telegram"
)

const USER_KEY = "user"

type UserApi interface {
	IsUser(userName string) bool

	SaveUserData(user *UserData) bool
	GetUser(inputMessage *core_telegram.TelegramMessage) *UserData
	GetUserData(userId int64) *UserData
	ClearUserData(userId int64) bool
	SaveContext(userId int64, context []int)
}

func (adapter *TelegramRedisImpl) IsUser(userName string) bool {
	return adapter.redis.SISMembers(USERS_KEY, userName)
}

func (adapter *TelegramRedisImpl) SaveUserData(userData *UserData) bool {
	var userJson = MapUserDataToJson(userData)
	err := adapter.redis.SetData(fmt.Sprint(userData.ChatId), userJson)
	return err == nil
}

func (adapter *TelegramRedisImpl) ClearUserData(userId int64) bool {
	err := adapter.redis.DeleteData(fmt.Sprint(userId))
	return err == nil
}

func (adapter *TelegramRedisImpl) GetUser(inputMessage *core_telegram.TelegramMessage) *UserData {
	fmt.Println(inputMessage)
	if inputMessage == nil || inputMessage.Chat.Id == 0 {
		return nil
	}
	role := adapter.getRole(inputMessage.From.UserName)
	if len(role) > 0 {
		result := adapter.redis.GetData(fmt.Sprint(inputMessage.From.Id))
		user := MapJsonToUserData(result)
		if user != nil && len(user.UserName) > 0 && user.Role == role {
			return user
		}
		var newUser = &UserData{
			ChatId:       inputMessage.From.Id,
			Role:         role,
			UserName:     inputMessage.From.UserName,
			CurrentState: WRITING_TO_AI_STATE,
		}
		adapter.SaveUserData(newUser)
		return newUser
	}
	adapter.ClearUserData(inputMessage.From.Id)
	return nil
}

func (adapter *TelegramRedisImpl) getRole(userName string) string {
	switch {
	case adapter.IsUser(userName):
		return USER_KEY
	default:
		return ""
	}
}

func (adapter *TelegramRedisImpl) GetUserData(userId int64) *UserData {
	result := adapter.redis.GetData(fmt.Sprint(userId))
	userData := MapJsonToUserData(result)
	return userData
}

func (adapter *TelegramRedisImpl) SaveContext(userId int64, context []int) {
	result := adapter.redis.GetData(fmt.Sprint(userId))
	userData := MapJsonToUserData(result)
	userData.Context = context
	adapter.SaveUserData(userData)
}
