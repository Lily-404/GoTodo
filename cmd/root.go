package cmd

import (
	"github.com/Lily-404/todo/internal/renderer"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var version = "v1.0.0"

var rootCmd = &cobra.Command{
	Use:   "todo",
	Short: color.HiCyanString("Todo - A Minimalist Terminal Task Manager"),
	Long: color.HiWhiteString(`Todo is a minimalist yet powerful terminal task manager,
focused on helping you efficiently manage your todos.

Command Aliases:
  add (a)  - Add task
  list (l) - List tasks
  done (d) - Complete task
  clean (c) - Clean tasks

Features:`) + color.HiCyanString(`
  - Clean command-line interface
  - Efficient task management
  - Hacker-style terminal experience`),
	Version: version,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		renderer.ShowBanner()
	},
}

func Execute() error {
	return rootCmd.Execute()
}
