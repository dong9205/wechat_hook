package cmd

import (
	"github.com/dong9205/wechat_hook/pkg/serve"
	"github.com/spf13/cobra"
)

var serverCmd = cobra.Command{
	Use:   "serve",
	Short: "启动服务",
	Long:  `启动服务`,
	Run: func(cmd *cobra.Command, args []string) {
		serve.Serve()
	},
}

func init() {
	rootCmd.AddCommand(&serverCmd)
}
