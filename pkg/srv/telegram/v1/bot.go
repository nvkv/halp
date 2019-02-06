package telegram

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/nvkv/halp/pkg/types/datasource/v1"
)

type Bot struct {
	Token            string
	Datasource       datasource.Datasource
	WhitelistedChats []int64
	bot              *tgbotapi.BotAPI
}

// Reply only in whitelisted chats
func (b Bot) isItSafeToReply(update tgbotapi.Update) bool {
	for _, chatid := range b.WhitelistedChats {
		if chatid == update.Message.Chat.ID {
			return true
		}
	}
	return false
}

func (b Bot) Start() error {
	bot, err := tgbotapi.NewBotAPI(b.Token)
	if err != nil {
		return err
	}

	b.bot = bot
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		fmt.Println(err)
	}

	for update := range updates {
		// Skip everything except commands
		if update.Message == nil || update.Message.IsCommand() != true {
			continue
		}

		if b.isItSafeToReply(update) == false {
			foffMsg := tgbotapi.NewMessage(update.Message.Chat.ID, "Кто вы? Идите на хер! Я вас не знаю!")
			if _, err := bot.Send(foffMsg); err != nil {
				fmt.Println(err)
			}
			continue
		}

		b.processCommand(update)
	}

	return nil
}
