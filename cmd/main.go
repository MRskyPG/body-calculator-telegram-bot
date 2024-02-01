package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var (
	bot          *tgbotapi.BotAPI
	startMessage = "Привет! Данный калькулятор позволит рассчитать ИМТ, количество калорий, которое необходимо " +
		"вашему организму в зависимости от вашего роста, веса, возраста и степени физической активности. Для начала перейдите в меню:"
	goToMenu = tgbotapi.NewInlineKeyboardMarkup(tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData("Меню", "menu")))
	menu     = tgbotapi.NewInlineKeyboardMarkup(
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
	cb := update.CallbackQuery.Data
	chatId := update.CallbackQuery.Message.Chat.ID
	switch cb {
	case "menu":
		msg := tgbotapi.NewMessage(chatId, "Выберите необходимую опцию:")
		msg.ReplyMarkup = menu
		sendMessage(msg)
	}
}

func commands(update tgbotapi.Update) {
	command := update.Message.Command()
	chatId := update.Message.Chat.ID
	switch command {
	case "start":
		msg := tgbotapi.NewMessage(chatId, startMessage)
		msg.ReplyMarkup = goToMenu
		sendMessage(msg)
	case "menu":
		msg := tgbotapi.NewMessage(chatId, "Выберите необходимую опцию:")
		msg.ReplyMarkup = menu
		sendMessage(msg)
	}
}

func sendMessage(msg tgbotapi.Chattable) {
	if _, err := bot.Send(msg); err != nil {
		log.Panicf("Failed to send message: %v", err)
	}
	log.Println("Message was sended.")
}
