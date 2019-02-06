package telegram

import (
	"fmt"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/nvkv/halp/pkg/schedule/v1"
)

func (b Bot) processCommand(update tgbotapi.Update) {
	placeholderMsg := tgbotapi.NewMessage(update.Message.Chat.ID, "Дайте-ка подумаю...")
	if _, err := b.bot.Send(placeholderMsg); err != nil {
		fmt.Println(err)
	}

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
	switch update.Message.Command() {
	case "today":
		schedule, err := schedule.ScheduleDay(time.Now(), b.Datasource)
		if err != nil {
			msg.Text = err.Error()
		} else {
			msg.Text = formatDaySchedule(schedule)
		}

	case "tomorrow":
		tomorrow := time.Now().AddDate(0, 0, 1)
		schedule, err := schedule.ScheduleDay(tomorrow, b.Datasource)
		if err != nil {
			msg.Text = err.Error()
		} else {
			msg.Text = formatDaySchedule(schedule)
		}
	}

	if _, err := b.bot.Send(msg); err != nil {
		fmt.Println(err)
	}
}
