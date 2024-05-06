package main

import (
	"body-calculator-tg-bot/internal/db"
	"body-calculator-tg-bot/internal/functions"
	sqlPkg "body-calculator-tg-bot/internal/sql"
	"database/sql"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func main() {
	//Create ".env" file in project's folder and paste the following: TG_BOT_TOKEN=YOUR_TOKEN
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Failed to load .env: %v", err)
	}

	//Connect db
	var DB *sql.DB
	DB, err = db.GetDB()
	if err != nil {
		log.Fatalf("Can't connect to database: %s", err.Error())
	}
	defer DB.Close()
	functions.SaveDB(DB)

	//Get products types
	prodTypes, err := sqlPkg.GetProductTypes(DB)
	if err != nil || prodTypes == nil {
		log.Fatalf("Cannot get product types from Database")
	}
	functions.SaveProductTypes(prodTypes)

	//Get Products
	err = functions.InitialHandlers(DB)
	if err != nil {
		log.Printf("Error with getting all products same type:", err.Error())
	}

	//Bot
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
