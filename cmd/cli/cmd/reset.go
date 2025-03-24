package cmd

import (
	"fmt"

	"github.com/slem7451/anti_bruteforce/internal/server/grpc/pb"
	"github.com/spf13/cobra"
)

func resetCmd() *cobra.Command {
	var login string
	var ip string

	var resetCmd = &cobra.Command{
		Use:   "reset-auth",
		Short: "Cбросить попытки авторизации по логину и IP",
		Run: func(cmd *cobra.Command, args []string) {
			res, err := client.Reset(ctx, &pb.Credits{Login: login, Ip: ip})
			if err != nil {
				fmt.Println(err)
				return
			}

			if res.Ok {
				fmt.Println("Сброс выполнен")
			} else {
				fmt.Println("Не удалось сбросить")
				if res.GetMsg() != "" {
					fmt.Println(res.GetMsg())
				}
			}
		},
	}

	resetCmd.Flags().StringVar(&ip, "ip", "", "IP пользвателя")
	resetCmd.Flags().StringVarP(&login, "login", "l", "", "Логин пользователя")

	resetCmd.MarkFlagRequired("ip")
	resetCmd.MarkFlagRequired("login")

	return resetCmd
}
