package models

type Starships struct {
	Name                   string      `json:"name"`
	Model                  string      `json:"model"`
	Length                 string      `json:"length"`
	Max_atmosphering_speed string      `json:"max_atmosphering_speed"`
	Passengers             string      `json:"passengers"`
	Cargo_capacity         string      `json:"cargo_capacity"`
	Pilots                 interface{} `json:"pilots"`
	Films                  interface{} `json:"films"`
}
