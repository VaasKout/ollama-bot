package core_telegram

const (
	SendMessageMethod  = "sendMessage"
	SendPhotoMethod    = "sendPhoto"
	SendDocumentMethod = "sendDocument"
	SendVideoMethod    = "sendVideo"
	SendAudioMethod    = "sendAudio"
	SendStickerMethod  = "sendSticker"
	SendVoiceMethod    = "sendVoice"

	EditMessageTextMethod = "editMessageText"
	EditMessageKeyboard   = "editMessageReplyMarkup"

	VoiceMessagesForbiddenError = "Bad Request: VOICE_MESSAGES_FORBIDDEN"

	VoiceMessagesForbiddenMessage = "Вам пытались отправить голосовое сообщение, но стоит запрет на отправку"
)

const (
	HtmlParseMode = "HTML"
)
