package cmd

import (
	"gotodo/internal/renderer"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var version = "v0.1.0"

var rootCmd = &cobra.Command{
	Use:   "gotodo",
	Short: color.HiCyanString("GoTodo - A Minimalist Terminal Task Manager"),
	Long: color.HiWhiteString(`GoTodo is a minimalist yet powerful terminal task manager,
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
