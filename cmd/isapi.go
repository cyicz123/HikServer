/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// isapiCmd represents the isapi command
var isapiCmd = &cobra.Command{
	Use:   "isapi",
	Short: "ISAPI监听服务器",
	Long: `ISAPI监听服务器基于ISAPI协议接收设备传输的json/xml报文和二进制文件，并将其存放于设备IP同名文件夹下`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("isapi called")
	},
}

func init() {
	rootCmd.AddCommand(isapiCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// isapiCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// isapiCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
