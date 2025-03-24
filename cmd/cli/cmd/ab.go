package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/slem7451/anti_bruteforce/internal/server/grpc/pb"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var client pb.AuthClient
var ctx context.Context
var address string
var conn *grpc.ClientConn

func Execute() {
	var rootCmd = &cobra.Command{
		Use:   "ab",
		Args:  cobra.MinimumNArgs(1),
		Short: "CLI к анти-брутфорсу",
		Long:  `Позволяет управлять black/white листами и сбросить попытки авторизации по логину и IP`,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			var err error

			conn, err = grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			client = pb.NewAuthClient(conn)
    		ctx = context.Background()
		},
		PersistentPostRun: func(cmd *cobra.Command, args []string) {
			if conn != nil {
				conn.Close()
			}
		},
	}

	rootCmd.PersistentFlags().StringVarP(&address, "address", "a", "", "Адрес к GRPC-серверу")
	rootCmd.MarkPersistentFlagRequired("address")

	rootCmd.AddCommand(resetCmd(), addCmd(), deleteCmd())

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
