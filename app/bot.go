package main

import (
	"fmt"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var gBot *tgbotapi.BotAPI
var gToken string
var gChatId int64

func init() {
	_ = os.Setenv(TOKEN_NAME_IN_OS, "6635332877:AAE_EA1Rvlfw6sLLJxfhIyRrDTtmyd20PG8")

	if gToken = os.Getenv(TOKEN_NAME_IN_OS); gToken == "" {
		panic(fmt.Errorf(`failed to load enviroment variable "%s"`, TOKEN_NAME_IN_OS))
	}
	var err error
	gBot, err := tgbotapi.NewBotAPI("gToken")
	if err != nil {
		log.Panic(err)
	}
	gBot.Debug = true
}

func isStartMessage(update *tgbotapi.Update) bool {
	return update.Message != nil && update.Message.Text == "/start"
}

func main() {

	log.Printf("Authorized on account %s", gBot.Self.UserName)

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60

	for update := range gBot.GetUpdatesChan(updateConfig) {
		if isStartMessage(&update) {
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			msg.ReplyToMessageID = update.Message.MessageID

			gBot.Send(msg)
		}
		//TODO: implement proccessing other types of updates
	}
}
