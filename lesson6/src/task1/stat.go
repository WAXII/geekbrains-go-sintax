package task1

func Average(xs []float64) float64 {
	return SumList(xs) / float64(len(xs))
}

func SumList(xs []float64) float64 {
	total := float64(0)
	for _, x := range xs {
		total += x
	}
	return total
}
