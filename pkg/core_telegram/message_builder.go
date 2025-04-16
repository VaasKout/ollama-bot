package core_telegram

import "html"

type OutputMessage struct {
	ChatId            int64     `json:"chat_id"`
	Text              string    `json:"text,omitempty"`
	PhotoId           string    `json:"photo,omitempty"`
	DocumentId        string    `json:"document,omitempty"`
	AudioId           string    `json:"audio,omitempty"`
	VideoId           string    `json:"video,omitempty"`
	StickerId         string    `json:"sticker,omitempty"`
	VoiceId           string    `json:"voice,omitempty"`
	Caption           string    `json:"caption,omitempty"`
	DisableWebPreview bool      `json:"disable_web_page_preview,omitempty"`
	ParseMode         string    `json:"parse_mode,omitempty"`
	Keyboard          IKeyboard `json:"reply_markup,omitempty"`
}

type IMessageBuilder interface {
	ChatId(chatId int64) IMessageBuilder
	Text(text string, needParse bool) IMessageBuilder
	PhotoId(photoId string) IMessageBuilder
	DocumentId(documentId string) IMessageBuilder
	VideoId(videoId string) IMessageBuilder
	AudioId(videoId string) IMessageBuilder
	StickerId(videoId string) IMessageBuilder
	VoiceId(voiceId string) IMessageBuilder
	Caption(caption string, needParse bool) IMessageBuilder
	DisableWebPreview(disableWebPreview bool) IMessageBuilder
	ParseMode(parseMode string) IMessageBuilder
	ReplyKeyboard(buttonsText *[][]string) IMessageBuilder
	InlineKeyboard(buttons *[][]InlineKeyboardButton) IMessageBuilder
	Build() *OutputMessage
}

type MessageBuilder struct {
	chatId            int64
	text              string
	photoId           string
	documentId        string
	audioId           string
	videoId           string
	stickerId         string
	voiceId           string
	caption           string
	disableWebPreview bool
	parseMode         string
	keyboard          IKeyboard
}

func NewMessageBuilder() *MessageBuilder {
	return &MessageBuilder{
		disableWebPreview: true,
		parseMode:         HtmlParseMode,
	}
}

func (builder *MessageBuilder) ChatId(chatId int64) IMessageBuilder {
	builder.chatId = chatId
	return builder
}

func (builder *MessageBuilder) Text(text string, needParse bool) IMessageBuilder {
	if needParse {
		builder.text = html.EscapeString(text)
	} else {
		builder.text = text
	}

	return builder
}

func (builder *MessageBuilder) PhotoId(photoId string) IMessageBuilder {
	builder.photoId = photoId
	return builder
}

func (builder *MessageBuilder) DocumentId(documentId string) IMessageBuilder {
	builder.documentId = documentId
	return builder
}

func (builder *MessageBuilder) AudioId(audioId string) IMessageBuilder {
	builder.audioId = audioId
	return builder
}

func (builder *MessageBuilder) VideoId(videoId string) IMessageBuilder {
	builder.videoId = videoId
	return builder
}

func (builder *MessageBuilder) StickerId(stickerId string) IMessageBuilder {
	builder.stickerId = stickerId
	return builder
}

func (builder *MessageBuilder) VoiceId(voiceId string) IMessageBuilder {
	builder.voiceId = voiceId
	return builder
}

func (builder *MessageBuilder) Caption(caption string, needParse bool) IMessageBuilder {
	if needParse {
		builder.caption = html.EscapeString(caption)
	} else {
		builder.caption = caption
	}
	return builder
}

func (builder *MessageBuilder) DisableWebPreview(disableWebPreview bool) IMessageBuilder {
	builder.disableWebPreview = disableWebPreview
	return builder
}

func (builder *MessageBuilder) ParseMode(parseMode string) IMessageBuilder {
	builder.parseMode = parseMode
	return builder
}

func (builder *MessageBuilder) ReplyKeyboard(buttonsText *[][]string) IMessageBuilder {
	keyboard := NewReplyKeyboardMarkup(buttonsText)

	if keyboard == nil {
		return builder
	}

	builder.keyboard = keyboard
	return builder
}

func (builder *MessageBuilder) InlineKeyboard(buttons *[][]InlineKeyboardButton) IMessageBuilder {
	keyboard := NewInlineKeyboardMarkup(buttons)

	if keyboard == nil {
		return builder
	}

	builder.keyboard = keyboard
	return builder
}

func (builder *MessageBuilder) Build() *OutputMessage {
	return &OutputMessage{
		ChatId:            builder.chatId,
		Text:              builder.text,
		PhotoId:           builder.photoId,
		DocumentId:        builder.documentId,
		AudioId:           builder.audioId,
		VideoId:           builder.videoId,
		VoiceId:           builder.voiceId,
		StickerId:         builder.stickerId,
		Caption:           builder.caption,
		DisableWebPreview: builder.disableWebPreview,
		ParseMode:         builder.parseMode,
		Keyboard:          builder.keyboard,
	}
}
