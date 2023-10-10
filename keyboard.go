package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func keyboardGet(keytext, keycode string) []tgbotapi.InlineKeyboardButton {
	return tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData(keytext, keycode))
}

func showMenu(update *tgbotapi.Update) {
	msg := tgbotapi.NewMessage(ChatId, "Menu")
	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		keyboardGet(KEY_TEXT_HELLO, KEY_CODE_HELLO),
		keyboardGet(KEY_TEXT_WRITE, KEY_CODE_WRITE),
	)
	BotaZamena.Send(msg)
}

func hello() {
	msg := tgbotapi.NewMessage(ChatId, "Привет")
	BotaZamena.Send(msg)
}
