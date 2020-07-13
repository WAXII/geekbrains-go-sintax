package helper

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

const (
	DoTerminate  = 0
	IllegalValue = 1
	IsOK         = 2
)

func RequestString(msg string, toLower bool) string {
	var input string
	fmt.Print(msg + " ")
	fmt.Scanln(&input)
	if toLower {
		input = strings.ToLower(input)
	}
	return input
}

func PrettyFloat64(v float64) float64 {
	return math.Round(v*100) / 100
}

func StringToFloat64(str string) (float64, error) {
	str = strings.Replace(str, ",", ".", 1)
	return strconv.ParseFloat(str, 64)
}

func RequestInputFloat64(msg string, exitMsg string) (float64, int) {
	input := RequestString(msg, true)
	if input == exitMsg {
		return 0, DoTerminate
	}
	result, err := StringToFloat64(input)
	if err != nil {
		return 0, IllegalValue
	}
	return result, IsOK
}

func RequestInputInt(msg string, exitMsg string) (int, int) {
	input := RequestString(msg, true)
	if input == exitMsg {
		return 0, DoTerminate
	}
	result, err := strconv.Atoi(input)
	if err != nil {
		return 0, IllegalValue
	}
	return result, IsOK
}

func RequestInputYesNo(msg string, exitMsg string) (bool, int) {
	input := RequestString(msg, true)
	if input == exitMsg {
		return false, DoTerminate
	}
	if input != "y" && input != "n" {
		return false, IllegalValue
	}
	if input == "y" {
		return true, IsOK
	}
	return false, IsOK
}
