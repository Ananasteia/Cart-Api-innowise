package main

import (
	"Cart_Api_New/config"
	"Cart_Api_New/internal/database"
	"Cart_Api_New/internal/handlers"
	"Cart_Api_New/internal/repositories"
	"Cart_Api_New/internal/services"
	"context"
	"flag"
	"log"
	"net"
	"net/http"
	"os/signal"
	"syscall"
)

var (
	cfgFile = flag.String("cfg", "/build/config.yml", "path to config file")
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(),
		syscall.SIGHUP, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGABRT, syscall.SIGTERM)
	defer cancel()

	cfg, err := config.ReadConfig(*cfgFile)
	if err != nil {
		log.Fatal(err)
	}

	db, err := database.New(ctx, cfg.DBConfig)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	newRepository := repositories.New(db)
	newService := services.New(newRepository)
	newHandler := handlers.New(newService)
	newServer := server(newHandler, cfg.Server)

	log.Printf("server is running on: %s:%s", cfg.Server.Host, cfg.Server.Port)
	if err := newServer.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

func server(handler handlers.Handler, cfg config.Server) *http.Server {
	return &http.Server{
		Addr:    net.JoinHostPort(cfg.Host, cfg.Port),
		Handler: handler.Handle(),
	}
}
