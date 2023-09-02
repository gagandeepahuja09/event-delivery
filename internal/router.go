package internal

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"

	log "github.com/sirupsen/logrus"
)

func serveRoutes(router *mux.Router) {
	srv := &http.Server{
		Addr:    "localhost:9410",
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout
	quit := make(chan os.Signal, 1)

	log.Info("Server is up...")
	signal.Notify(quit, syscall.SIGUSR1, syscall.SIGINT, syscall.SIGTERM)

	<-quit

	log.Info("Shutting down Server ...")

	time.Sleep(2 * time.Second)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown with error : ", err)
	}

	log.Info("Server exiting")
}

func InitRouter() {
	router := mux.NewRouter()

	router.HandleFunc("/v1/proxy", proxy).Methods("POST")
	serveRoutes(router)
}