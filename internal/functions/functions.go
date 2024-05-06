package functions

import (
	sqlPkg "body-calculator-tg-bot/internal/sql"
	"body-calculator-tg-bot/pkg/bodycalc"
	"database/sql"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"math"
	"strconv"
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
	ProductTypesState
)

const (
	minWeight       = 3.0
	maxWeight       = 450.0
	minHeight       = 40.0
	maxHeight       = 225.0
	minAge          = 0
	maxAge          = 130
	EMOJI_CHECKMARK = "\U00002705"
	EMOJI_BACK      = "\U0001F519"
	EMOJI_PLAY      = "\U000025B6"
	EMOJI_HI        = "\U0001F44B"
	EMOJI_ONE       = "\U00000031\U0000FE0F\U000020E3 "
	EMOJI_TWO       = "\U00000032\U0000FE0F\U000020E3 "
	EMOJI_THREE     = "\U00000033\U0000FE0F\U000020E3 "
	EMOJI_FOUR      = "\U00000034\U0000FE0F\U000020E3 "
	EMOJI_FIVE      = "\U00000035\U0000FE0F\U000020E3 "
	EMOJI_MAN       = " \U0001F468"
	EMOJI_WOMAN     = " \U0001F469"
)

var (
	Bot          *tgbotapi.BotAPI
	DB           *sql.DB
	bmi          = bodycalc.NewBMI(0.0, 0.0)
	calories     = bodycalc.NewDailyNormOfCalories("", 0.0, 0.0, 0, 0.0)
	UserStates   map[int64]UserState
	startMessage = "Привет " + EMOJI_HI + " Данный калькулятор позволит рассчитать ИМТ, количество калорий, которое необходимо " +
		"вашему организму в зависимости от вашего роста, веса, возраста и степени физической активности.\nДля начала перейдите в меню:"
	goToMenu = tgbotapi.NewInlineKeyboardMarkup(tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData("Меню "+EMOJI_PLAY, "menu")))
	menu     = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(EMOJI_ONE+"ИМТ", "bmi"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(EMOJI_TWO+"Норма калорий", "calories"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(EMOJI_THREE+"Список продуктов (В процессе разработки)", "products"),
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
		tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData("Мужской"+EMOJI_MAN, "male")),
		tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData("Женский"+EMOJI_WOMAN, "female")))
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
	case "products":
		var menuPrTypes tgbotapi.InlineKeyboardMarkup
		UserStates[chatId] = ProductTypesState
		err := DB.Ping()
		if err != nil {
			msg := tgbotapi.NewMessage(chatId, "Ошибка! Не удалось подключиться к базе данных продуктов. Вернуться в меню: /menu "+EMOJI_BACK)
			fmt.Println(err.Error())
			sendMessage(msg)
			UserStates[chatId] = InitialState
			return
		}
		productsTypes, err := sqlPkg.GetProductTypes(DB)
		if err != nil || productsTypes == nil {
			msg := tgbotapi.NewMessage(chatId, "Ошибка! Не найдены типы продуктов. Вернуться в меню: /menu "+EMOJI_BACK)
			fmt.Println(err.Error())
			sendMessage(msg)
			UserStates[chatId] = InitialState
			return
		}
		for _, t := range productsTypes {
			var row []tgbotapi.InlineKeyboardButton
			btn := tgbotapi.NewInlineKeyboardButtonData(t, t)
			row = append(row, btn)
			menuPrTypes.InlineKeyboard = append(menuPrTypes.InlineKeyboard, row)
		}
		msg := tgbotapi.NewMessage(chatId, "Выберите тип продукта:")
		msg.ReplyMarkup = menuPrTypes
		sendMessage(msg)
		/*TODO: Choose State*/

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
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Неизвестная команда. Вернуться в меню: /menu "+EMOJI_BACK)
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
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Неизвестная команда. Вернуться в меню: /menu "+EMOJI_BACK)
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

		aboutActivities := EMOJI_ONE + "Минимальная: сидячая работа, отсутствие спорта;\n\n" +
			EMOJI_TWO + "Легкая: легкие физические упражнения около 3 раз за неделю, ежедневная утренняя зарядка, пешие прогулки;\n\n" +
			EMOJI_THREE + "Средняя: спорт до 5 раз за неделю;\n\n" +
			EMOJI_FOUR + "Высокая: активный образ жизни вкупе с ежедневными интенсивными тренировками;\n\n" +
			EMOJI_FIVE + "Экстремальная: максимальная активность - спортивный образ жизни, тяжелый физический труд, длительные тяжелые тренировки каждый день."
		msg := tgbotapi.NewMessage(chatId, "Выберите вашу физическую активность:\n"+aboutActivities)
		msg.ReplyMarkup = activitiesMenu
		sendMessage(msg)

		userStates[chatId] = ActivityStateDNOC
	case ActivityStateDNOC:
		msg := tgbotapi.NewMessage(chatId, "Ошибка. Выберите вашу физическую активность:")
		msg.ReplyMarkup = activitiesMenu
		sendMessage(msg)
	case ProductTypesState:
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Неизвестная команда. Вернуться в меню: /menu "+EMOJI_BACK)
		sendMessage(msg)
		userStates[chatId] = InitialState
	}

}

func SendDNOC(update tgbotapi.Update, dnoc *bodycalc.DailyNormOfCalories) {
	chatID := update.CallbackQuery.Message.Chat.ID
	dnocValue := dnoc.CalculateDailyNormOfCalories()
	resString := fmt.Sprintf("Для сохранения веса суточная потебность в калориях возрастом %v лет и весом %.2f кг., при росте %.2f см. и вашей активности:\n%d ккал%v\n\n", dnoc.Age, dnoc.Weight, dnoc.Height, dnocValue, EMOJI_CHECKMARK)
	minCal := int(math.Round(float64(dnocValue) * 1.15))
	maxCal := int(math.Round(float64(dnocValue) * 0.85))
	resString += fmt.Sprintf("Для набора веса рекомендуется потребление %d ккал%v в день, а для похудения стоит употреблять порядка %d ккал%v в день.", minCal, EMOJI_CHECKMARK, maxCal, EMOJI_CHECKMARK)
	msg := tgbotapi.NewMessage(chatID, resString)
	sendMessage(msg)
	UserStates[chatID] = InitialState
}
