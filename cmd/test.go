package cmd

import (
	"os"
	"os/exec"
	"time"

	"github.com/Lily-404/todo/internal/config"
	"github.com/Lily-404/todo/internal/i18n"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var testCmd = &cobra.Command{
	Use:     "test",
	Aliases: []string{"t"},
	Short:   i18n.GetMessage(config.GetConfig().Language, "cmd_test_short"),
	RunE: func(cmd *cobra.Command, args []string) error {
		info := color.New(color.FgHiCyan)
		success := color.New(color.FgHiGreen)
		fail := color.New(color.FgHiRed)

		info.Println("\n  " + i18n.GetMessage(config.GetConfig().Language, "building_project"))
		buildCmd := exec.Command("go", "build", "-o", "todo")
		buildCmd.Stdout = os.Stdout
		buildCmd.Stderr = os.Stderr
		if err := buildCmd.Run(); err != nil {
			fail.Println("  " + i18n.GetMessage(config.GetConfig().Language, "build_failed"))
			return err
		}
		success.Println("  " + i18n.GetMessage(config.GetConfig().Language, "build_success"))

		info.Println("\n  " + i18n.GetMessage(config.GetConfig().Language, "starting_tests"))

		// 清理旧数据
		runCommand("rm", "-f", "notes/notes.json")

		// 测试添加任务
		info.Println("1. Testing task addition:")
		runCommand("./todo", "add", "Complete project docs", "-p", "high")
		runCommand("./todo", "add", "Prepare meeting presentation", "-p", "normal")
		runCommand("./todo", "add", "Reply to emails", "-p", "low")
		time.Sleep(time.Second)

		// 测试列表显示
		info.Println("\n2. Testing task list:")
		runCommand("./todo", "list")
		time.Sleep(time.Second)

		success.Println("\n  Tests completed!")
		return nil
	},
}

func runCommand(name string, args ...string) {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

func init() {
	rootCmd.AddCommand(testCmd)
}
