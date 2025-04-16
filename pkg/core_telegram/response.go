package core_telegram

//-------File Response-------

type TelegramFileResponse struct {
	Ok   bool         `json:"ok"`
	File TelegramFile `json:"result"`
}

type TelegramFile struct {
	FileId       string `json:"file_id"`
	FileUniqueId string `json:"file_unique_id"`
	FileSize     int64  `json:"file_size"`
	FilePath     string `json:"file_path"`
}

//-------Update Response------

type TelegramSendResponse struct {
	Ok     bool                       `json:"ok"`
	Result TelegramSendResultResponse `json:"result"`
}

type TelegramSendResultResponse struct {
	MessageId int64        `json:"message_id"`
	From      TelegramUser `json:"from"`
	Chat      TelegramChat `json:"chat"`
	Date      int64        `json:"date"`
	Text      string       `json:"text"`
}

type TelegramResponse struct {
	Ok     bool             `json:"ok"`
	Result []TelegramResult `json:"result"`
}

type TelegramResult struct {
	UpdateId int64             `json:"update_id"`
	Message  *TelegramMessage  `json:"message"`
	Callback *TelegramCallback `json:"callback_query"`
}

type TelegramCallback struct {
	Id           string                   `json:"id"`
	From         TelegramUser             `json:"from"`
	Message      *TelegramMessageCallback `json:"message"`
	ChatInstance string                   `json:"chat_instance"`
	Data         string                   `json:"data"`
	Keyboard     interface{}              `json:"reply_markup"`
}

type TelegramMessageCallback struct {
	Id       int64            `json:"message_id"`
	Date     int64            `json:"date"`
	Text     string           `json:"text"`
	Photo    []TelegramPhoto  `json:"photo"`
	Document TelegramDocument `json:"document"`
	Video    TelegramVideo    `json:"video"`
	Sticker  TelegramSticker  `json:"sticker"`
	Audio    TelegramAudio    `json:"audio"`
	Voice    TelegramVoice    `json:"voice"`
	Caption  string           `json:"caption"`
	Chat     TelegramChat     `json:"chat"`
	From     TelegramUser     `json:"from"`
}

type TelegramMessage struct {
	MessageId int64            `json:"message_id"`
	Date      int64            `json:"date"`
	Text      string           `json:"text"`
	Photo     []TelegramPhoto  `json:"photo"`
	Document  TelegramDocument `json:"document"`
	Video     TelegramVideo    `json:"video"`
	Sticker   TelegramSticker  `json:"sticker"`
	Audio     TelegramAudio    `json:"audio"`
	Voice     TelegramVoice    `json:"voice"`
	Caption   string           `json:"caption"`
	Chat      TelegramChat     `json:"chat"`
	From      TelegramUser     `json:"from"`
}

type TelegramPhoto struct {
	FileId   string `json:"file_id"`
	FileSize int64  `json:"file_size"`
}

type TelegramDocument struct {
	FileName string `json:"file_name"`
	MimeType string `json:"mime_type"`
	FileId   string `json:"file_id"`
	FileSize int64  `json:"file_size"`
}

type TelegramAudio struct {
	FileName string `json:"file_name"`
	MimeType string `json:"mime_type"`
	FileId   string `json:"file_id"`
	FileSize int64  `json:"file_size"`
}

type TelegramSticker struct {
	Type     string `json:"type"`
	FileId   string `json:"file_id"`
	FileSize int64  `json:"file_size"`
}

type TelegramVideo struct {
	FileName string `json:"file_name"`
	MimeType string `json:"mime_type"`
	FileId   string `json:"file_id"`
	FileSize int64  `json:"file_size"`
	Duration int64  `json:"duration"`
}

type TelegramVoice struct {
	MimeType string `json:"mime_type"`
	FileId   string `json:"file_id"`
	FileSize int64  `json:"file_size"`
	Duration int64  `json:"duration"`
}

type TelegramUser struct {
	Id       int64  `json:"id"`
	UserName string `json:"username"`
}

type TelegramChat struct {
	Id   int64  `json:"id"`
	Type string `json:"type"`
}

type ErrorMessage struct {
	Description string `json:"description"`
}
