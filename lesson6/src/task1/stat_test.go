package task1

import "testing"

type StatCase struct {
	Name            string
	Input           []float64
	ExpectedAverage float64
	ExpectedSum     float64
}

func getTestCases() []StatCase {
	return []StatCase{
		{
			Name:            "AverageCase1",
			Input:           []float64{4, 4, 4},
			ExpectedAverage: 4,
			ExpectedSum:     12,
		},
		{
			Name:            "AverageCase2",
			Input:           []float64{1, 2, 3},
			ExpectedAverage: 2,
			ExpectedSum:     6,
		},
	}
}

func TestAverage(t *testing.T) {
	for _, c := range getTestCases() {
		result := Average(c.Input)
		if result != c.ExpectedAverage {
			t.Errorf("Failed %s! Expexted: %2f, got %2f", c.Name, c.ExpectedAverage, result)
		}
	}
}

func TestSumList(t *testing.T) {
	for _, c := range getTestCases() {
		result := SumList(c.Input)
		if result != c.ExpectedSum {
			t.Errorf("Failed %s! Expexted: %2f, got %2f", c.Name, c.ExpectedSum, result)
		}
	}
}
