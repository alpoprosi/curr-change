package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/alpoprosi/curr-change/internal/db"
	"github.com/alpoprosi/curr-change/internal/handlers"

	"github.com/AdguardTeam/golibs/log"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const (
	pathBTCUSDT = "/btcusdt"
	pathCurrs   = "/currencies"
	pathLatest  = "/latest"
)

func main() {
	cfg, err := parseConfig()
	if err != nil {
		log.Fatalf("[FATAL] getting config: %v", err)
	}

	log.SetLevel(log.Level(cfg.LogLevel))

	srv := echo.New()
	sg := srv.Group("/api")

	srv.HideBanner = true
	srv.HidePort = true

	srv.Use(middleware.Logger())
	srv.Use(middleware.RemoveTrailingSlashWithConfig(middleware.TrailingSlashConfig{
		RedirectCode: http.StatusMovedPermanently,
	}))

	db, err := db.NewDB(cfg.DBConfig)
	if err != nil {
		log.Fatal(err)
	}

	h := handlers.NewHandler(db)

	sg.Add(string(http.MethodGet), pathBTCUSDT, h.BTCUSDTLast)
	sg.Add(string(http.MethodPost), pathBTCUSDT, h.BTCUSDTHistory)

	sg.Add(string(http.MethodGet), pathBTCUSDT, h.CurrsLast)
	sg.Add(string(http.MethodPost), pathBTCUSDT, h.CurrsHistory)

	sg.Add(string(http.MethodGet), pathBTCUSDT, h.BTCCurrLast)

	addr := fmt.Sprintf("%s:%s", cfg.Addr, cfg.Port)
	log.Fatalf("Starting server: %v", srv.Start(addr))

	graceShutdown(context.Background(), srv)
}

func graceShutdown(c context.Context, srv *echo.Echo) {
	sigterm := make(chan os.Signal)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		for {
			sig := <-sigterm
			fmt.Printf("[INFO] received signal %q", sig)
			srv.Shutdown(c)
			os.Exit(0)
		}
	}()
}
