package cmd

import (
	"github.com/Lily-404/todo/internal/config"
	"github.com/Lily-404/todo/internal/i18n"
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
	Short:   i18n.GetMessage(config.GetConfig().Language, "cmd_list_short"),
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

		renderer.RenderNotes(notes, true, filterPriority) // 将 showAll 改为 true
		return nil
	},
}

func init() {
	help := color.New(color.FgHiBlue).SprintFunc()
	listCmd.Flags().BoolVarP(&showAll, "pending", "p", false, help(i18n.GetMessage(config.GetConfig().Language, "show_pending_only")))
	listCmd.Flags().StringVarP(&filterPriority, "priority", "r", "", help(i18n.GetMessage(config.GetConfig().Language, "filter_by_priority")))
	rootCmd.AddCommand(listCmd)
}
