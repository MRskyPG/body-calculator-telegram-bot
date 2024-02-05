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
