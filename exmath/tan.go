package exmath

import (
	"fmt"
	"math"
)

func TanExample(degree float64) {
	a := math.Round(math.Tan(DegreeToRadian(degree)))
	fmt.Println("Tan 45度: ", a)
	fmt.Println("ATan 45度: ", RadianToDegree(math.Atan(a)))

	fmt.Println(math.Tan(DegreeToRadian(45)))
	fmt.Println(math.Atan(1))
	fmt.Println(DegreeToRadian(90))
	fmt.Println(DegreeToRadian(0))
}

// RadianToDegree 弧度转角度
func RadianToDegree(x float64) float64 {
	return x * (180 / math.Pi)
}

// DegreeToRadian 角度转弧度
func DegreeToRadian(x float64) float64 {
	return x * (math.Pi / 180)
}
