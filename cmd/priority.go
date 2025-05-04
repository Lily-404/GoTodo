package cmd

import (
	"fmt"
	"github.com/Lily-404/todo/internal/renderer"
	"github.com/Lily-404/todo/internal/storage"
	"github.com/Lily-404/todo/pkg/logger"
	"os"

	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(priorityCmd)
}

var priorityCmd = &cobra.Command{
	Use:     "priority",
	Aliases: []string{"p"},
	Short:   "修改任务优先级",
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
			color.Yellow("没有待处理的任务")
			return nil
		}

		// 选择任务
		templates := &promptui.SelectTemplates{
			Label:    "{{ . }}",
			Active:   "→ {{ .Title | cyan }} {{ .Content | white | bold }} [{{ .Priority | red }}]",
			Inactive: "  {{ .Title | faint }} {{ .Content | faint }} [{{ .Priority | faint }}]",
			Selected: "✓ {{ .Title | green }} {{ .Content | green | bold }} [{{ .Priority | green }}]",
		}

		prompt := promptui.Select{
			Label:     "选择要修改优先级的任务",
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

		// 选择新的优先级（从低到高排序）
		priorities := []string{"low", "normal", "high"}
		priorityPrompt := promptui.Select{
			Label: "选择新的优先级",
			Items: priorities,
		}

		p, _, err := priorityPrompt.Run()
		if err != nil {
			return err
		}

		// 更新选中任务的优先级
		selectedNote := unfinishedNotes[i]
		for i := range notes {
			if notes[i].ID == selectedNote.ID {
				notes[i].Priority = priorities[p]
				break
			}
		}

		// 保存更新后的任务列表
		if err := storage.SaveNotes(notes); err != nil {
			return err
		}

		logger.Success(fmt.Sprintf("任务优先级已更新: %s -> %s", selectedNote.Content, priorities[p]))

		// 显示更新后的任务列表
		color.New(color.FgHiCyan).Println("\n当前任务列表:")
		renderer.RenderNotes(notes, false, "")
		return nil
	},
}
