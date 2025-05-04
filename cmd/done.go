package cmd

import (
	"fmt"
	"gotodo/internal/renderer"
	"gotodo/internal/storage"
	"gotodo/pkg/logger"
	"os"

	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

// 将 delete.go 重命名为 done.go
var doneCmd = &cobra.Command{
	Use:     "done",
	Aliases: []string{"d"},  // 保留 d 作为快捷命令
	Short:   "标记任务为已完成",
	RunE: func(cmd *cobra.Command, args []string) error {
		notes, err := storage.ListNotes()
		if err != nil {
			return err
		}

		// 过滤出未完成的任务
		var unfinishedNotes []storage.Note
		for _, note := range notes {
			if note.Status != "done" {
				unfinishedNotes = append(unfinishedNotes, note)
			}
		}

		if len(unfinishedNotes) == 0 {
			color.Yellow("No pending tasks")
			return nil
		}

		// 创建选择提示
		// 更新模板设计
		templates := &promptui.SelectTemplates{
			Label:    "{{ . }}",
			Active:   "→ {{ .Title | cyan }} {{ .Content | white | bold }} {{ if .DueDate }}({{ .DueDate | magenta }}){{ end }} [{{ .Priority | red }}]",
			Inactive: "  {{ .Title | faint }} {{ .Content | faint }} {{ if .DueDate }}({{ .DueDate | faint }}){{ end }} [{{ .Priority | faint }}]",
			Selected: "✓ {{ .Title | green }} {{ .Content | green | bold }} {{ if .DueDate }}({{ .DueDate | green }}){{ end }} [{{ .Priority | green }}]",
		}

		prompt := promptui.Select{
			Label:     "选择要完成的任务",
			Items:     unfinishedNotes,
			Templates: templates,
			Size:      10,
		}

		i, _, err := prompt.Run()
		if err != nil {
			if err == promptui.ErrInterrupt {
				os.Exit(0)
			}
			return err
		}

		// 更新选中的任务状态
		selectedNote := unfinishedNotes[i]
		for i := range notes {
			if notes[i].ID == selectedNote.ID {
				notes[i].Status = "done"
				break
			}
		}

		// 保存更新后的任务列表
		if err := storage.SaveNotes(notes); err != nil {
			return err
		}

		logger.Success(fmt.Sprintf("Task completed: %s", selectedNote.Content))

		// 显示更新后的任务列表
		color.New(color.FgHiCyan).Println("\nCurrent tasks:")
		renderer.RenderNotes(notes, false, "")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(doneCmd)
}
