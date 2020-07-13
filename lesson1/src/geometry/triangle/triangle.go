package triangle

import (
	"math"
)

type Triangle struct {
	SideA      float64
	SideB      float64
	Hipotinuse float64
}

func (t *Triangle) GetScuare() float64 {
	return t.SideA * t.SideB / 2
}

func (t *Triangle) GetHipotinuse() {
	t.Hipotinuse = math.Sqrt(t.SideA*t.SideA + t.SideB*t.SideB)
}

func (t *Triangle) SumSides() float64 {
	return t.SideA + t.SideB + t.Hipotinuse
}
