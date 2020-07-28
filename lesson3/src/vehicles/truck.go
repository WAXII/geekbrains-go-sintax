package vehicles

type truck struct {
	Info  *vecicleInfo       `json:"vecicleInfo"`
	State *vehiclestate      `json:"vehiclestate"`
	Trunc map[string]float64 `json:"boxesInTrunc"`
}

func (v *truck) openWindows() {
	v.State.WindowsOpened = true
}

func (v *truck) closeWindows() {
	v.State.WindowsOpened = false
}

func (v *truck) startEngine() {
	v.State.EngineStarted = true
}

func (v *truck) stopEngine() {
	v.State.EngineStarted = false
}
