package config

type Config struct {
	Mysql  Mysql
	Log    Log
	Server Server
	I18n   I18n
}

type Mysql struct {
	DBHost     string
	DBPort     string
	DBName     string
	DBUser     string
	DBPassword string
}

type Server struct {
	Mode        string
	ServicePort int
	ServiceHost string
}

type Log struct {
	Level string
}

type I18n struct {
	Locale string
}
