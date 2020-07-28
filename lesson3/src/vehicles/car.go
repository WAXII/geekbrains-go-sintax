package vehicles

type car struct {
	Info  *vecicleInfo       `json:"vecicleInfo"`
	State *vehiclestate      `json:"vehiclestate"`
	Trunc map[string]float64 `json:"boxesInTrunc"`
}

func (v *car) openWindows() {
	v.State.WindowsOpened = true
}

func (v *car) closeWindows() {
	v.State.WindowsOpened = false
}

func (v *car) startEngine() {
	v.State.EngineStarted = true
}

func (v *car) stopEngine() {
	v.State.EngineStarted = false
}
