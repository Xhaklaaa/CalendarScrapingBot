package main

import (
	"fmt"
	"log"
	"os"
	"time"

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

func delay(seconds uint8) {
	time.Sleep(time.Second * time.Duration(seconds))
}

func printSystemMessageWithDelay(delayInSec uint8, message string) {
	gBot.Send(tgbotapi.NewMessage(gChatId, message))
	delay(delayInSec)

}

func printIntro(update *tgbotapi.Update) {
	printSystemMessageWithDelay(2, "Привет!"+EMOJI_HUGGING)
}

func main() {

	log.Printf("Authorized on account %s", gBot.Self.UserName)

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60

	for update := range gBot.GetUpdatesChan(updateConfig) {
		if isStartMessage(&update) {
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
			gChatId = update.Message.Chat.ID
			printIntro(&update)
		}
		//TODO: implement proccessing other types of updates
	}
}
