package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var keyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Создать картинку"),
		tgbotapi.NewKeyboardButton("Помощь"),
	),
)

func main() {
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

	log.Printf("Authorized on account %s", bot.Self.UserName)

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
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		switch update.Message.Command() {
		case "start":
			msg.ReplyMarkup = keyboard
		default:
			switch update.Message.Text {
			case "Создать картинку":
				img := tgbotapi.NewPhoto(update.Message.Chat.ID, tgbotapi.FilePath("image.png"))
				if _, err := bot.Send(img); err != nil {
					panic(err)
				}
			case "Помощь":
				msg.Text = "Это помощь"
			default:
				msg.Text = "Для начала работы введите /start"
			}
		}
		if _, err := bot.Send(msg); err != nil {
			panic(err)
		}
	}
}
