/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	var (
		port uint16
		sdkCmd = &cobra.Command{
			Use:   "sdk",
			Short: "SDK监听服务",
			Long: `SDK监听服务，基于SDK协议接收设备透明传输的json/xml报文和二进制文件，并将其存放于设备IP同名文件夹下`,
			PreRun: func(cmd *cobra.Command, args []string) {
				if port != 0 {
					cfg.Sdk.Port = port
				}
			},
			Run: sdkEntry,
		}
	)

	rootCmd.AddCommand(sdkCmd)
	sdkCmd.Flags().Uint16VarP(&port, "port", "p", 0, "sdk port")
}

func sdkEntry(cmd *cobra.Command, args []string) {
	fmt.Printf("sdk port %d\n", cfg.Sdk.Port)
}