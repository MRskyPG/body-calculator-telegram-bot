package bodycalc

type UserState int

type BMI struct {
	Height float64 //cm
	Weight float64 //kg
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
