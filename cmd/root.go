/*
Copyright © 2022 chengyue5 chengyue5@hikvision.com.cn

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"os"
	"github.com/cyicz123/HikServer/config"
	"github.com/spf13/cobra"
)



// rootCmd represents the base command when called without any subcommands
var (
	rootCmd = &cobra.Command{
		Use:   "hikServer",
		Short: "测试设备ISAPI、SDK监听功能的工具",
		Long: `hikServer能够建立一个ISAPI或者SDK监听服务器
并且将设备上传的报文和图片放在设备ip同名的文件夹下`,
		// Uncomment the following line if your bare application
		// has an action associated with it:
		// Run: func(cmd *cobra.Command, args []string) { },
	}
	cfg *config.Config
	cfgErr error
)
// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cfg, cfgErr = config.NewConfig()
	if cfgErr != cfgErr {
		println(cfgErr)
		os.Exit(1)
	}
	cfgErr = rootCmd.Execute()
	if cfgErr != nil {
		println(cfgErr)
		os.Exit(1)
	}
}



