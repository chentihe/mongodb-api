package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/chentihe/mongodb-api/config"
	"github.com/chentihe/mongodb-api/config/svc"
	_ "github.com/chentihe/mongodb-api/docs"
	"github.com/chentihe/mongodb-api/routes"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	config, err := config.LoadConfig(".", "../")
	if err != nil {
		log.Fatal("Could not load environment variables", err)
	}

	ctx := context.TODO()
	svc, err := svc.NewServiceContext(config, ctx)
	if err != nil {
		log.Fatal("Could not initial the service context", err)
	}

	router := gin.Default()
	routes.RegisterRouters(router, svc)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	defer svc.DB.Client().Disconnect(ctx)

	srv := &http.Server{
		Addr:    ":" + config.Port,
		Handler: router,
	}
	GracefulShutdown(srv)
}

func GracefulShutdown(server *http.Server) {
	go func() {
		// service connections
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds
	quit := make(chan os.Signal)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	//catching ctx.Done(). timeout of 5 seconds
	select {
	case <-ctx.Done():
		log.Println("timeout of 5 seconds")
	}
	log.Println("Server exiting")
}
