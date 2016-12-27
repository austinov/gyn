package config

import (
	"encoding/hex"
	"time"
)

type (
	JWTConfig struct {
		Issuer     string        `mapstructure:"issuer"`
		SignKeyHex string        `mapstructure:"sign-key"`
		Expiration time.Duration `mapstructure:"expiration"`
		SignKey    []byte
	}

	DBConfig struct {
		Type             string `mapstructure:"type"`
		ConnectionString string `mapstructure:"connection-string"`
	}

	Config struct {
		DebugMode      bool      `mapstructure:"-"`
		ListenAddr     string    `mapstructure:"listen-addr"`
		AuthCookieName string    `mapstructure:"auth-cookie-name"`
		JWT            JWTConfig `mapstructure:"jwt-token"`
		DB             DBConfig  `mapstructure:"db"`
	}
)

func (c *Config) init() {
	if c.AuthCookieName == "" {
		c.AuthCookieName = "X-App-Auth"
	}
	var err error
	c.JWT.SignKey, err = hex.DecodeString(c.JWT.SignKeyHex)
	if err != nil {
		panic(err)
	}
}
