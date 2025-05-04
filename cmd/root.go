package cmd

import (
	"hacknote/internal/renderer"
	"hacknote/pkg/logger"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var verbose bool

var rootCmd = &cobra.Command{
	Use:   "gotodo",
	Short: color.HiCyanString("GoTodo - 简约的终端任务管理工具"),
	Long: color.HiWhiteString(`GoTodo 是一个简约而强大的终端任务管理工具，
专注于帮助你高效管理待办事项。

特性:`) + color.HiCyanString(`
  - 简洁的命令行界面
  - 高效的任务管理
  - 黑客风格的终端体验`),
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		logger.SetVerbose(verbose)
		renderer.ShowBanner()
	},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "启用详细日志输出")
}
