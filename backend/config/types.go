package config

import (
	"encoding/hex"
	"strings"
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

	CtxConfig struct {
		Key string
	}

	Config struct {
		DebugMode      bool      `mapstructure:"-"`
		ListenAddr     string    `mapstructure:"listen-addr"`
		TLSCertFile    string    `mapstructure:"tls-cert-file"`
		TLSKeyFile     string    `mapstructure:"tls-key-file"`
		AuthCookieName string    `mapstructure:"auth-cookie-name"`
		DocxDir        string    `mapstructure:"docx-dir"`
		JWT            JWTConfig `mapstructure:"jwt-token"`
		DB             DBConfig  `mapstructure:"db"`
		Ctx            CtxConfig
	}
)

func (c *Config) init() {
	if c.AuthCookieName == "" {
		c.AuthCookieName = "X-App-Auth"
	}
	if c.DocxDir == "" {
		c.DocxDir = "./docx/"
	}
	if !strings.HasSuffix(c.DocxDir, "/") {
		c.DocxDir += "/"
	}
	c.Ctx.Key = "user-context"
	var err error
	c.JWT.SignKey, err = hex.DecodeString(c.JWT.SignKeyHex)
	if err != nil {
		panic(err)
	}
}
