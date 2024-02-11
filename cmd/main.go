package main

import (
	"body-calculator-tg-bot/int/functions"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	//Create ".env" file in project's folder and paste the following: TG_BOT_TOKEN=YOUR_TOKEN
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Failed to load .env: %v", err)
	}

	functions.Bot, err = tgbotapi.NewBotAPI(os.Getenv("TG_BOT_TOKEN"))
	if err != nil {
		log.Fatalf("Uncorrect token.")
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := functions.Bot.GetUpdatesChan(u)

	log.Println("Bot was started.")

	functions.UserStates = make(map[int64]functions.UserState) // Хранение состояний пользователей

	for update := range updates {
		if update.CallbackQuery != nil {
			functions.Callbacks(update)
		} else if update.Message.IsCommand() {
			functions.Commands(update)
		} else if update.Message != nil {
			functions.ChooseState(update, functions.UserStates)
		}
	}

}
