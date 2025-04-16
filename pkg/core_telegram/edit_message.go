package core_telegram

import "encoding/json"

type CallbackOperations interface{}

type EditMessage struct {
	ChatId    int64  `json:"chat_id"`
	MessageId int64  `json:"message_id"`
	Text      string `json:"text"`
	CallbackOperations
}

func (e *EditMessage) ToJson() string {
	result, err := json.Marshal(e)
	if err != nil {
		return ""
	}
	return string(result)
}

func EditMessageFromJson(body string) *EditMessage {
	var result EditMessage
	err := json.Unmarshal([]byte(body), &result)
	if err != nil {
		return &EditMessage{}
	}
	return &result
}

type EditKeyboard struct {
	ChatId    int64 `json:"chat_id"`
	MessageId int64 `json:"message_id"`
	IKeyboard `json:"reply_markup,omitempty"`
	CallbackOperations
}
