package config

import (
	"flag"
	"fmt"
	"log"
	"sync"

	"github.com/spf13/viper"
)

const (
	defPath = "./env"
	defName = "dev"
)

var (
	cfg       Config
	cfgPath   string
	cfgName   string
	debugMode bool
	once      sync.Once
)

func init() {
	flag.StringVar(&cfgPath, "cfg-dir", "", "dir with app's config")
	flag.StringVar(&cfgName, "cfg-name", "", "app's config base file name")
	flag.BoolVar(&debugMode, "dbg", false, "debug mode")
	flag.Parse()
}

func Get() Config {
	once.Do(func() {
		if cfgPath != "" {
			viper.AddConfigPath(cfgPath)
		}
		viper.AddConfigPath(defPath)
		viper.AddConfigPath(".")
		if cfgName == "" {
			cfgName = defName
		}
		viper.SetConfigName(cfgName)

		err := viper.ReadInConfig()
		if err != nil {
			panic(fmt.Errorf("Fatal error config file: %s \n", err))
		}
		err = viper.Unmarshal(&cfg)
		if err != nil {
			panic(err)
		}
		cfg.DebugMode = debugMode
		cfg.init()

		if cfg.DebugMode {
			log.Printf("Loaded config: %#v\n", cfg)
		}
	})
	return cfg
}
