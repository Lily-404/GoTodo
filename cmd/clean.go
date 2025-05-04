package cmd

import (
	"fmt"
	"gotodo/internal/storage"
	"gotodo/pkg/logger"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var cleanCmd = &cobra.Command{
	Use:     "clean",
	Aliases: []string{"c"},
	Short:   "Clean completed tasks",
	RunE: func(cmd *cobra.Command, args []string) error {
		notes, err := storage.ListNotes()
		if err != nil {
			return err
		}

		// 过滤出未完成的任务
		var unfinishedNotes []storage.Note
		var finishedCount int

		for _, note := range notes {
			if note.Status != "done" {
				unfinishedNotes = append(unfinishedNotes, note)
			} else {
				finishedCount++
			}
		}

		if finishedCount == 0 {
			color.Yellow("  No completed tasks to clean")
			return nil
		}

		// 保存未完成的任务
		if err := storage.SaveNotes(unfinishedNotes); err != nil {
			return err
		}

		logger.Success(fmt.Sprintf("Cleaned %d completed tasks", finishedCount))
		return nil
	},
}

func init() {
	rootCmd.AddCommand(cleanCmd)
}
