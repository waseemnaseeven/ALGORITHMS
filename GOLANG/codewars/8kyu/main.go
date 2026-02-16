package main

import (
	"fmt"
)

func BinaryToDecimal(bin string) int {
	result := 0
	for _,digit := range bin {
		result = result * 2 // decalage a gauche
		if digit == '1' {
			result++ // ajouter 1 si le bit est a 1
		}
	}
	return result
}


func guessBlue(blueStart, redStart, bluePulled, redPulled int) float64 {
	blueRemain := blueStart - bluePulled
	redRemain := redStart - redPulled
	totalRemain := blueRemain + redRemain

	if totalRemain == 0 {
		return 0.0
	}

	ret := float64(blueRemain) / float64(totalRemain)

	return ret

}

func upperCaseStr(str string) string {
	b := []byte(str)

	for i := range b {
		if b[i] >= 'a' && b[i] <= 'z' {
			b[i] -= 32
		}
	}
	return string(b)
}

func main() {
	fmt.Println(BinaryToDecimal("0110"))
	fmt.Println(guessBlue(5,5,2,3))
	fmt.Println(upperCaseStr("hello world"))
}