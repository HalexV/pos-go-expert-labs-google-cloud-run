package dto

type GetBrasilAPIOutput struct {
	Cep          string   `json:"cep"`
	State        string   `json:"state"`
	City         string   `json:"city"`
	Neighborhood string   `json:"neighborhood"`
	Street       string   `json:"street"`
	Location     Location `json:"location"`
}

type Coordinates struct {
	Longitude string `json:"longitude"`
	Latitude  string `json:"latitude"`
}

type Location struct {
	Type        string      `json:"type"`
	Coordinates Coordinates `json:"coordinates"`
}

type GetWeatherAPIOutput struct {
	Current Current `json:"current"`
}
type Current struct {
	TempC float64 `json:"temp_c"`
	TempF float64 `json:"temp_f"`
}

type APIOutput struct {
	TempC float64 `json:"temp_C"`
	TempF float64 `json:"temp_F"`
	TempK float64 `json:"temp_K"`
}
