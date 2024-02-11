package bodycalc

import (
	"math"
)

// For Daily Norm Of Calories
const (
	MINIMAL  = 1.2
	LIGHT    = 1.35
	MEDIUM   = 1.55
	HIGH     = 1.75
	EXTREMAL = 1.95
)

type UserState int

type BMI struct {
	Height, Weight float64 //cm, kg
}

type DailyNormOfCalories struct {
	Sex            string
	Height, Weight float64
	Age            uint8
	Activity       float64
}

func NewBMI(h float64, w float64) *BMI {
	return &BMI{h, w}
}

func (bmi *BMI) CalcBMI() float64 {
	if bmi.Weight == 0.0 || bmi.Height == 0.0 {
		return 0.0
	}
	h := bmi.Height / 100 //meters
	bmiValue := bmi.Weight / (h * h)
	return bmiValue

}

func DefineBMI(bmi float64) string {
	switch {
	case bmi < 16.0:
		return "Выраженный дефицит массы тела."
	case bmi < 18.5:
		return "Недостаточная (дефицит) массы тела."
	case bmi < 25.0:
		return "Норма."
	case bmi < 30:
		return "Избыточная масса тела (предожирение)."
	case bmi < 35:
		return "Ожирение первой степени."
	case bmi < 40:
		return "Ожирение второй степени."
	default:
		return "Ожирение третьей степени (морбидное)."
	}
}

func NewDailyNormOfCalories(sex string, h float64, w float64, age uint8, activity float64) *DailyNormOfCalories {
	return &DailyNormOfCalories{sex, h, w, age, activity}
}

func (c *DailyNormOfCalories) CalculateDailyNormOfCalories() int {
	basalMetabolism := 0.0
	switch c.Sex {
	case "male":
		basalMetabolism = 9.99*c.Weight + 6.25*c.Height - 4.92*float64(c.Age) + 5.0
	case "female":
		basalMetabolism = 9.99*c.Weight + 6.25*c.Height - 4.92*float64(c.Age) - 161.0
	}
	dnoc := basalMetabolism * c.Activity
	return int(math.Round(dnoc))

}
