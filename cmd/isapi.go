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
		url string
		isapiCmd = &cobra.Command{
			Use:   "isapi",
			Short: "ISAPI监听服务器",
			Long: `ISAPI监听服务器基于ISAPI协议接收设备传输的json/xml报文和二进制文件，并将其存放于设备IP同名文件夹下`,
			PreRun: func(cmd *cobra.Command, args []string) {
				if port != 0 {
					cfg.Isapi.Port = port
				}
				if url != "" {
					cfg.Isapi.Url = url
				}
			},
			Run: isapiEntry,
		}
		
	)

	rootCmd.AddCommand(isapiCmd)
	isapiCmd.Flags().Uint16VarP(&port, "port", "p", 0, "isapi port")
	isapiCmd.Flags().StringVarP(&url, "url", "u", "", "isapi server url")
}

func isapiEntry(cmd *cobra.Command, args []string) {
	fmt.Printf("isapi port %d\n", cfg.Isapi.Port)
	fmt.Printf("host url %s\n", cfg.Isapi.Url)
}