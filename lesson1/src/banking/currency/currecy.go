package currency

import "math"

const rate = 70.89

func ConvertRurToUsd(rur float64) (float64, float64) {
	usd := rur / rate
	return humanizeCurrency(rur), humanizeCurrency(usd)
}

func humanizeCurrency(value float64) float64 {
	return math.Round(value*100) / 100
}
