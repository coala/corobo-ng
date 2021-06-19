package config

type DatabaseConfiguration struct {
	Port     string
	Host     string
	Name     string
	User     string
	Password string
	Driver   string
	LogMode  bool
}
