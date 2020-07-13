package main

import (
	"banking/deposit"
	"fmt"
	"helper"
)

const (
	ExitMsg             = "q"
	IllegalValueMessage = "Вы ввели недопустимое значение. Расчёт будет перезапущен."
	TooLowValueMessage  = "Введённое Вами значение меньше 1. Расчёт будет перезапущен."
)

func main() {
	var status int
AppLoop:
	for {
		fmt.Println("=========================")
		fmt.Println("Добро пожаловать в расчёт депозита!")
		fmt.Println("Чтобы выйти введите q")
		fmt.Println("Все вводимые числовые значения не могут быть меньше 1")
		fmt.Println("=========================")
		d := deposit.Deposit{}
	QuestionsLoop:
		for i := 0; i < 5; i++ {
			switch i {
			case 0:
				d.StartSum, status = helper.RequestInputFloat64("Введите начальную сумму вклада:", ExitMsg)
			case 1:
				d.Percent, status = helper.RequestInputFloat64("Введите годовой процент по вкладу:", ExitMsg)
			case 2:
				d.Months, status = helper.RequestInputInt("Введите количество месяцев (60 месяцев = 5 лет):", ExitMsg)
			case 3:
				d.Capitalization, status = helper.RequestInputYesNo("Есть ли капитализация (y/n):", ExitMsg)
			case 4:
				if !d.Capitalization {
					break QuestionsLoop
				}
				d.CapitalizationPeriod, status = helper.RequestInputInt("Введите период капитализации процентов (1 - каждый месяц):", ExitMsg)
			}
			switch status {
			case helper.DoTerminate:
				break AppLoop
			case helper.IllegalValue:
				fmt.Println(IllegalValueMessage)
				continue AppLoop
			default:
				switch i {
				case 0:
					if d.StartSum < 1 {
						fmt.Println(TooLowValueMessage)
						continue AppLoop
					}
				case 1:
					if d.Percent < 1 {
						fmt.Println(TooLowValueMessage)
						continue AppLoop
					}
				case 2:
					if d.Months < 1 {
						fmt.Println(TooLowValueMessage)
						continue AppLoop
					}
				case 4:
					if d.Capitalization && d.CapitalizationPeriod < 1 {
						fmt.Println(TooLowValueMessage)
						break QuestionsLoop
					}
				}
			}
		}
		d.Count()
		fmt.Println("=========================")
		for _, line := range d.GetCountResults().MonthlyList {
			fmt.Println("| Месяц:", line.Month, "| Сумма депозита:", helper.PrettyFloat64(line.DepositSum), "| Проценты:", helper.PrettyFloat64(line.PercentsSum), "|")
		}
		fmt.Println("Доход:", helper.PrettyFloat64(d.GetCountResults().PercentsSum))
		fmt.Println("Итоговая сумма вклада:", helper.PrettyFloat64(d.GetCountResults().FinalSum))
		fmt.Println("=========================")
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
