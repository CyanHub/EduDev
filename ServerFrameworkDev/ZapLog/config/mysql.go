package config

type MySQL struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"username"`
	Password string `json:"password"`
	Database string `json:"database"`
	Config   string `json:"config"`
}
