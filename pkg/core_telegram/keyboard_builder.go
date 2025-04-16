package core_telegram

type IKeyboard interface{}

type ReplyKeyboardMarkup struct {
	Keyboard        *[][]ReplyKeyboardButton `json:"keyboard"`
	OneTimeKeyboard bool                     `json:"one_time_keyboard"`
	ResizeKeyboard  bool                     `json:"resize_keyboard"`
	IKeyboard
}

type ReplyKeyboardButton struct {
	Text string `json:"text"`
}

type InlineKeyboardMarkup struct {
	Keyboard *[][]InlineKeyboardButton `json:"inline_keyboard"`
}

type InlineKeyboardButton struct {
	Text     string `json:"text"`
	Callback string `json:"callback_data"`
}

func NewInlineKeyboardMarkup(buttonsCallback *[][]InlineKeyboardButton) IKeyboard {
	if buttonsCallback == nil {
		return nil
	}

	return &InlineKeyboardMarkup{
		Keyboard: buttonsCallback,
	}
}

func NewReplyKeyboardMarkup(buttonsText *[][]string) IKeyboard {
	var buttons = make([][]ReplyKeyboardButton, 0)

	if buttonsText == nil {
		return nil
	}

	for _, textRow := range *buttonsText {
		var row []ReplyKeyboardButton
		for _, buttonText := range textRow {
			row = append(row, ReplyKeyboardButton{
				Text: buttonText,
			})
		}
		buttons = append(buttons, row)
	}

	return &ReplyKeyboardMarkup{
		Keyboard:        &buttons,
		OneTimeKeyboard: false,
		ResizeKeyboard:  true,
	}
}
