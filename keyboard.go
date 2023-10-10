package main

import (
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func keyboradGet(keytext, keycode string) []tgbotapi.InlineKeyboardButton {
	return tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData(keytext, keycode))
}

func showMenu(update *tgbotapi.Update) {
	msg := tgbotapi.NewMessage(ChatId, "")
	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		keyboradGet(KEY_TEXT_HELLO, KEY_CODE_HELLO),
		keyboradGet(KEY_TEXT_DATA, KEY_CODE_WRITE),
	)
	BotaZamena.Send(msg)
}

func hello() {
	msg := tgbotapi.NewMessage(ChatId, "Привет")
	BotaZamena.Send(msg)
}
func data() {
	GenerateCalendar(time.Now().Year(), time.Now().Month())
}
