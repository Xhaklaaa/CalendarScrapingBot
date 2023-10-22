package main

import (
	"log"
	"os"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var BotaZamena *tgbotapi.BotAPI
var Token string
var ChatId int64

func startMessage(update *tgbotapi.Update) bool {
	return update.Message != nil && update.Message.Text == "/start"
}

func CallbackQuerybot(update *tgbotapi.Update) bool {
	return update.CallbackQuery != nil && update.CallbackQuery.Data != ""
}

func init() {
	_ = os.Setenv(TOKEN_NAME, "6635332877:AAE_EA1Rvlfw6sLLJxfhIyRrDTtmyd20PG8")
	Token = os.Getenv(TOKEN_NAME)

	var err error
	if BotaZamena, err = tgbotapi.NewBotAPI(Token); err != nil {
		log.Panic(err)
	}
	BotaZamena.Debug = true
}
func updateProcess(update *tgbotapi.Update) {
	choice := update.CallbackQuery.Data
	log.Printf("[%T] %s", time.Now(), choice)

}

func main() {

	BotaZamena.Debug = true

	log.Printf("Authorized on account %s", BotaZamena.Self.UserName)

	update_config := tgbotapi.NewUpdate(0)
	update_config.Timeout = UPDATE_CONFIG_TIMEOUT

	updates := BotaZamena.GetUpdatesChan(update_config)

	for update := range updates {
		if CallbackQuerybot(&update) {
			updateProcess(&update)
		} else if startMessage(&update) {
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
			ChatId = update.Message.Chat.ID

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Привет, я полезный телеграм бот! Что бы выбрать дату введите '/calendar'")
			BotaZamena.Send(msg)
			calendar(&update)
		}
	}
}
