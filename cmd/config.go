package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var configCmd = cobra.Command{
	Use:   "config [option]",
	Short: "配置文件相关",
	Long:  `配置文件相关`,
}

var configPrintCmd = cobra.Command{
	Use:   "print",
	Short: "打印配置文件",
	Long:  `打印配置文件`,
	Run: func(cmd *cobra.Command, args []string) {
		for key, value := range viper.AllSettings() {
			fmt.Println(key, value)
		}
	},
}

func init() {
	rootCmd.AddCommand(&configCmd)
	configCmd.AddCommand(&configPrintCmd)
}
