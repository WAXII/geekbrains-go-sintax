package main

import (
	"fmt"
	"geometry/triangle"
	"helper"
)

func main() {
	fmt.Println("=========================")
	fmt.Println("Чтобы выйти введите q")
	fmt.Println("=========================")
AppLoop:
	for {
		a, status := helper.RequestInputFloat64("Введите катет A прямоугольного треугольника: ", "q")
		switch status {
		case helper.DoTerminate:
			break AppLoop
		case helper.IllegalValue:
			fmt.Println("Вы ввели недопустимое значение. Расчёт будет перезапущен.")
			continue AppLoop
		}
		b, status := helper.RequestInputFloat64("Введите катет B прямоугольного треугольника: ", "q")
		switch status {
		case helper.DoTerminate:
			break AppLoop
		case helper.IllegalValue:
			fmt.Println("Вы ввели недопустимое значение. Расчёт будет перезапущен.")
			continue AppLoop
		}
		t := triangle.Triangle{
			SideA: a,
			SideB: b,
		}
		t.GetHipotinuse()
		fmt.Println("Площадь:", t.GetScuare())
		fmt.Println("Гипотенуза:", t.Hipotinuse)
		fmt.Println("Периметр:", t.SumSides())
		relaunch, status := helper.RequestInputYesNo("Перезапустить расчёт? (y/n)", "q")
		switch status {
		case helper.DoTerminate:
			break AppLoop
		case helper.IllegalValue:
			continue AppLoop
		default:
			if !relaunch {
				break AppLoop
			}
		}
	}
}
