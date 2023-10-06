package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func keyboradGet(keytext, keycode string) []tgbotapi.InlineKeyboardButton {
	return tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData(keytext, keycode))
}

func showMenu() {
	message := tgbotapi.NewMessage(ChatId, "Меню")
	message.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		keyboradGet(KEY_TEXT_HELLO, KEY_CODE_HELLO),
		keyboradGet(KEY_TEXT_BYE, KEY_CODE_HELLO),
	)
	BotaZamena.Send(message)
}

func hello() {
	sendSystemMessageWithDelay(2, "Пизда")
}
func bye() {
	sendSystemMessageWithDelay(2, "Не пизда")
}
