package cmd

import (
	"fmt"
	"hacknote/internal/storage"
	"hacknote/pkg/logger"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "清理已完成的任务",
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
			color.Yellow("  没有已完成的任务需要清理")
			return nil
		}

		// 保存未完成的任务
		if err := storage.SaveNotes(unfinishedNotes); err != nil {
			return err
		}

		logger.Success(fmt.Sprintf("已清理 %d 个已完成的任务", finishedCount))
		return nil
	},
}

func init() {
	rootCmd.AddCommand(cleanCmd)
}
