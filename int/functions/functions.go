package functions

import (
	"body-calculator-tg-bot/pkg/bodycalc"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"math"
	"strconv"
)

var (
	Bot          *tgbotapi.BotAPI
	bmi          = bodycalc.NewBMI(0.0, 0.0)
	calories     = bodycalc.NewDailyNormOfCalories("", 0.0, 0.0, 0, 0.0)
	UserStates   map[int64]UserState
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
	activitiesMenu = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Минимальная", "minimal"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Легкая", "light"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Средняя", "medium"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Высокая", "high"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Экстремальная", "extremal"),
		),
	)
	sexMenu = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData("Мужской", "male")),
		tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData("Женский", "female")))
)

type UserState int

const (
	InitialState UserState = iota
	WeightStateBMI
	HeightStateBMI
	SexStateDNOC
	WeightStateDNOC
	HeightStateDNOC
	AgeStateDNOC
	ActivityStateDNOC
)

const (
	minWeight = 3.0
	maxWeight = 450.0
	minHeight = 40.0
	maxHeight = 225.0
	minAge    = 0
	maxAge    = 100
)

func Callbacks(update tgbotapi.Update) {
	cb := update.CallbackQuery.Data
	chatId := update.CallbackQuery.Message.Chat.ID
	switch cb {
	case "menu":
		UserStates[chatId] = InitialState
		msg := tgbotapi.NewMessage(chatId, "Выберите необходимую опцию:")
		msg.ReplyMarkup = menu
		sendMessage(msg)
	case "bmi":
		msg := tgbotapi.NewMessage(chatId, "Введите ваш вес:")
		sendMessage(msg)

		UserStates[chatId] = WeightStateBMI
	case "calories":
		msg := tgbotapi.NewMessage(chatId, "Выберите ваш пол: ")
		msg.ReplyMarkup = sexMenu
		sendMessage(msg)

		UserStates[chatId] = SexStateDNOC
	case "male":
		calories.Sex = "male"
		msg := tgbotapi.NewMessage(chatId, "Введите ваш вес:")
		sendMessage(msg)

		UserStates[chatId] = WeightStateDNOC
	case "female":
		calories.Sex = "female"
		msg := tgbotapi.NewMessage(chatId, "Введите ваш вес:")
		sendMessage(msg)

		UserStates[chatId] = WeightStateDNOC
	case "minimal":
		calories.Activity = bodycalc.MINIMAL
		SendDNOC(update, calories)
	case "light":
		calories.Activity = bodycalc.LIGHT
		SendDNOC(update, calories)
	case "medium":
		calories.Activity = bodycalc.MEDIUM
		SendDNOC(update, calories)
	case "high":
		calories.Activity = bodycalc.HIGH
		SendDNOC(update, calories)
	case "extremal":
		calories.Activity = bodycalc.EXTREMAL
		SendDNOC(update, calories)

	}
}

func Commands(update tgbotapi.Update) {
	command := update.Message.Command()
	chatId := update.Message.Chat.ID
	switch command {
	case "start":
		UserStates[chatId] = InitialState
		msg := tgbotapi.NewMessage(chatId, startMessage)
		msg.ReplyMarkup = goToMenu
		sendMessage(msg)
	case "menu":
		UserStates[chatId] = InitialState
		msg := tgbotapi.NewMessage(chatId, "Выберите необходимую опцию:")
		msg.ReplyMarkup = menu
		sendMessage(msg)
	default:
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Неизвестная команда. Вернуться в меню: /menu")
		sendMessage(msg)
	}
}

func sendMessage(msg tgbotapi.Chattable) {
	if _, err := Bot.Send(msg); err != nil {
		log.Panicf("Failed to send message: %v", err)
	}
	log.Println("Message was sended.")
}

func ChooseState(update tgbotapi.Update, userStates map[int64]UserState) {
	chatId := update.Message.Chat.ID
	switch userStates[chatId] {
	case InitialState:
		bmi = bodycalc.NewBMI(0.0, 0.0)
		calories = bodycalc.NewDailyNormOfCalories("", 0.0, 0.0, 0, 0.0)
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Неизвестная команда. Вернуться в меню: /menu")
		sendMessage(msg)
	case WeightStateBMI:
		weight, err := strconv.ParseFloat(update.Message.Text, 64)
		if err != nil {
			msg := tgbotapi.NewMessage(chatId, "Ошибка! Введите число для веса.")
			sendMessage(msg)
			return
		}
		if weight < minWeight || weight > maxWeight {
			msg := tgbotapi.NewMessage(chatId, "Ошибка! Введите корректное значение веса.")
			sendMessage(msg)
			return
		}

		bmi.Weight = weight

		msg := tgbotapi.NewMessage(chatId, "Введите ваш рост:")
		sendMessage(msg)

		userStates[chatId] = HeightStateBMI

	case HeightStateBMI:
		height, err := strconv.ParseFloat(update.Message.Text, 64)
		if err != nil {
			msg := tgbotapi.NewMessage(chatId, "Ошибка! Введите число для роста.")
			sendMessage(msg)
			return
		}
		if height < minHeight || height > maxHeight {
			msg := tgbotapi.NewMessage(chatId, "Ошибка! Введите корректное значение роста.")
			sendMessage(msg)
			return
		}
		//calculateBMI
		bmi.Height = height
		bmiValue := bmi.CalcBMI()

		bmiInfo := bodycalc.DefineBMI(bmiValue)

		msg := tgbotapi.NewMessage(chatId, "Ваш ИМТ: "+strconv.FormatFloat(bmiValue, 'f', 2, 64)+". "+bmiInfo)
		sendMessage(msg)

		userStates[chatId] = InitialState
	case SexStateDNOC:
		msg := tgbotapi.NewMessage(chatId, "Ошибка. Выберите ваш пол: ")
		msg.ReplyMarkup = sexMenu

		sendMessage(msg)
	case WeightStateDNOC:
		weight, err := strconv.ParseFloat(update.Message.Text, 64)
		if err != nil {
			msg := tgbotapi.NewMessage(chatId, "Ошибка! Введите число для веса.")
			sendMessage(msg)
			return
		}
		if weight < minWeight || weight > maxWeight {
			msg := tgbotapi.NewMessage(chatId, "Ошибка! Введите корректное значение веса.")
			sendMessage(msg)
			return
		}
		calories.Weight = weight

		msg := tgbotapi.NewMessage(chatId, "Введите ваш рост:")
		sendMessage(msg)

		userStates[chatId] = HeightStateDNOC
	case HeightStateDNOC:
		height, err := strconv.ParseFloat(update.Message.Text, 64)
		if err != nil {
			msg := tgbotapi.NewMessage(chatId, "Ошибка! Введите число для роста.")
			sendMessage(msg)
			return
		}
		if height < minHeight || height > maxHeight {
			msg := tgbotapi.NewMessage(chatId, "Ошибка! Введите корректное значение роста.")
			sendMessage(msg)
			return
		}
		calories.Height = height

		msg := tgbotapi.NewMessage(chatId, "Введите ваш возраст:")
		sendMessage(msg)

		userStates[chatId] = AgeStateDNOC
	case AgeStateDNOC:
		age, err := strconv.ParseUint(update.Message.Text, 10, 8)
		if err != nil {
			msg := tgbotapi.NewMessage(chatId, "Ошибка! Некорректное число для возраста")
			sendMessage(msg)
			return
		}
		if age < minAge || age > maxAge {
			msg := tgbotapi.NewMessage(chatId, "Ошибка! Введите возраст в пределах жизни человека")
			sendMessage(msg)
			return
		}

		calories.Age = uint8(age)

		msg := tgbotapi.NewMessage(chatId, "Выберите вашу физическую активность:")
		msg.ReplyMarkup = activitiesMenu
		sendMessage(msg)

		userStates[chatId] = ActivityStateDNOC
	case ActivityStateDNOC:
		msg := tgbotapi.NewMessage(chatId, "Ошибка. Выберите вашу физическую активность:")
		msg.ReplyMarkup = activitiesMenu
		sendMessage(msg)
	}

}

func SendDNOC(update tgbotapi.Update, dnoc *bodycalc.DailyNormOfCalories) {
	chatID := update.CallbackQuery.Message.Chat.ID
	dnocValue := dnoc.CalculateDailyNormOfCalories()
	resString := fmt.Sprintf("Для сохранения веса суточная потебность в калориях возрастом %v лет и весом %.2f кг., при росте %.2f см. и вашей активности:\n%d ккал\n\n", dnoc.Age, dnoc.Weight, dnoc.Height, dnocValue)
	minCal := int(math.Round(float64(dnocValue) * 1.15))
	maxCal := int(math.Round(float64(dnocValue) * 0.85))
	resString += fmt.Sprintf("Для набора веса рекомендуется потребление %d ккал в день, а для похудения стоит употреблять порядка %d ккал в день.", minCal, maxCal)
	msg := tgbotapi.NewMessage(chatID, resString)
	sendMessage(msg)
	UserStates[chatID] = InitialState
}
