/*
Copyright © 2024 sfk <sfk@live.cn>

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
	"github.com/isfk/get-cdnjs/internal/qiniu"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete [目录名]",
	Short: "删除七牛云中的目录",
	Long: `删除七牛云中指定目录及其所有文件。

示例:
  ./get-cdnjs delete jquery        # 删除 jquery 目录
  ./get-cdnjs delete jquery@3.7.1  # 删除指定版本目录`,
	Run: func(cmd *cobra.Command, args []string) {
		targetPath := ""
		if len(args) > 0 {
			targetPath = args[0]
		}
		qiniu.RunDelete(targetPath)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
