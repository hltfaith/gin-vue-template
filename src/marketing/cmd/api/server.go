package api

import (
	"github.com/spf13/cobra"
	"github.com/gin-gonic/gin"

	"marketing/app/admin/router"
	"marketing/common/middleware"
)

var StartCmd = &cobra.Command{
	Use:			"run",
	Short:			"运行API服务",
	Example:		"server run",
	Run: func(cmd *cobra.Command, args []string) {
		run()
	},
}

func run() {
	r := gin.Default()
	r.Use(middleware.Cors())
	router.InitRouter(r)
	r.Run(":8002")
}
