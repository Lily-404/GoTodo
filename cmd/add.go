package cmd

import (
	"hacknote/internal/renderer"
	"hacknote/internal/storage"
	"hacknote/pkg/logger"
	"time"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var (
	title    string
	priority string
	dueDate  string
)

var addCmd = &cobra.Command{
	Use:   "add [内容]",
	Short: "添加一条新的笔记",
	Example: `  gotodo add "完成项目文档" -t "文档" -p high
  gotodo add "回复邮件" -t "工作" -d "2024-01-20"`,
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

		if err := storage.AddNote(note); err != nil {
			return err
		}

		// 添加成功后显示提示
		logger.Success("任务添加成功")

		// 获取并显示所有未完成的任务
		notes, err := storage.ListNotes()
		if err != nil {
			return err
		}

		color.New(color.FgHiCyan).Println("\n当前未完成的任务：")
		renderer.RenderNotes(notes, false, "")
		return nil
	},
}

func init() {
	addCmd.Flags().StringVarP(&title, "title", "t", "", "笔记标题")
	addCmd.Flags().StringVarP(&priority, "priority", "p", "normal", "优先级 (high/normal/low)")
	addCmd.Flags().StringVarP(&dueDate, "due", "d", "", "截止日期 (YYYY-MM-DD)")
	rootCmd.AddCommand(addCmd)
}
