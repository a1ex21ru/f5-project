package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"f5-project/cmd/server"
	"f5-project/internal/adapters/postgres"
	"f5-project/internal/handlers"
	"f5-project/internal/repository"
	"f5-project/internal/routes"
	"f5-project/internal/service"
)

func main() {

	ctx := context.Background()

	store, err := postgres.NewClient(ctx)
	if err != nil {
		log.Println(ctx, "Failed to connect to database", "err", err.Error())
		os.Exit(1)
	}

	repo := repository.NewRepository(store.GetGormConn())
	serv := service.NewService(repo)
	handler := handlers.NewHandler(serv)
	router := routes.NewHandler(handler)

	srv := new(server.Server)
	go func() {
		if err = srv.Run(os.Getenv("port"), router.InitRoutes()); err != nil {
			log.Println(ctx, "Failed to start server", "err", err.Error())
		}
	}()

	log.Println(`Server started on your http://localhost:8080`)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	if err = srv.Shutdown(ctx); err != nil {
		log.Println(ctx, "error server shutting", "err", err.Error())
	}

	if err = store.Close(); err != nil {
		log.Println(ctx, "error db connect close", "err", err.Error())
	}

}
