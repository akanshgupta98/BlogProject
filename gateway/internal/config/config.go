package config

func New() Config {

	return Config{
		Srv: newServerCfg(),
	}
}
func newServerCfg() ServerConfig {

	srv := ServerConfig{
		IP:   "127.0.0.1",
		Port: "8080",
	}
	return srv

}
