package cmd

import (
	"fmt"
	"gotodo/internal/renderer"
	"gotodo/internal/storage"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:     "remove",
	Aliases: []string{"r", "-"},
	Short:   "删除一个未完成的任务",
	RunE: func(cmd *cobra.Command, args []string) error {
		// 获取所有任务
		notes, err := storage.ListNotes()
		if err != nil {
			return err
		}

		// 获取未完成的任务
		var unfinishedNotes []storage.Note
		for _, note := range notes {
			if note.Status != "done" {
				unfinishedNotes = append(unfinishedNotes, note)
			}
		}

		if len(unfinishedNotes) == 0 {
			fmt.Println("没有未完成的任务可供删除")
			return nil
		}

		// 创建选择提示
		templates := &promptui.SelectTemplates{
			Label:    "{{ . }}",
			Active:   "➤ {{ .Content | cyan }}",
			Inactive: "  {{ .Content }}",
			Selected: "✓ {{ .Content | green }}",
		}

		prompt := promptui.Select{
			Label:     "选择要删除的任务",
			Items:     unfinishedNotes,
			Templates: templates,
		}

		idx, _, err := prompt.Run()
		if err != nil {
			return fmt.Errorf("选择任务失败: %v", err)
		}

		// 删除选中的任务
		if err := storage.DeleteNote(unfinishedNotes[idx].ID); err != nil {
			return err
		}

		// 重新获取并显示剩余的任务
		updatedNotes, err := storage.ListNotes()
		if err != nil {
			return err
		}

		renderer.RenderNotes(updatedNotes, false, "")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
