package config

import (
	"flag"
	"fmt"
	"log"
	"os"
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
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "\n")
		fmt.Fprintf(os.Stderr, "Please, make sure you have the configuration files and use flags to setup its.\n")
		fmt.Fprintf(os.Stderr, "Flags:\n")
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\n")
	}
	flag.StringVar(&cfgPath, "cfg-dir", defPath, "dir with app's config")
	flag.StringVar(&cfgName, "cfg-name", defName, "app's config base file name")
	flag.BoolVar(&debugMode, "dbg", false, "debug mode")
	flag.Parse()
}

func Get() Config {
	once.Do(func() {
		if err := func() error {
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
				return fmt.Errorf("Fatal error config file: %s \n", err)
			}
			err = viper.Unmarshal(&cfg)
			if err != nil {
				return err
			}
			cfg.DebugMode = debugMode
			cfg.init()

			if cfg.DebugMode {
				log.Printf("Loaded config: %#v\n", cfg)
			}
			return err
		}(); err != nil {
			flag.Usage()
			panic(err)
		}
	})
	return cfg
}
