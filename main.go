package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"goAdminStudy/config"
	"goAdminStudy/route"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {

	gin.SetMode(gin.DebugMode)

	engine := route.InitRoute()

	server := &http.Server{
		Addr:    config.ServerConfig.GetHost() + ":" + string(config.ServerConfig.GetPort()),
		Handler: engine,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Shutdown...", err)
	}
	log.Println("Server exit...")

}
