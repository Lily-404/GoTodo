package cmd

import (
	"fmt"
	"time"

	"github.com/Lily-404/todo/internal/config"
	"github.com/Lily-404/todo/internal/i18n"
	"github.com/Lily-404/todo/internal/renderer"
	"github.com/Lily-404/todo/internal/storage"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var (
	title    string
	priority string
	dueDate  string
)

var addCmd = &cobra.Command{
	Use:     "add [content]",
	Aliases: []string{"a", "+"},
	Short:   i18n.GetMessage(config.GetConfig().Language, "cmd_add_short"),
	Example: `  todo add "Complete project documentation" -t "docs" -p high
  todo a "Reply to email" -t "work" -d "2024-01-20"`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		content := args[0]
		note := storage.Note{
			Title:     title,
			Content:   content,
			Priority:  priority,
			DueDate:   dueDate,
			Status:    "pending",
			CreatedAt: time.Now(),
		}

		// use i18n
		priorityPrompt := promptui.Select{
			Label: i18n.GetMessage(config.GetConfig().Language, "select_priority"),
			Items: []string{"low", "normal", "high"},
			Templates: &promptui.SelectTemplates{
				Label:    "{{ . }}",
				Active:   "➤ {{ . | cyan }}",
				Inactive: "  {{ . }}",
				Selected: "✓ {{ . | green }}",
			},
		}

		priorityIdx, _, err := priorityPrompt.Run()
		if err != nil {
			return fmt.Errorf(i18n.GetMessage(config.GetConfig().Language, "priority_select_failed"), err)
		}

		priorities := []string{"low", "normal", "high"} // 这里也需要保持相同的顺序
		note.Priority = priorities[priorityIdx]

		if addErr := storage.AddNote(note); addErr != nil {
			return err
		}

		// 获取并显示所有未完成的任务
		notes, err := storage.ListNotes()
		if err != nil {
			return err
		}

		renderer.RenderNotes(notes, false, "")
		return nil
	},
}

func init() {
	addCmd.Flags().StringVarP(&title, "title", "t", "", i18n.GetMessage(config.GetConfig().Language, "flag_title"))
	addCmd.Flags().StringVarP(&priority, "priority", "p", "low", i18n.GetMessage(config.GetConfig().Language, "flag_priority"))
	addCmd.Flags().StringVarP(&dueDate, "due", "d", "", i18n.GetMessage(config.GetConfig().Language, "flag_due_date"))
	rootCmd.AddCommand(addCmd)
}
