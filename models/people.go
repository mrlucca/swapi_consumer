package models

type People struct {
	Name          string      `json:"name"`
	Gender        string      `json:"gender"`
	Participation interface{} `json:"films"`
	// SpecieInfo    SpecieInfo
}
