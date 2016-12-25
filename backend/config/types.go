package config

const (
	defPath = "./env"
	defName = "dev"
)

type Config struct {
	ListenAddr string   `mapstructure:"listen-addr"`
	DB         DBConfig `mapstructure:"db"`
}

func (c *Config) init() {
}

type DBConfig struct {
}
