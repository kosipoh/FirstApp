package common

import (
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func GetButt(str string, strBack string) tg.InlineKeyboardMarkup {

	mas := make([][]string, 0)
	if str == "/start" {
		temp := make([]string, 3)
		temp[0] = "Рейтинг"
		temp[1] = "Добавить"
		temp[2] = "Мои"
		mas = append(mas, []string{"Рейтинг", "/start"})
		mas = append(mas, []string{"Добавить", "/start"})
		mas = append(mas, []string{"Мои", "/start"})
	} else if str == "Назад" {
		mas = append(mas, []string{"Назад", strBack})
	} else if str == "Рейтинг" {
		mas = append(mas, []string{"оаткоткоа", "/start"})
	}

	masBut := make([]tg.InlineKeyboardButton, 0)
	for _, elem := range mas {
		masBut = append(masBut, tg.NewInlineKeyboardButtonData(elem[0], elem[1]))
	}

	//masBut = append(masBut, tg.NewInlineKeyboardButtonData("Назад", "Назад"))
	return tg.NewInlineKeyboardMarkup(tg.NewInlineKeyboardRow(masBut...))

}
