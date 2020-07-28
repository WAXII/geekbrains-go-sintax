package main

import (
	"fmt"
	"task3"
)

func main() {
	input := ""
MainLoop:
	for {
		fmt.Print("> ")
		if _, err := fmt.Scanln(&input); err != nil {
			fmt.Println(err)
			continue MainLoop
		}
		switch input {
		case "exit":
			break MainLoop
		case "help":
			fmt.Println(task3.ShowHelp())
			continue MainLoop
		default:
			if res, err := task3.Calculate(input); err == nil {
				fmt.Printf("Результат: %v\n", res)
			} else {
				fmt.Println("Не удалось произвести вычисление")
			}
		}
	}
}
