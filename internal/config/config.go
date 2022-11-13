package config

type Config struct {
	Address string
}

func New() *Config {
	return &Config{
		Address: ":3333",
	}
}
