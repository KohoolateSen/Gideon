package bot

import (
	"sort"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (tb *TelegramBot) SendRawMessage(chat_id int64, text string) {
	msg := tgbotapi.NewMessage(chat_id, text)
	tb.Send(msg)
}

func (tb *TelegramBot) CustomSingleButtonKeyboard(label string, value string) *tgbotapi.InlineKeyboardMarkup {
	inline_keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonURL(label, value),
		),
	)

	return &inline_keyboard
}

func (tb *TelegramBot) CustomMultiButtonKeyboard(options map[int]string) *tgbotapi.InlineKeyboardMarkup {
	var buttons []tgbotapi.InlineKeyboardButton
	var keys []int

	for key := range options {
		keys = append(keys, key)
	}

	sort.Ints(keys)

	for _, key := range keys {
		value := options[key]
		label := strings.Split(value, ":")[1]
		buttons = append(buttons, tgbotapi.NewInlineKeyboardButtonData(label, value))
	}

	inline_keyboard := tgbotapi.NewInlineKeyboardMarkup(tgbotapi.NewInlineKeyboardRow(buttons...))
	return &inline_keyboard
}

func (tb *TelegramBot) SendMessageWithMultiButtonKeyboard(chat_id int64, message string, options map[int]string) {
	inline_keyboard := tb.CustomMultiButtonKeyboard(options)

	msg := tgbotapi.NewMessage(chat_id, message)
	msg.ReplyMarkup = inline_keyboard
	msg.Text = message

	tb.Send(msg)
}

func (tb *TelegramBot) SendMessageWithSingleButtonKeyboard(chat_id int64, message, label string, value string) {
	inline_keyboard := tb.CustomSingleButtonKeyboard(label, value)

	msg := tgbotapi.NewMessage(chat_id, message)
	msg.ReplyMarkup = inline_keyboard
	msg.Text = message

	tb.Send(msg)
}
