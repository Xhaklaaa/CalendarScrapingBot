package app

import (
	"log"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	tb_cal "github.com/oramaz/telebot-calendar"
	tb "gopkg.in/telebot.v3"
)

func calendar(update *tgbotapi.Update) {
	b, err := tb.NewBot(tb.Settings{
		Token:  "6635332877:AAE_EA1Rvlfw6sLLJxfhIyRrDTtmyd20PG8",
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("/calendar", func(c tb.Context) error {
		calendar := tb_cal.NewCalendar(b, tb_cal.Options{})
		return c.Send("Select a date", &tb.ReplyMarkup{
			InlineKeyboard: calendar.GetKeyboard(),
		})
	})

	b.Handle(tb.OnText, func(c tb.Context) error {
		date, err := time.Parse("02.01.2006", c.Data())
		if err != nil {
			return err
		}
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, date.String())
		BotaZamena.Send(msg)

		c.Respond()
		return nil
	})

	b.Start()
}
