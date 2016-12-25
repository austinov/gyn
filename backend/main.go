package main

import (
	"flag"
	"log"

	"github.com/austinov/gyn/backend/config"
	"github.com/austinov/gyn/backend/route"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine"
	"github.com/labstack/echo/engine/standard"
)

var (
	cfgPath, cfgName string
	debugMode        bool
)

func main() {
	flag.StringVar(&cfgPath, "cfg-dir", "", "dir with app's config")
	flag.StringVar(&cfgName, "cfg-name", "", "app's config base file name")
	flag.BoolVar(&debugMode, "dbg", false, "debug mode")
	flag.Parse()
	config.Init(cfgPath, cfgName)
	c := config.Get()

	e := echo.New()
	if debugMode {
		e.SetDebug(true)
		e.SetLogLevel(0)
	}

	route.Init(e)

	log.Printf("Serving at address: '%s'.", c.ListenAddr)
	log.Printf("Press Ctrl+C to exit.")

	e.Run(standard.WithConfig(engine.Config{
		Address: c.ListenAddr,
	}))
}
