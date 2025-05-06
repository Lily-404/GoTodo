package i18n

import "fmt"

var messages = map[string]map[string]string{
	"en": {
		"cmd_add_short":        "Add a new note",
		"cmd_clean_short":      "Clean completed tasks",
		"cmd_done_short":       "Mark task as completed",
		"cmd_list_short":       "List all tasks",
		"cmd_priority_short":   "Modify task priority",
		"cmd_remove_short":     "Remove specified task",
		"cmd_test_short":       "Test and build project",
		"cmd_help_short":       "Help about any command",
		"cmd_completion_short": "Generate the autocompletion script",
		"cmd_root_short":       "Todo - A Minimalist Terminal Task Manager",
		"cmd_root_long": `Todo is a minimalist yet powerful terminal task manager,
focused on helping you efficiently manage your todos.

Command Aliases:
  add (a)  - Add task
  list (l) - List tasks
  done (d) - Complete task
  clean (c) - Clean tasks`,
		"cmd_lang_short":          "Set default language",
		"select_language":         "Select default language",
		"language_english":        "English",
		"language_chinese":        "Chinese",
		"language_select_failed":  "Failed to select language: %v",
		"language_changed":        "Language has been changed to English",
		"building_project":        "Building project...",
		"build_failed":            "Build failed",
		"build_success":           "Build successful",
		"starting_tests":          "Starting functional tests...",
		"show_pending_only":       "Show pending tasks only",
		"filter_by_priority":      "Filter by priority (high/normal/low)",
		"flag_title":              "Note title",
		"flag_priority":           "Priority (high/normal/low)",
		"flag_due_date":           "Due date (YYYY-MM-DD)",
		"total_tasks":             "Total: %d tasks (%d completed)",
		"help_flag_short":         "help for todo",
		"version_flag_short":      "version for todo",
		"available_commands":      "Available Commands:",
		"flags":                   "Flags:",
		"help_usage":              "Use \"{{.CommandPath}} [command] --help\" for more information about a command.",
		"add_task":                "Add task",
		"list_tasks":              "List tasks",
		"complete_task":           "Complete task",
		"clean_tasks":             "Clean tasks",
		"select_priority":         "Select task priority",
		"no_pending_tasks":        "No pending tasks",
		"task_completed":          "Task completed: %s",
		"current_tasks":           "Current tasks:",
		"select_task":             "Select task to complete",
		"delete_task":             "Delete task",
		"modify_priority":         "Modify task priority",
		"new_priority":            "Select new priority",
		"select_new_priority":     "Select new priority",
		"no_completed_tasks":      "No completed tasks to clean",
		"tasks_to_clean":          "Completed tasks to be cleaned:",
		"cleaned_tasks":           "Cleaned %d completed tasks",
		"priority_updated":        "Task priority updated: %s -> %s",
		"no_tasks":                "No tasks...",
		"banner_subtitle":         "Focus on what matters.",
		"priority_select_failed":  "Failed to select priority: %v",
		"select_task_priority":    "Select task to modify priority",
		"priority_low":            "Low",
		"priority_normal":         "Normal",
		"priority_high":           "High",
		"select_task_to_remove":   "Select task to remove",
		"select_task_to_complete": "Select task to complete",
		"task_select_failed":      "Failed to select task: %v",
		"task_deleted":            "Task deleted: %s",
	},
	"zh": {
		"cmd_add_short":        "添加新笔记",
		"cmd_clean_short":      "清理已完成的任务",
		"cmd_done_short":       "标记任务为已完成",
		"cmd_list_short":       "列出所有任务",
		"cmd_priority_short":   "修改任务优先级",
		"cmd_remove_short":     "删除指定的任务",
		"cmd_test_short":       "测试并构建项目",
		"cmd_help_short":       "显示命令帮助信息",
		"cmd_completion_short": "生成自动补全脚本",
		"cmd_root_short":       "Todo - 极简终端任务管理器",
		"cmd_root_long": `Todo 是一个极简但功能强大的终端任务管理器，
专注于帮助你高效管理待办事项。

命令别名：
  add (a)  - 添加任务
  list (l) - 列出任务
  done (d) - 完成任务
  clean (c) - 清理任务`,
		"cmd_lang_short":          "选择默认语言",
		"select_language":         "选择默认语言",
		"language_english":        "English",
		"language_chinese":        "Chinese",
		"language_select_failed":  "选择语言失败: %v",
		"language_changed":        "语言已切换为中文",
		"building_project":        "正在构建项目...",
		"build_failed":            "构建失败",
		"build_success":           "构建成功",
		"starting_tests":          "开始功能测试...",
		"show_pending_only":       "仅显示待处理任务",
		"filter_by_priority":      "按优先级筛选（高/中/低）",
		"flag_title":              "任务标题",
		"flag_priority":           "优先级（高/中/低）",
		"flag_due_date":           "截止日期（YYYY-MM-DD）",
		"total_tasks":             "总计：%d 个任务（%d 已完成）",
		"help_flag_short":         "显示帮助信息",
		"version_flag_short":      "显示版本信息",
		"available_commands":      "可用命令：",
		"flags":                   "标志：",
		"help_usage":              "使用 \"{{.CommandPath}} [命令] --help\" 获取关于命令的更多信息。",
		"add_task":                "添加任务",
		"list_tasks":              "列出任务",
		"complete_task":           "完成任务",
		"clean_tasks":             "清理任务",
		"select_priority":         "选择任务优先级",
		"no_pending_tasks":        "没有待处理的任务",
		"task_completed":          "任务已完成: %s",
		"current_tasks":           "当前任务列表：",
		"select_task":             "选择要完成的任务",
		"delete_task":             "删除任务",
		"modify_priority":         "修改任务优先级",
		"no_completed_tasks":      "没有已完成的任务需要清理",
		"tasks_to_clean":          "将要清理的已完成任务：",
		"cleaned_tasks":           "已清理 %d 个已完成的任务",
		"select_new_priority":     "选择新的优先级",
		"priority_updated":        "任务优先级已更新: %s -> %s",
		"no_tasks":                "没有任务...",
		"banner_subtitle":         "专注于重要的事。",
		"priority_select_failed":  "选择优先级失败: %v",
		"select_task_priority":    "选择要修改优先级的任务",
		"priority_low":            "低",
		"priority_normal":         "中",
		"priority_high":           "高",
		"select_task_to_remove":   "选择要删除的任务",
		"select_task_to_complete": "选择要完成的任务",
		"task_select_failed":      "选择任务失败: %v",
		"task_deleted":            "已删除任务: %s",
	},
}

func GetMessage(lang, key string, args ...interface{}) string {
	if msg, ok := messages[lang][key]; ok {
		if len(args) > 0 {
			return fmt.Sprintf(msg, args...)
		}
		return msg
	}
	return key
}
