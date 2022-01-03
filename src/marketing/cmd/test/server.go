package test

import (
	// "fmt"
	"github.com/spf13/cobra"

	"marketing/common/models"
)

var StartCmd = &cobra.Command{
	Use:			"test",
	Short:			"仅测试使用",
	Example:		"server test",
	Run: func(cmd *cobra.Command, args []string) {
		run()
	},
}

func run() {
	var num uint = 1
	models.GetRoleIDFromRId(&num)
	// fmt.Println(res)
}
