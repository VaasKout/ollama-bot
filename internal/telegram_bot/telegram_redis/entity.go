package telegram_redis

import (
	"encoding/json"
	"fmt"
)

const (
	WRITING_TO_AI_STATE       = "WRITING_TO_AI_STATE"
	AWAITING_RESPONSE_FROM_AI = "AWAITING_RESPONSE_FROM_AI"
)

type UserData struct {
	UserName      string `json:"user_name"`
	ChatId        int64  `json:"chat_id"`
	Role          string `json:"role"`
	LastInputText string `json:"last_input_text"`
	LastMessageID int64  `json:"last_message_id"`
	CurrentState  string `json:"current_state"`
	Context       []int  `json:"context"`
}

func (userData *UserData) UpdateUserState(newState string) {
	if userData != nil {
		userData.CurrentState = newState
	}
}

func MapUserDataToJson(userData *UserData) string {
	result, err := json.Marshal(userData)
	if err != nil {
		fmt.Println("MapUserDataToJson")
		fmt.Println(err)
		return ""
	}
	return string(result)
}

func MapJsonToUserData(body string) *UserData {
	var userData UserData
	err := json.Unmarshal([]byte(body), &userData)
	if err != nil {
		fmt.Println("MapJsonToUserData")
		fmt.Println(err)
		return nil
	}
	return &userData
}
