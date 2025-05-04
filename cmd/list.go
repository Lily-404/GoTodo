package cmd

import (
	"gotodo/internal/renderer"
	"gotodo/internal/storage"

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
	Example: `  gotodo list
  gotodo list --all
  gotodo list -p high`,
	RunE: func(cmd *cobra.Command, args []string) error {
		info := color.New(color.FgHiCyan, color.Bold)
		notes, err := storage.ListNotes()
		if err != nil {
			return err
		}
		if len(notes) == 0 {
			info.Println("  No tasks yet...")
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
	help := color.New(color.FgHiBlue).SprintFunc()
	listCmd.Flags().BoolVarP(&showAll, "all", "a", false, help("Show all tasks (including completed)"))
	listCmd.Flags().StringVarP(&filterPriority, "priority", "p", "", help("Filter by priority (high/normal/low)"))
	rootCmd.AddCommand(listCmd)
}
