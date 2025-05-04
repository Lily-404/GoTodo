package cmd

import (
	"github.com/Lily-404/todo/internal/renderer"
	"github.com/Lily-404/todo/internal/storage"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var (
	showAll        bool
	filterPriority string
)

var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"l"},
	Short:   "List all tasks",
	Example: `  todo list
  todo list --all
  todo list -p high`,
	RunE: func(cmd *cobra.Command, args []string) error {
		info := color.New(color.FgHiCyan, color.Bold)
		notes, err := storage.ListNotes()
		if err != nil {
			return err
		}
		if len(notes) == 0 {
			info.Println("  没有任务...")
			return nil
		}

		renderer.RenderNotes(notes, true, filterPriority)  // 将 showAll 改为 true
		return nil
	},
}

func init() {
	help := color.New(color.FgHiBlue).SprintFunc()
	listCmd.Flags().BoolVarP(&showAll, "pending", "p", false, help("只显示未完成的任务"))  // 修改参数含义
	listCmd.Flags().StringVarP(&filterPriority, "priority", "r", "", help("按优先级筛选 (high/normal/low)"))
	rootCmd.AddCommand(listCmd)
}
