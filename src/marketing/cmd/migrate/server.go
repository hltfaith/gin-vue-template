package migrate

import (
	"fmt"
	"github.com/spf13/cobra"
	_ "github.com/go-sql-driver/mysql"

	"marketing/app/admin/models"
	"marketing/common/database"
)

var StartCmd = &cobra.Command{
	Use:			"migrate",
	Short:			"初始化数据库表结构",
	Example:		"server migrate",
	Run: func(cmd *cobra.Command, args []string) {
		run()
	},
}

func run() {
	// 开始执行相关流程
	fmt.Println(`start migrate`)
	initDB()
	fmt.Println(`migrate success.`)
}

func initDB() {
	// 初始化数据库
	db, err := database.Connect()
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	db.AutoMigrate(&models.Permission{})
	db.AutoMigrate(&models.RoleBePermission{})
	db.AutoMigrate(&models.Role{})
	db.AutoMigrate(&models.UserProfile{})
	db.AutoMigrate(&models.User{})
}
