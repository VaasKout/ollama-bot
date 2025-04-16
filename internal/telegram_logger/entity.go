package telegram_logger

import (
	"fmt"
	"time"
)

const BOT_LOG_FORMAT = "<b>bot_answer:\n</b>%s"

func GetBotAnswerMessage(message string) string {
	if message == "" {
		return ""
	}
	outputMessage := fmt.Sprintf(BOT_LOG_FORMAT, message)
	messageWithDate := outputMessage + fmt.Sprintf("\n<b>date:</b> %s", getTimeDate())
	return messageWithDate
}

func getTimeDate() string {
	return time.Now().Format("15:04:05 02-01-06")
}
