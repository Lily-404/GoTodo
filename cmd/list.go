package cmd

import (
	"hacknote/internal/renderer"
	"hacknote/internal/storage"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var (
	showAll        bool
	filterPriority string
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "列出所有任务",
	Example: `  gotodo list
  gotodo list --all
  gotodo list -p high`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// 使用不同颜色显示命令执行信息
		info := color.New(color.FgHiCyan, color.Bold)
		notes, err := storage.ListNotes()
		if err != nil {
			return err
		}
		if len(notes) == 0 {
			info.Println("  暂无任务...")
		}

		notes, err = storage.ListNotes()
		if err != nil {
			return err
		}

		renderer.RenderNotes(notes, showAll, filterPriority)
		return nil
	},
}

func init() {
	// 使用彩色提示来显示帮助信息
	help := color.New(color.FgHiBlue).SprintFunc()
	listCmd.Flags().BoolVarP(&showAll, "all", "a", false, help("显示所有任务（包括已完成）"))
	listCmd.Flags().StringVarP(&filterPriority, "priority", "p", "", help("按优先级筛选 (high/normal/low)"))
	rootCmd.AddCommand(listCmd)
}
