package cmd

import (
	"fmt"

	"github.com/Lily-404/todo/internal/config"
	"github.com/Lily-404/todo/internal/i18n"

	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var langCmd = &cobra.Command{
	Use:     "lang",
	Aliases: []string{"language"},
	Short:   i18n.GetMessage(config.GetConfig().Language, "cmd_lang_short"),
	RunE: func(cmd *cobra.Command, args []string) error {
		// 创建语言选择提示
		prompt := promptui.Select{
			Label: i18n.GetMessage(config.GetConfig().Language, "select_language"),
			Items: []string{
				i18n.GetMessage(config.GetConfig().Language, "language_english"),
				i18n.GetMessage(config.GetConfig().Language, "language_chinese"),
			},
			Templates: &promptui.SelectTemplates{
				Label:    "{{ . }}",
				Active:   "➤ {{ . | cyan }}",
				Inactive: "  {{ . }}",
				Selected: "✓ {{ . | green }}",
			},
		}

		idx, _, err := prompt.Run()
		if err != nil {
			return fmt.Errorf(i18n.GetMessage(config.GetConfig().Language, "language_select_failed"), err)
		}

		// 根据选择设置语言
		lang := "en"
		if idx == 1 {
			lang = "zh"
		}

		// 更新配置
		config.DefaultConfig.Language = lang

		// 保存配置到文件
		if err := config.SaveConfig(); err != nil {
			return fmt.Errorf("保存语言配置失败: %v", err)
		}

		// 显示成功消息
		color.Green(i18n.GetMessage(lang, "language_changed"))
		return nil
	},
}

func init() {
	rootCmd.AddCommand(langCmd)
}
