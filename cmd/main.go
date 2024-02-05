package main

import (
	"body-calculator-tg-bot/pkg/bodycalc"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

type UserState int

const (
	InitialState UserState = iota
	WeightState
	HeightState
)

const (
	minWeight = 3.0
	maxWeight = 450.0
	minHeight = 40.0
	maxHeight = 225.0
)

var (
	bot          *tgbotapi.BotAPI
	bmi          = bodycalc.NewBMI(0.0, 0.0)
	userStates   map[int64]UserState
	startMessage = "Привет! Данный калькулятор позволит рассчитать ИМТ, количество калорий, которое необходимо " +
		"вашему организму в зависимости от вашего роста, веса, возраста и степени физической активности. Для начала перейдите в меню:"
	goToMenu = tgbotapi.NewInlineKeyboardMarkup(tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData("Меню", "menu")))
	menu     = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("ИМТ", "bmi"),
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

	userStates = make(map[int64]UserState) // Хранение состояний пользователей

	for update := range updates {
		if update.CallbackQuery != nil {
			callbacks(update)
		} else if update.Message.IsCommand() {
			commands(update)
		} else if update.Message != nil {
			chooseState(userStates, update)
		}
	}

}

func callbacks(update tgbotapi.Update) {
	cb := update.CallbackQuery.Data
	chatId := update.CallbackQuery.Message.Chat.ID
	switch cb {
	case "menu":
		userStates[chatId] = InitialState
		msg := tgbotapi.NewMessage(chatId, "Выберите необходимую опцию:")
		msg.ReplyMarkup = menu
		sendMessage(msg)
	case "bmi":
		chatID := update.CallbackQuery.Message.Chat.ID

		msg := tgbotapi.NewMessage(chatID, "Введите ваш вес:")
		bot.Send(msg)

		userStates[chatID] = WeightState

	}
}

func commands(update tgbotapi.Update) {
	command := update.Message.Command()
	chatId := update.Message.Chat.ID
	switch command {
	case "start":
		userStates[chatId] = InitialState
		msg := tgbotapi.NewMessage(chatId, startMessage)
		msg.ReplyMarkup = goToMenu
		sendMessage(msg)
	case "menu":
		userStates[chatId] = InitialState
		msg := tgbotapi.NewMessage(chatId, "Выберите необходимую опцию:")
		msg.ReplyMarkup = menu
		sendMessage(msg)
	default:
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Неизвестная команда. Вернуться в меню: /menu")
		sendMessage(msg)
	}
}

func sendMessage(msg tgbotapi.Chattable) {
	if _, err := bot.Send(msg); err != nil {
		log.Panicf("Failed to send message: %v", err)
	}
	log.Println("Message was sended.")
}

func chooseState(userStates map[int64]UserState, update tgbotapi.Update) {
	chatId := update.Message.Chat.ID
	switch userStates[chatId] {
	case InitialState:
		bmi = bodycalc.NewBMI(0.0, 0.0)
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Неизвестная команда. Вернуться в меню: /menu")
		sendMessage(msg)
	case WeightState:
		weight, err := strconv.ParseFloat(update.Message.Text, 64)
		if err != nil {
			msg := tgbotapi.NewMessage(chatId, "Ошибка! Введите число для веса.")
			sendMessage(msg)
			return
		}
		if weight < minWeight || weight > maxWeight {
			msg := tgbotapi.NewMessage(chatId, "Ошибка! Введите корректное значение веса")
			sendMessage(msg)
			return
		}

		bmi.Weight = weight

		msg := tgbotapi.NewMessage(chatId, "Введите ваш рост:")
		bot.Send(msg)

		userStates[chatId] = HeightState

	case HeightState:
		height, err := strconv.ParseFloat(update.Message.Text, 64)
		if err != nil {
			msg := tgbotapi.NewMessage(chatId, "Ошибка! Введите число для роста.")
			sendMessage(msg)
			return
		}
		if height < minHeight || height > maxHeight {
			msg := tgbotapi.NewMessage(chatId, "Ошибка! Введите корректное значение роста")
			sendMessage(msg)
			return
		}
		//calculateBMI
		bmi.Height = height
		bmiValue := bmi.CalcBMI()

		msg := tgbotapi.NewMessage(chatId, "Ваш ИМТ: "+strconv.FormatFloat(bmiValue, 'f', 2, 64))
		bot.Send(msg)

		userStates[chatId] = InitialState

	}

	//TODO: more info about bmi
}
