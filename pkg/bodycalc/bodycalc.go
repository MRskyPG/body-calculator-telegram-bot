package bodycalc

import "errors"

type BMI struct {
	Height float32 //cm
	Weight float32 //kg
}

func NewBMI(h float32, w float32) *BMI {
	return &BMI{h, w}
}

func (bmi *BMI) CalcBMI() (float32, error) {

	/*TODO: if User input symbols?*/

	if bmi.Weight == 0.0 || bmi.Height == 0.0 {
		return 0.0, errors.New("Нет данных о весе или росте")

	} else if bmi.Weight < 0.0 || bmi.Height < 0.0 || bmi.Weight > 500.0 || bmi.Height > 250.0 {
		return 0.0, errors.New("Данные о росте или весе некорректны")
	} else {
		h := bmi.Height / 100 //meters
		bmiValue := bmi.Weight / (h * h)
		return bmiValue, nil
	}
}
