package tests

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/slem7451/anti_bruteforce/internal/server/grpc/pb"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestABLogic(t *testing.T) { //nolint:funlen
	err := godotenv.Load("../configs/.app.env")
	require.NoError(t, err)

	ipLim, err := strconv.Atoi(os.Getenv("MAX_IP"))
	require.NoError(t, err)

	passwordLim, err := strconv.Atoi(os.Getenv("MAX_PASSWORD"))
	require.NoError(t, err)

	loginLim, err := strconv.Atoi(os.Getenv("MAX_LOGIN"))
	require.NoError(t, err)

	ttl, err := strconv.Atoi(os.Getenv("REDIS_TTL"))
	require.NoError(t, err)

	ctx, cancel := context.WithTimeout(context.Background(), 90*time.Second)
	defer cancel()

	conn, err := grpc.NewClient("anti-bruteforce:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	require.NoError(t, err)

	defer conn.Close()

	client := pb.NewAuthClient(conn)

	blackSubnet := "192.168.1.1/25"
	whiteSubnet := "192.168.1.192/32"

	res, err := client.AddToBlacklist(ctx, &pb.Subnet{Subnet: blackSubnet})
	require.NoError(t, err)
	require.Equal(t, true, res.Ok)

	res, err = client.AddToWhitelist(ctx, &pb.Subnet{Subnet: whiteSubnet})
	require.NoError(t, err)
	require.Equal(t, true, res.Ok)

	res, err = client.AddToBlacklist(ctx, &pb.Subnet{Subnet: "1"})
	require.NoError(t, err)
	require.Equal(t, false, res.Ok)
	require.NotEmpty(t, res.Msg)

	res, err = client.AddToWhitelist(ctx, &pb.Subnet{Subnet: "1"})
	require.NoError(t, err)
	require.Equal(t, false, res.Ok)
	require.NotEmpty(t, res.Msg)

	password := "p"

	res, err = client.Auth(ctx, &pb.Credits{Ip: "192.168.1.1", Login: "l", Password: &password})
	require.NoError(t, err)
	require.Equal(t, false, res.Ok)

	res, err = client.Auth(ctx, &pb.Credits{Ip: "192.168.1.127", Login: "l", Password: &password})
	require.NoError(t, err)
	require.Equal(t, false, res.Ok)

	res, err = client.Auth(ctx, &pb.Credits{Ip: "192.168.1.128", Login: "l", Password: &password})
	require.NoError(t, err)
	require.Equal(t, true, res.Ok)

	for i := 0; i < ipLim+1; i++ {
		res, err = client.Auth(ctx, &pb.Credits{Ip: "192.168.1.192", Login: "l", Password: &password})
		require.NoError(t, err)
		require.Equal(t, true, res.Ok)
	}

	res, err = client.DeleteFromBlacklist(ctx, &pb.Subnet{Subnet: blackSubnet})
	require.NoError(t, err)
	require.Equal(t, true, res.Ok)

	res, err = client.DeleteFromWhitelist(ctx, &pb.Subnet{Subnet: whiteSubnet})
	require.NoError(t, err)
	require.Equal(t, true, res.Ok)

	for i := 0; i < loginLim; i++ {
		res, err := client.Auth(ctx, &pb.Credits{Ip: "192.168.1.192", Login: "login", Password: &password})
		require.NoError(t, err)
		require.Equal(t, true, res.Ok)
	}

	res, err = client.Auth(ctx, &pb.Credits{Ip: "192.168.1.192", Login: "login", Password: &password})
	require.NoError(t, err)
	require.Equal(t, false, res.Ok)

	password = "pp"
	for i := 0; i < passwordLim; i++ {
		res, err := client.Auth(ctx, &pb.Credits{Ip: "192.168.1.192", Login: fmt.Sprintf("%dl", i), Password: &password})
		require.NoError(t, err)
		require.Equal(t, true, res.Ok)
	}

	res, err = client.Auth(ctx, &pb.Credits{Ip: "192.168.1.192", Login: "l", Password: &password})
	require.NoError(t, err)
	require.Equal(t, false, res.Ok)

	for i := 0; i < ipLim; i++ {
		password := fmt.Sprintf("%dp", i)
		res, err := client.Auth(ctx, &pb.Credits{Ip: "192.168.1.193", Login: fmt.Sprintf("%dl", i), Password: &password})
		require.NoError(t, err)
		require.Equal(t, true, res.Ok)
	}

	password = "ppp"
	res, err = client.Auth(ctx, &pb.Credits{Ip: "192.168.1.193", Login: "l", Password: &password})
	require.NoError(t, err)
	require.Equal(t, false, res.Ok)

	time.Sleep(time.Duration(ttl) * time.Second)

	password = "pp"
	res, err = client.Auth(ctx, &pb.Credits{Ip: "192.168.1.192", Login: "login", Password: &password})
	require.NoError(t, err)
	require.Equal(t, true, res.Ok)

	for i := 0; i < loginLim; i++ {
		res, err := client.Auth(ctx, &pb.Credits{Ip: "192.168.1.192", Login: "l", Password: &password})
		require.NoError(t, err)
		require.Equal(t, true, res.Ok)
	}

	res, err = client.Auth(ctx, &pb.Credits{Ip: "192.168.1.193", Login: "l", Password: &password})
	require.NoError(t, err)
	require.Equal(t, false, res.Ok)

	res, err = client.Reset(ctx, &pb.Credits{Login: "l", Ip: "192.168.1.193"})
	require.NoError(t, err)
	require.Equal(t, true, res.Ok)

	res, err = client.Auth(ctx, &pb.Credits{Ip: "192.168.1.193", Login: "l", Password: &password})
	require.NoError(t, err)
	require.Equal(t, true, res.Ok)
}
