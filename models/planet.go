package models

type Planet struct {
	Name            string      `json:"name"`
	Rotation_period string      `json:"rotation_period"`
	Orbital_period  string      `json:"orbital_period"`
	Diameter        string      `json:"diameter"`
	Climate         string      `json:"climate"`
	Gravity         string      `json:"gravity"`
	Population      string      `json:"population"`
	Residents       string      `json:"residents"`
	Films           interface{} `json:"films"`
}
