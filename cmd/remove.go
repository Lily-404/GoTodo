package cmd

import (
	"fmt"
	"github.com/Lily-404/todo/internal/renderer"
	"github.com/Lily-404/todo/internal/storage"
	"github.com/Lily-404/todo/pkg/logger"

	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:     "remove",
	Aliases: []string{"rm"},
	Short:   "删除指定的任务",
	RunE: func(cmd *cobra.Command, args []string) error {
		notes, err := storage.ListNotes()
		if err != nil {
			return err
		}

		// 过滤未完成的任务
		var unfinishedNotes []storage.Note
		for _, note := range notes {
			if note.Status != "done" {
				unfinishedNotes = append(unfinishedNotes, note)
			}
		}

		if len(unfinishedNotes) == 0 {
			color.Yellow("没有未完成的任务可供删除")
			return nil
		}

		// 创建任务选择提示
		templates := &promptui.SelectTemplates{
			Label:    "{{ . }}",
			Active:   "→ {{ .Title | cyan }} {{ .Content | white | bold }} {{ if .DueDate }}({{ .DueDate | magenta }}){{ end }} [{{ .Priority | red }}]",
			Inactive: "  {{ .Title | faint }} {{ .Content | faint }} {{ if .DueDate }}({{ .DueDate | faint }}){{ end }} [{{ .Priority | faint }}]",
			Selected: "✓ {{ .Title | green }} {{ .Content | green | bold }} {{ if .DueDate }}({{ .DueDate | green }}){{ end }} [{{ .Priority | green }}]",
		}

		prompt := promptui.Select{
			Label:     "选择要删除的任务",
			Items:     unfinishedNotes,
			Templates: templates,
			Size:      10,
		}

		idx, _, err := prompt.Run()
		if err != nil {
			return fmt.Errorf("选择任务失败: %v", err)
		}

		selectedNote := unfinishedNotes[idx]

		// 删除选中的任务
		if err := storage.DeleteNote(selectedNote.ID); err != nil {
			return err
		}

		logger.Success(fmt.Sprintf("已删除任务: %s", selectedNote.Content))

		// 显示更新后的任务列表
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
	rootCmd.AddCommand(removeCmd)
}