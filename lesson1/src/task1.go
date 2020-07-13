package main

import (
	"banking/currency"
	"fmt"
	"helper"
)

func main() {
	fmt.Println("=========================")
	fmt.Println("Добро пожаловать в конвертор валют RUR > USD!")
	fmt.Println("Чтобы выйти введите q")
	fmt.Println("=========================")
AppLoop:
	for {
		rurFloat, status := helper.RequestInputFloat64("Введите конвертируемую сумму в рублях: ", "q")
		switch status {
		case helper.DoTerminate:
			break AppLoop
		case helper.IllegalValue:
			fmt.Println("Вы ввели некорректное значение. Попробуйюе снова или введите q для выхода")
			continue AppLoop
		}
		rur, usd := currency.ConvertRurToUsd(rurFloat)
		fmt.Println(">>>", rur, "RUR =", usd, "USD")
	}

}
