package bmi

import (
	"fmt"
	"strconv"
	"strings"
)

func Calculate() {
	fmt.Println(mainTitle)
	fmt.Println(separator)
	fmt.Print(weightPrompt)
	weightInput, weightInputErr := reader.ReadString('\n')
	if weightInputErr != nil {
		return
	}
	fmt.Print(heightPrompt)
	heightInput, heightInputErr := reader.ReadString('\n')
	if heightInputErr != nil {
		return
	}
	// Sanitize userInputs
	weightInput = strings.Replace(weightInput, "\n", "", -1)
	heightInput = strings.Replace(heightInput, "\n", "", -1)

	// Parse userInputs
	weight, weightErr := strconv.ParseFloat(weightInput, 64)
	if weightErr != nil {
		return
	}
	height, heightErr := strconv.ParseFloat(heightInput, 64)
	if heightErr != nil {
		return
	}
	bmi := weight / (height * height)

	fmt.Println(separator)
	fmt.Printf("Your BMI: %.2f", bmi)

}
