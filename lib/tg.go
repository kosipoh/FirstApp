package lib

import (
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func GetCon(token string) (tg.BotAPI, error) {
	bot, err := tg.NewBotAPI(token)
	return *bot, err
}
