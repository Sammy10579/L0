package main

import (
	"log"
	"os"

	"L0/app"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}

	// todo ensure variables exist

	application := app.NewApplication(app.Config{
		ListenedAddr: os.Getenv("HTTP_LISTENED_ADDR"),
		DBConnection: os.Getenv("DATABASE_URL"),
		Nats: struct {
			URL       string
			ClusterID string
			ClientID  string
		}{
			URL:       os.Getenv("NATS_STREAMING_URL"),
			ClusterID: os.Getenv("STAN_CLUSTER_ID"),
			ClientID:  os.Getenv("STAN_CLIENT_ID"),
		},
	})

	application.Run()
}
