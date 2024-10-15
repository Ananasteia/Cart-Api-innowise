package main

import (
	"Cart_Api_New/config"
	"Cart_Api_New/internal/database"
	"Cart_Api_New/internal/handlers"
	"Cart_Api_New/internal/repositories"
	"Cart_Api_New/internal/services"
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os/signal"
	"syscall"
)

var (
	cfgFile = flag.String("cfg", "./config/config.yml", "path to config file")
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(),
		syscall.SIGHUP, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGABRT, syscall.SIGTERM)
	defer cancel()

	cfg := config.ReadConfig(*cfgFile)

	db, err := database.New(ctx, cfg.DBConfig)

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close() // in database.New -???????

	newRepository := repositories.New(db)
	newApp := services.New(newRepository)
	newApi := handlers.New(newApp)
	serverHTTP := http.Server{
		Addr:    net.JoinHostPort(cfg.Server.Host, cfg.Server.Port),
		Handler: newApi,
	}

	fmt.Printf("server is running on: %s:%s", cfg.Server.Host, cfg.Server.Port)
	log.Fatal(serverHTTP.ListenAndServe())
}
