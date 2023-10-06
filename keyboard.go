package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func keyboradGet(keytext, keycode string) []tgbotapi.InlineKeyboardButton {
	return tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData(keytext, keycode))
}

func showMenu(update *tgbotapi.Update) {
	msg := tgbotapi.NewMessage(ChatId, "Меню")
	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		keyboradGet(KEY_TEXT_HELLO, KEY_CODE_HELLO),
		keyboradGet(KEY_TEXT_BYE, KEY_CODE_HELLO),
	)
	BotaZamena.Send(msg)
}

func hello() {
	msg := tgbotapi.NewMessage(ChatId, "Привет")
	BotaZamena.Send(msg)
}
func bye() {
	msg := tgbotapi.NewMessage(ChatId, "Пока")
	BotaZamena.Send(msg)
}
