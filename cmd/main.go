package main

import (
	"context"
	"electronics-store-go/internal/app/controller"
	"electronics-store-go/internal/config"
	"electronics-store-go/internal/database"
	"electronics-store-go/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	router := gin.New()
	log := logger.NewLogger()
	cnf, err := config.LoadConfig()
	if err != nil {
		log.Error("Error with load config files!")
	}
	db, err := database.NewDatabase(&cnf)
	if err != nil {
		log.Error("Error with connect Database!")
	}

	controller.RouteV1(&controller.Handler{Db: db}, router)

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	// Graceful server shutdown - https://github.com/gin-gonic/examples/blob/master/graceful-shutdown/graceful-shutdown/server.go
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to initialize server: %v\n", err)
		}
	}()

	log.Printf("Listening on port %v\n", server.Addr)

	// Wait for kill signal of channel
	quit := make(chan os.Signal)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// This blocks until a signal is passed into the quit channel
	<-quit

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Shutdown server
	log.Println("Shutting down server...")
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v\n", err)
	}

	//newProduct := models.Product{Name: "Mahsulot Nomi", Title: "Mahsulot Sarlavhasi", Price: 13.0, Specification: "Tavsifi"}

}
