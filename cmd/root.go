package cmd

import (
	"github.com/dong9205/wechat_hook/pkg/config"
	"github.com/dong9205/wechat_hook/pkg/logger"
	"github.com/dong9205/wechat_hook/pkg/wx"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "wechat-webook",
	Short: "调用微信接口",
	Long:  `通过微信的接口发送消息，公众号推送消息`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	// config
	config.ConfigInit(rootCmd.Flags().String("config", "./config.yaml", "配置文件路径"))
	// 日志
	logger.InitLogger()
	// wechat bot
	_ = wx.GetBotMgr()
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
