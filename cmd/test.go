package cmd

import (
	"os"
	"os/exec"
	"time"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "测试并构建项目",
	RunE: func(cmd *cobra.Command, args []string) error {
		info := color.New(color.FgHiCyan)
		success := color.New(color.FgHiGreen)
		fail := color.New(color.FgHiRed)

		// 构建项目
		info.Println("\n  正在构建项目...")
		buildCmd := exec.Command("go", "build", "-o", "gotodo")
		buildCmd.Stdout = os.Stdout
		buildCmd.Stderr = os.Stderr
		if err := buildCmd.Run(); err != nil {
			fail.Println("  构建失败")
			return err
		}
		success.Println("  构建成功")

		// 测试各个命令
		info.Println("\n  开始功能测试...\n")

		// 清理旧数据
		runCommand("rm", "-f", "notes/notes.json")

		// 测试添加任务
		info.Println("1. 测试添加任务:")
		runCommand("./gotodo", "add", "完成项目文档", "-p", "high")
		runCommand("./gotodo", "add", "准备周会演示", "-p", "normal")
		runCommand("./gotodo", "add", "回复邮件", "-p", "low")
		time.Sleep(time.Second)

		// 测试列表显示
		info.Println("\n2. 测试任务列表:")
		runCommand("./gotodo", "list")
		time.Sleep(time.Second)

		success.Println("\n  测试完成!")
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
