package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	// подключение .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// подключение к боту через токен
	bot, err := tgbotapi.NewBotAPI(os.Getenv("API_TOKEN"))
	if err != nil {
		panic(err)
	}

	bot.Debug = true

	// сообщаем о том, что мы обработали предыдущие сообщенияи нам не нужно их повторять
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30

	// Проверка наличия обновлений
	updates := bot.GetUpdatesChan(updateConfig)

	// Цикл идет по всем обновлениям, полученным каналом
	for update := range updates {
		// если это не сообщение, то пропускаем это обновление
		if update.Message == nil {
			continue
		}

		// отправка сообщения
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
		switch update.Message.Command() {
		case "create":
			msg.Text = "Напишите текст, который хотите поместить на картинке"
		case "help":
			msg.Text = "Пока здесь ничего нет"
		default:
			msg.Text = "Я не знаю такой команды"
		}
		if _, err := bot.Send(msg); err != nil {
			panic(err)
		}
	}
}
