package cmd //nolint:dupl

import (
	"fmt"

	"github.com/slem7451/anti_bruteforce/internal/server/grpc/pb"
	"github.com/spf13/cobra"
)

func deleteCmd() *cobra.Command {
	var subnet string
	var bType bool
	var wType bool

	printFunc := func(res *pb.Response, err error) {
		if err != nil {
			fmt.Println(err)
			return
		}

		if res.Ok {
			fmt.Println("Подсеть удалена из списка")
		} else {
			fmt.Println("Не удалось удалить подсеть из списка")
			if res.GetMsg() != "" {
				fmt.Println(res.GetMsg())
			}
		}
	}

	deleteCmd := &cobra.Command{
		Use:   "delete-from-list",
		Short: "Удалить подсеть из black/white листа",
		Long:  `Подсеть вида "IP + маска" удаляется из заданного списка`,
		Run: func(_ *cobra.Command, _ []string) {
			var res *pb.Response
			var err error

			if bType && wType {
				fmt.Println("Подсеть удалится из обеих списков")
			}

			if bType {
				res, err = client.DeleteFromBlacklist(ctx, &pb.Subnet{Subnet: subnet})
				printFunc(res, err)
			}

			if wType {
				res, err = client.DeleteFromWhitelist(ctx, &pb.Subnet{Subnet: subnet})
				printFunc(res, err)
			}
		},
	}

	deleteCmd.Flags().StringVarP(&subnet, "subnet", "s", "", "Подсеть (IP + маска)")
	deleteCmd.Flags().BoolVar(&bType, "b", false, "Blacklist")
	deleteCmd.Flags().BoolVar(&wType, "w", false, "Whitelist")

	deleteCmd.MarkFlagRequired("subnet")
	deleteCmd.MarkFlagsOneRequired("b", "w")

	return deleteCmd
}
