package main

import (
	"context"
	"fmt"
	"log"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/slem7451/anti_bruteforce/internal/app"
	"github.com/slem7451/anti_bruteforce/internal/entity/request"
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

	/*fmt.Println(app.AddToBlacklist(ctx, "192.1.1.1/32"))
	fmt.Println(app.AddToWhitelist(ctx, "192.1.1.2/32"))*/
	/*fmt.Println(app.DeleteFromWhitelist(ctx, "192.1.1.2/32"))
	fmt.Println(app.DeleteFromBlacklist(ctx, "192.1.1.1/32"))*/

	/*fmt.Println(app.ValidateAuth(ctx, request.Credits{IP: "192.1.1.1", Login: "lll1", Password: "pswd1"}))
	fmt.Println(app.ValidateAuth(ctx, request.Credits{IP: "192.1.1.2", Login: "lll2", Password: "pswd2"}))*/

	for i := 0; i < 11; i++ {
		fmt.Println(app.ValidateAuth(ctx, request.Credits{IP: "192.1.1.3", Login: "lll1", Password: "pswd1"}))
	}

	fmt.Println(app.ValidateAuth(ctx, request.Credits{IP: "192.1.1.3", Login: "lll2", Password: "pswd2"}))
	fmt.Println(app.ValidateAuth(ctx, request.Credits{IP: "192.1.1.4", Login: "lll1", Password: "pswd1"}))
	fmt.Println(app.ValidateAuth(ctx, request.Credits{IP: "192.1.1.4", Login: "lll2", Password: "pswd1"}))
}
