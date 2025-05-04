package cmd

import (
	"fmt"
	"hacknote/internal/storage"
	"hacknote/pkg/logger"
	"os"

	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var doneCmd = &cobra.Command{
	Use:   "done",
	Short: "交互式选择并完成任务",
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
			color.Yellow("没有待完成的任务")
			return nil
		}

		// 创建选择提示
		templates := &promptui.SelectTemplates{
			Label:    "{{ . }}",
			Active:   "→ {{ .Content | magenta }}", // 使用洋红色突出当前选中项
			Inactive: "  {{ .Content | faint }}",   // 使用暗淡效果显示未选中项
			Selected: "✓ {{ .Content | yellow }}",  // 使用黄色显示确认选择
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

		logger.Success(fmt.Sprintf("任务已完成: %s", selectedNote.Content))
		return nil
	},
}

func init() {
	rootCmd.AddCommand(doneCmd)
}
