package cmd

import (
	"fmt"

	"github.com/Lily-404/todo/internal/config"
	"github.com/Lily-404/todo/internal/i18n"
	"github.com/Lily-404/todo/internal/renderer"
	"github.com/Lily-404/todo/internal/storage"
	"github.com/Lily-404/todo/pkg/logger"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var cleanCmd = &cobra.Command{
	Use:     "clean",
	Aliases: []string{"c"},
	Short:   i18n.GetMessage(config.GetConfig().Language, "cmd_clean_short"),
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
			color.Yellow(i18n.GetMessage(config.GetConfig().Language, "no_completed_tasks"))
			return nil
		}

		// 显示已完成的任务列表
		color.HiCyan("\n" + i18n.GetMessage(config.GetConfig().Language, "tasks_to_clean"))
		for i, note := range finishedNotes {
			color.HiBlack(fmt.Sprintf("  %d. %s", i+1, note.Content))
		}
		fmt.Println()

		// 保存未完成的任务
		if saveErr := storage.SaveNotes(unfinishedNotes); saveErr != nil {
			return err
		}

		logger.Success(i18n.GetMessage(config.GetConfig().Language, "cleaned_tasks", len(finishedNotes)))

		fmt.Println("\n" + i18n.GetMessage(config.GetConfig().Language, "current_tasks"))
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
