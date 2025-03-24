package cmd

import (
	"fmt"

	"github.com/slem7451/anti_bruteforce/internal/server/grpc/pb"
	"github.com/spf13/cobra"
)

func addCmd() *cobra.Command {
	var subnet string
	var bType bool
	var wType bool

	printFunc := func (res *pb.Response, err error) {
		if err != nil {
			fmt.Println(err)
			return
		}

		if res.Ok {
			fmt.Println("Подсеть добавлена в список")
		} else {
			fmt.Println("Не удалось добавить подсеть в список")
			if res.GetMsg() != "" {
				fmt.Println(res.GetMsg())
			}
		}
	}

	var addCmd = &cobra.Command{
		Use:   "add-to-list",
		Short: "Добавить подсеть в black/white лист",
		Long: `Подсеть вида "IP + маска" добавляется в заданный список`,
		Run: func(cmd *cobra.Command, args []string) {
			var res *pb.Response
			var err error

			if bType && wType {
				fmt.Println("Подсеть добавится в оба списка")
			}

			if bType {
				res, err = client.AddToBlacklist(ctx, &pb.Subnet{Subnet: subnet})
				printFunc(res, err)
			}

			if wType {
				res, err = client.AddToWhitelist(ctx, &pb.Subnet{Subnet: subnet})
				printFunc(res, err)
			}
		},
	}

	addCmd.Flags().StringVarP(&subnet, "subnet", "s", "", "Подсеть (IP + маска)")
	addCmd.Flags().BoolVar(&bType, "b", false, "Blacklist")
	addCmd.Flags().BoolVar(&wType, "w", false, "Whitelist")

	addCmd.MarkFlagRequired("subnet")
	addCmd.MarkFlagsOneRequired("b", "w")

	return addCmd
}