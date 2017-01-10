package main

import (
	"log"

	"github.com/austinov/gyn/backend/config"
	"github.com/austinov/gyn/backend/handler"
	"github.com/austinov/gyn/backend/route"
	"github.com/austinov/gyn/backend/store"
	"github.com/austinov/gyn/backend/store/pg"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine"
	"github.com/labstack/echo/engine/standard"
)

func main() {
	c := config.Get()

	e := echo.New()
	if c.DebugMode {
		e.SetDebug(true)
		e.SetLogLevel(0)
	}

	dao := createDao(c.DB)
	defer dao.Close()

	h := handler.New(c, dao, handler.NewErrorCustomizer())

	route.Init(e, h)

	log.Printf("Serving at address: '%s'.", c.ListenAddr)
	if c.TLSKeyFile != "" && c.TLSCertFile != "" {
		log.Printf("Use 'https://' prefix in browser.")
	}
	log.Printf("Press Ctrl+C to exit.")

	e.Run(standard.WithConfig(engine.Config{
		Address:     c.ListenAddr,
		TLSCertFile: c.TLSCertFile,
		TLSKeyFile:  c.TLSKeyFile,
	}))
}

func createDao(cfg config.DBConfig) store.Dao {
	switch cfg.Type {
	case "pg":
		return pg.New(cfg)
	}
	log.Fatal("Unknown db type " + cfg.Type)
	return nil
}
