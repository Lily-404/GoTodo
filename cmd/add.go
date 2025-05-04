package cmd

import (
	"fmt"
	"github.com/Lily-404/todo/internal/renderer"
	"github.com/Lily-404/todo/internal/storage"
	"time"

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
	Aliases: []string{"a","+"},
	Short:   "Add a new note",
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

		// 创建优先级选择
		priorityPrompt := promptui.Select{
			Label: "选择任务优先级",
			Items: []string{"low", "normal", "high"},  // 改为从低到高的顺序
			Templates: &promptui.SelectTemplates{
				Label:    "{{ . }}",
				Active:   "➤ {{ . | cyan }}",
				Inactive: "  {{ . }}",
				Selected: "✓ {{ . | green }}",
			},
		}

		priorityIdx, _, err := priorityPrompt.Run()
		if err != nil {
			return fmt.Errorf("选择优先级失败: %v", err)
		}

		priorities := []string{"low", "normal", "high"}  // 这里也需要保持相同的顺序
		note.Priority = priorities[priorityIdx]

		if err := storage.AddNote(note); err != nil {
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
	addCmd.Flags().StringVarP(&title, "title", "t", "", "Note title")
	addCmd.Flags().StringVarP(&priority, "priority", "p", "low", "Priority (high/normal/low)")
	addCmd.Flags().StringVarP(&dueDate, "due", "d", "", "Due date (YYYY-MM-DD)")
	rootCmd.AddCommand(addCmd)
}
