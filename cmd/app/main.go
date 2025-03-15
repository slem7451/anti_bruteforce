package main

import (
	"context"
	"fmt"
	"log"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/slem7451/anti_bruteforce/internal/app"
	"github.com/slem7451/anti_bruteforce/internal/pgsql"
	"github.com/slem7451/anti_bruteforce/internal/redis"
	"github.com/slem7451/anti_bruteforce/internal/request"
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

	for i := 0; i < 11; i++ {
		fmt.Println(app.ValidateAuth(ctx, request.Credits{IP: "1921", Login: "lll1", Password: "pswd1"}))
	}

	fmt.Println(app.ValidateAuth(ctx, request.Credits{IP: "1921", Login: "lll2", Password: "pswd2"}))
	fmt.Println(app.ValidateAuth(ctx, request.Credits{IP: "1922", Login: "lll1", Password: "pswd1"}))
	fmt.Println(app.ValidateAuth(ctx, request.Credits{IP: "1922", Login: "lll2", Password: "pswd1"}))
}
