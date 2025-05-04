package cmd

import (
	"os"
	"os/exec"
	"time"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var testCmd = &cobra.Command{
	Use:     "test",
	Aliases: []string{"t"},
	Short:   "Test and build project",
	RunE: func(cmd *cobra.Command, args []string) error {
		info := color.New(color.FgHiCyan)
		success := color.New(color.FgHiGreen)
		fail := color.New(color.FgHiRed)

		// 构建项目
		info.Println("\n  Building project...")
		buildCmd := exec.Command("go", "build", "-o", "todo")
		buildCmd.Stdout = os.Stdout
		buildCmd.Stderr = os.Stderr
		if err := buildCmd.Run(); err != nil {
			fail.Println("  构建失败")
			return err
		}
		success.Println("  Build successful")

		// 测试各个命令
		info.Println("\n  Starting functional tests...")

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
