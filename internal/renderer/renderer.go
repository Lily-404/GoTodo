package renderer

import (
	"hacknote/internal/storage"

	"github.com/fatih/color"
)

func ShowBanner() {
	bold := color.New(color.FgHiCyan, color.Bold) // 使用青色作为标题
	gray := color.New(color.FgHiBlack)            // 使用深灰色作为分隔线

	bold.Println("\n  GoTodo")
	gray.Println("  ─────────────────────────") // 简化分隔线
}

func RenderNotes(notes []storage.Note, showAll bool, filterPriority string) {
	// 定义统一的颜色方案
	checkbox := color.New(color.FgHiBlue) // 方框使用蓝色
	task := color.New(color.FgHiWhite)    // 所有未完成任务统一使用亮白色
	done := color.New(color.FgHiBlack)    // 已完成任务使用深灰色
	stats := color.New(color.FgHiGreen)   // 统计信息使用绿色

	// 分类任务
	var unfinishedNotes []storage.Note
	var finishedNotes []storage.Note

	for _, note := range notes {
		if note.Status == "done" {
			finishedNotes = append(finishedNotes, note)
		} else {
			unfinishedNotes = append(unfinishedNotes, note)
		}
	}

	// 显示未完成任务，使用统一的颜色
	for _, note := range unfinishedNotes {
		checkbox.Print("  □ ")
		task.Println(note.Content)
	}

	// 显示已完成任务
	if len(finishedNotes) > 0 {
		done.Println("\n  ─────────────────────────")
		for _, note := range finishedNotes {
			done.Print("  ■ ")
			done.Println(note.Content)
		}
	}

	// 显示统计信息
	if len(notes) > 0 {
		stats.Printf("\n  %d 个任务 (%d 已完成)\n", len(notes), len(finishedNotes))
	}
}
