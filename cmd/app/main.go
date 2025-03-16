package main

import (
	"context"
	"log"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/slem7451/anti_bruteforce/internal/app"
	"github.com/slem7451/anti_bruteforce/internal/server/grpc"
	"github.com/slem7451/anti_bruteforce/internal/storage/pgsql"
	"github.com/slem7451/anti_bruteforce/internal/storage/redis"
)

func init() {
	err := godotenv.Load("configs/.app.env")

	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(),
		syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer cancel()

	redisC, err := redis.NewClient(ctx)

	if err != nil {
		log.Fatalf("Redis: error - %s", err)
	}

	pgC, err := pgsql.NewClient(ctx)

	if err != nil {
		log.Fatalf("PostgreSQL: error - %s", err)
	}

	app, err := app.NewApp(redisC, pgC)

	if err != nil {
		log.Fatalf("App: error - %s", err)
	}

	server := grpc.NewServer(app)

	go func () {
		if err := server.Start(ctx); err != nil {
			log.Fatalf("Server: start error - %s", err)
		}
	}()

	go func () {
		<-ctx.Done()

		if err := server.Stop(ctx); err != nil {
			log.Printf("Server: stop error - %s", err)
		}
	}()

	log.Println("App is running...")

	<-ctx.Done()
}
