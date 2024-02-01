package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var (
	bot           *tgbotapi.BotAPI
	start_message = "Привет! Данный калькулятор позволит рассчитать ИМТ, количество калорий, которое необходимо " +
		"вашему организму в зависимости от вашего роста, веса, возраста и степени физической активности. Выберите, что хотите рассчитать:"
	start_menu = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("ИМТ", "BMI"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Норма калорий", "calories"),
		),
	)
)

func main() {
	//Create ".env" file in project's folder and paste the following: TG_BOT_TOKEN=YOUR_TOKEN
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Failed to load .env: %v", err)
	}

	bot, err = tgbotapi.NewBotAPI(os.Getenv("TG_BOT_TOKEN"))
	if err != nil {
		log.Fatalf("Uncorrect token.")
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	log.Println("Bot was started.")

	for update := range updates {
		if update.CallbackQuery != nil {
			callbacks(update)
		} else if update.Message.IsCommand() {
			commands(update)
		} else {
			//messages
		}
	}

}

func callbacks(update tgbotapi.Update) {

}

func commands(update tgbotapi.Update) {
	command := update.Message.Command()
	switch command {
	case "start":
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, start_message)
		msg.ReplyMarkup = start_menu
		sendMessage(msg)

	}
}

func sendMessage(msg tgbotapi.Chattable) {
	if _, err := bot.Send(msg); err != nil {
		log.Panicf("Failed to send message: %v", err)
	}
	log.Println("Message was sended.")
}
