package core_telegram

import (
	"encoding/json"
	"fmt"
)

func InitOutputMessage(
	chatId int64,
	text string,
	replyKeyboard *[][]string,
	inlineKeyboard *[][]InlineKeyboardButton,
) *OutputMessage {
	message := NewMessageBuilder().
		ChatId(chatId).
		Text(text, false)

	if replyKeyboard != nil && len(*replyKeyboard) > 0 {
		message.ReplyKeyboard(replyKeyboard)
	}

	if inlineKeyboard != nil && len(*inlineKeyboard) > 0 {
		message.InlineKeyboard(inlineKeyboard)
	}

	return message.Build()
}

func InitEditMessage(
	chatId int64,
	messageId int64,
	text string,
) *EditMessage {
	return &EditMessage{
		ChatId:    chatId,
		MessageId: messageId,
		Text:      text,
	}
}

func MapEditMessageToJson(message *EditMessage) string {
	result, err := json.Marshal(message)
	if err != nil {
		fmt.Println("MapMessageToJson")
		fmt.Println(err)
		return ""
	}
	return string(result)
}

func MapJsonToEditMessage(body string) *EditMessage {
	var message EditMessage
	err := json.Unmarshal([]byte(body), &message)
	if err != nil {
		fmt.Println("MapJsonToMessage")
		fmt.Println(err)
		return nil
	}
	return &message
}

func MapEditKeyboardMessageToJson(message *EditKeyboard) string {
	result, err := json.Marshal(message)
	if err != nil {
		fmt.Println("MapMessageToJson")
		fmt.Println(err)
		return ""
	}
	return string(result)
}

func MapJsonToEditKeyboardMessage(body string) *EditKeyboard {
	var message EditKeyboard
	err := json.Unmarshal([]byte(body), &message)
	if err != nil {
		fmt.Println("MapJsonToMessage")
		fmt.Println(err)
		return nil
	}
	return &message
}

func MapCallbackMessageToJson(message *TelegramCallback) string {
	result, err := json.Marshal(message)
	if err != nil {
		fmt.Println("MapMessageToJson")
		fmt.Println(err)
		return ""
	}
	return string(result)
}

func MapJsonToCallbackMessage(body string) *TelegramCallback {
	var message TelegramCallback
	err := json.Unmarshal([]byte(body), &message)
	if err != nil {
		fmt.Println("MapJsonToMessage")
		fmt.Println(err)
		return nil
	}
	return &message
}

func MapInputMessageToJson(message *TelegramMessage) string {
	result, err := json.Marshal(message)
	if err != nil {
		fmt.Println("MapMessageToJson")
		fmt.Println(err)
		return ""
	}
	return string(result)
}

func MapJsonToInputMessage(body string) *TelegramMessage {
	var message TelegramMessage
	err := json.Unmarshal([]byte(body), &message)
	if err != nil {
		fmt.Println("MapJsonToMessage")
		fmt.Println(err)
		return nil
	}
	return &message
}

func MapOutputMessageToJson(message *OutputMessage) string {
	result, err := json.Marshal(message)
	if err != nil {
		fmt.Println("MapMessageToJson")
		fmt.Println(err)
		return ""
	}
	return string(result)
}

func MapJsonToOutputMessage(body string) *OutputMessage {
	var message OutputMessage
	err := json.Unmarshal([]byte(body), &message)
	if err != nil {
		fmt.Println("MapJsonToMessage")
		fmt.Println(err)
		return nil
	}
	return &message
}
