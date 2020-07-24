package vehicles

import "encoding/json"

const (
	VECICLE_TYPE_CAR   = "car"
	VECICLE_TYPE_TRUCK = "truck"
)

type vecicleError struct {
	s string
}

func (e *vecicleError) Error() string {
	return e.s
}

type VehicleInterface interface {
	openWindows()
	closeWindows()
	startEngine()
	stopEngine()
}

type vecicleInfo struct {
	Alias       string  `json:"vechicleAlias"`
	Name        string  `json:"vechicleName"`
	VType       string  `json:"vechicleType"`
	Year        int     `json:"vechicleYear"`
	TrunkVolume float64 `json:"vechicleTruckVolume"`
}

type vehiclestate struct {
	EngineStarted bool    `json:"engineStarted"`
	WindowsOpened bool    `json:"windowOpened"`
	TrunkFill     float64 `json:"trunkFill"`
}

var vehiclesList map[string]VehicleInterface

func AddCar(alias, name string, year int, trunkVol float64) error {
	if len(vehiclesList) == 0 {
		vehiclesList = make(map[string]VehicleInterface)
	} else if _, exists := vehiclesList[alias]; exists {
		return &vecicleError{s: "Vecicle already exists"}
	}
	vehiclesList[alias] = &car{
		Info: &vecicleInfo{
			Alias:       alias,
			Name:        name,
			VType:       VECICLE_TYPE_CAR,
			Year:        year,
			TrunkVolume: trunkVol,
		},
		State: &vehiclestate{},
	}
	return nil
}

func AddTruck(alias, name string, year int, trunkVol float64) error {
	if len(vehiclesList) == 0 {
		vehiclesList = make(map[string]VehicleInterface)
	} else if _, exists := vehiclesList[alias]; exists {
		return &vecicleError{s: "Vecicle already exists"}
	}
	vehiclesList[alias] = &truck{
		Info: &vecicleInfo{
			Alias:       alias,
			Name:        name,
			VType:       VECICLE_TYPE_TRUCK,
			Year:        year,
			TrunkVolume: trunkVol,
		},
		State: &vehiclestate{},
	}
	return nil
}

func GetAllVehicles() string {
	if len(vehiclesList) == 0 {
		return "[]"
	}
	j, _ := json.Marshal(vehiclesList)
	return string(j)
}

func OpenCloseWindow(vehicleAlias string, open bool) error {
	v, e := getVehicle(vehicleAlias)
	if e != nil {
		return e
	}
	if open {
		v.openWindows()
	} else {
		v.closeWindows()
	}
	return nil
}

func StartStopEngine(vehicleAlias string, start bool) error {
	v, e := getVehicle(vehicleAlias)
	if e != nil {
		return e
	}
	if start {
		v.startEngine()
	} else {
		v.stopEngine()
	}
	return nil
}

func getVehicle(alias string) (v VehicleInterface, err error) {
	v, e := vehiclesList[alias]
	if !e {
		return nil, &vecicleError{s: "Vecicle already exists"}
	}
	return v, nil
}
