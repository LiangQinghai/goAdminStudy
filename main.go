package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"goAdminStudy/config"
	_ "goAdminStudy/docs"
	"goAdminStudy/route"
	"goAdminStudy/validator"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

// @title api
// @version 1.0
// @description admin
//
func main() {

	gin.SetMode(gin.DebugMode)
	binding.Validator = new(validator.NewValidator)

	engine := route.InitRoute()

	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", config.ServerConfig.GetHost(), config.ServerConfig.GetPort()),
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
