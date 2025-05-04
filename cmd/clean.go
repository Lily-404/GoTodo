package cmd

import (
	"fmt"
	"github.com/Lily-404/todo/internal/renderer"
	"github.com/Lily-404/todo/internal/storage"
	"github.com/Lily-404/todo/pkg/logger"

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

		// 过滤出已完成和未完成的任务
		var unfinishedNotes []storage.Note
		var finishedNotes []storage.Note
		for _, note := range notes {
			if note.Status == "done" {
				finishedNotes = append(finishedNotes, note)
			} else {
				unfinishedNotes = append(unfinishedNotes, note)
			}
		}

		if len(finishedNotes) == 0 {
			color.Yellow("  没有已完成的任务需要清理")
			return nil
		}

		// 显示已完成的任务列表
		color.HiCyan("\n将要清理的已完成任务：")
		for i, note := range finishedNotes {
			color.HiBlack(fmt.Sprintf("  %d. %s", i+1, note.Content))
		}
		fmt.Println()

		// 保存未完成的任务
		if saveErr := storage.SaveNotes(unfinishedNotes); saveErr != nil {
			return err
		}

		logger.Success(fmt.Sprintf("已清理 %d 个已完成的任务", len(finishedNotes)))
		
		// 显示当前所有任务
		fmt.Println("\n当前任务列表：")
		updatedNotes, err := storage.ListNotes()
		if err != nil {
			return err
		}
		renderer.RenderNotes(updatedNotes, false, "")
		
		return nil
	},
}

func init() {
	rootCmd.AddCommand(cleanCmd)
}
