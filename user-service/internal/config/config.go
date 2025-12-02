package config

type Config struct {
	Grpc GRPC_Cfg
}

type GRPC_Cfg struct {
	IP   string
	Port string
}

func NewConfig() Config {
	return Config{
		Grpc: GRPC_Cfg{
			IP:   ":",
			Port: "50001",
		},
	}
}
