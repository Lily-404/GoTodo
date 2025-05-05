package cmd

import (
	"fmt"
	"os"

	"github.com/Lily-404/todo/internal/config"
	"github.com/Lily-404/todo/internal/i18n"
	"github.com/Lily-404/todo/internal/renderer"
	"github.com/Lily-404/todo/internal/storage"
	"github.com/Lily-404/todo/pkg/logger"

	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var doneCmd = &cobra.Command{
	Use:     "done",
	Aliases: []string{"d"},
	Short:   i18n.GetMessage(config.GetConfig().Language, "cmd_done_short"),
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
			color.Yellow(i18n.GetMessage(config.GetConfig().Language, "no_pending_tasks"))
			return nil
		}

		templates := &promptui.SelectTemplates{
			Label:    "{{ . }}",
			Active:   "→ {{ .Title | cyan }} {{ .Content | white | bold }} {{ if .DueDate }}({{ .DueDate | magenta }}){{ end }} [{{ .Priority | red }}]",
			Inactive: "  {{ .Title | faint }} {{ .Content | faint }} {{ if .DueDate }}({{ .DueDate | faint }}){{ end }} [{{ .Priority | faint }}]",
			Selected: "✓ {{ .Title | green }} {{ .Content | green | bold }} {{ if .DueDate }}({{ .DueDate | green }}){{ end }} [{{ .Priority | green }}]",
		}

		prompt := promptui.Select{
			Label:     i18n.GetMessage(config.GetConfig().Language, "select_task_to_complete"),
			Items:     unfinishedNotes,
			Templates: templates,
			Size:      10,
		}

		i, _, err := prompt.Run()
		if err != nil {
			if err == promptui.ErrInterrupt {
				os.Exit(0)
			}
			return fmt.Errorf(i18n.GetMessage(config.GetConfig().Language, "task_select_failed"), err)
		}

		selectedNote := unfinishedNotes[i]
		for i := range notes {
			if notes[i].ID == selectedNote.ID {
				notes[i].Status = "done"
				break
			}
		}

		if err := storage.SaveNotes(notes); err != nil {
			return err
		}

		logger.Success(i18n.GetMessage(config.GetConfig().Language, "task_completed", selectedNote.Content))

		color.New(color.FgHiCyan).Println("\n" + i18n.GetMessage(config.GetConfig().Language, "current_tasks"))
		renderer.RenderNotes(notes, false, "")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(doneCmd)
}
