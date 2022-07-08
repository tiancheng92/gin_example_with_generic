package config

type Config struct {
	Mysql    string
	LogLevel string
	I18n     string
	Server   struct {
		Mode        string
		ServicePort int
		ServiceHost string
	}
}
