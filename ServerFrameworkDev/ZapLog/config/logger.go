package config

type Logger struct {
	Level    string `json:"level"`
	Director string `json:"director"`
	Layout   string `json:"layout"`
}
