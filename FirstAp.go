package main

import (
	"FirstApp/common"
	"FirstApp/lib"
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var numericKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("/start"),
	),
)

// var newb = tgbotapi.NewInlineKeyboardMarkup(
// 	tgbotapi.NewInlineKeyboardRow(
// 		tgbotapi.NewInlineKeyboardButtonData("frrf", "dede"),
// 	),
// )

func main() {

	bot, err := lib.GetCon("7930793489:AAGaZxjQgd1Qy4XAm_x6dSdpCWnCVMr74J0")
	if err != nil {
		fmt.Println("Ошибка подключения бота")
	}

	bot.Debug = true
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	var txt string
	var id int64

	lib.CheckUser(9888)
	return

	for update := range updates {
		var newb = common.GetButt("/start", "/start")
		if update.Message == nil && update.CallbackQuery == nil {
			continue
		} else if update.CallbackQuery != nil {
			// Respond to the callback query, telling Telegram to show the user
			// a message with the data received.
			callback := tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)
			if _, err := bot.Request(callback); err != nil {
				panic(err)
			}

			// And finally, send a message containing the data received.
			msgDel := tgbotapi.NewDeleteMessage(update.CallbackQuery.From.ID, update.CallbackQuery.Message.MessageID)
			bot.Send(msgDel)
			txt = update.CallbackQuery.Data
			newb = common.GetButt(update.CallbackQuery.Message.Text, update.CallbackQuery.Data)
			id = update.CallbackQuery.From.ID
		} else if update.Message != nil {
			txt = update.Message.Text

			id = update.Message.Chat.ID
		}

		msg := tgbotapi.NewMessage(id, txt)
		msg.ReplyMarkup = numericKeyboard
		msg.ReplyMarkup = newb
		bot.Send(msg)

	}

}
