package config

type Server struct {
	MySQL  MySQL
	Redis  Redis
	App    App
	Jwt    Jwt
	Logger Logger
}
