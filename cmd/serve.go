package cmd

import (
	"github.com/dong9205/wechat_hook/pkg/config"
	"github.com/dong9205/wechat_hook/pkg/logger"
	"github.com/dong9205/wechat_hook/pkg/serve"
	"github.com/dong9205/wechat_hook/pkg/wx"
	"github.com/spf13/cobra"
)

var serverCmd = cobra.Command{
	Use:   "serve",
	Short: "启动服务",
	Long:  `启动服务`,
	Run: func(cmd *cobra.Command, args []string) {
		// config
		config.ConfigInit(rootCmd.Flags().String("config", "./config.yaml", "配置文件路径"))
		// 日志
		logger.InitLogger()
		// wechat bot
		_ = wx.GetBotMgr()
		serve.Serve()
	},
}

func init() {
	rootCmd.AddCommand(&serverCmd)
}
