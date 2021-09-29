package migrate

import (
	"github.com/RealLiuSha/echo-admin/lib"
	"github.com/RealLiuSha/echo-admin/models"
	"github.com/spf13/cobra"
)

var configFile string

func init() {
	//PersistentFlags 全局性flag 可用于它所分配的命令以及该命令下的每个命令
	pf := StartCmd.PersistentFlags()
	pf.StringVarP(&configFile, "config", "c",
		"config/config.yaml", "this parameter is used to start the service application")

	//Required flagsRequired flags 必选flag
	cobra.MarkFlagRequired(pf, "config")
}

var StartCmd = &cobra.Command{
	Use:          "migrate",
	Short:        "Migrate database",
	Example:      "{execfile} migrate -c config/config.yaml",
	SilenceUsage: true,
	// *Run 函数按以下顺序执行：
	// * PersistentPreRun()
	// * PreRun()
	// * Run()
	// * PostRun()
	// * PersistentPostRun()
	// 所有函数获取相同的参数，命令名称后的参数。
	// // PersistentPreRun：此命令的子项将继承并执行。
	PreRun: func(cmd *cobra.Command, args []string) {
		lib.SetConfigPath(configFile)
	},
	Run: func(cmd *cobra.Command, args []string) {
		config := lib.NewConfig()
		logger := lib.NewLogger(config)
		db := lib.NewDatabase(config, logger)

		if err := db.ORM.AutoMigrate(
			&models.User{},
			&models.UserRole{},
			&models.Role{},
			&models.RoleMenu{},
			&models.Menu{},
			&models.MenuAction{},
			&models.MenuActionResource{},
		); err != nil {
			logger.Zap.Fatalf("Error to migrate database: %v", err)
		}
	},
}
