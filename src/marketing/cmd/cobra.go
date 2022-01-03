package cmd

import (
	"os"
	"fmt"
	"github.com/spf13/cobra"
	
	"marketing/cmd/migrate"
	"marketing/cmd/api"
	"marketing/cmd/test"
)

var rootCmd = &cobra.Command{
	Use:			"server",
	Short:			"server",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("请使用 -h 查看命令")
	},
}

func init() {
	// 新增子命令
	rootCmd.AddCommand(migrate.StartCmd)
	rootCmd.AddCommand(api.StartCmd)
	rootCmd.AddCommand(test.StartCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}