package config

type Config struct {
	Srv ServerConfig
}

type ServerConfig struct {
	IP   string
	Port string
}
