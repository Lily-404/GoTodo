package renderer

import (
	"fmt"
	"github.com/Lily-404/todo/internal/storage"
	"sort"

	"github.com/fatih/color"
)

func ShowBanner() {
	logo := color.New(color.FgHiCyan, color.Bold)
	border := color.New(color.FgHiBlack)

	logo.Print("\n  ┌────────────────────────┐\n")
	logo.Print("  │         Go Todo        │")
	logo.Print("\n  └────────────────────────┘\n")
	border.Println("    Focus on What Matters.")
}

func RenderNotes(notes []storage.Note, showAll bool, filterPriority string) {
	// 定义更简洁的颜色方案
	title := color.New(color.FgHiYellow, color.Bold)
	done := color.New(color.FgHiBlack)
	date := color.New(color.FgHiMagenta)
	divider := color.New(color.FgHiBlack)

	// 优先级只用颜色区分
	priority := map[string]*color.Color{
		"high":   color.New(color.FgHiRed),
		"normal": color.New(color.FgHiBlue),
		"low":    color.New(color.FgHiWhite),
	}

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

	// 获取未完成任务后，按优先级排序
	priorityWeight := map[string]int{
		"high":   3,
		"normal": 2,
		"low":    1,
	}

	// 对未完成任务进行排序
	sort.Slice(unfinishedNotes, func(i, j int) bool {
		weightI := priorityWeight[unfinishedNotes[i].Priority]
		weightJ := priorityWeight[unfinishedNotes[j].Priority]
		if weightI == weightJ {
			// 如果优先级相同，按创建时间排序（新的在前）
			return unfinishedNotes[i].CreatedAt.After(unfinishedNotes[j].CreatedAt)
		}
		return weightI > weightJ
	})

	// 显示分隔线和未完成任务
	divider.Println("  ──────────────────────────")
	for _, note := range unfinishedNotes {
		p := priority[note.Priority]
		if p == nil {
			p = priority["normal"]
		}

		p.Print("  ● ")
		if note.Title != "" {
			title.Printf("%s: ", note.Title)
		}
		p.Print(note.Content)
		if note.DueDate != "" {
			fmt.Print(" ")
			date.Printf("📅 %s", note.DueDate)
		}
		fmt.Println()
	}

	// 显示已完成任务
	if len(finishedNotes) > 0 {
		divider.Println("  ──────────────────────────")
		for _, note := range finishedNotes {
			done.Printf("  ✓ %s\n", note.Content)
		}
	}

	// 显示进度统计
	if len(notes) > 0 {
		divider.Println("  ──────────────────────────")
		ShowProgressBar(len(notes), len(finishedNotes))
		// 使用更柔和的颜色组合来显示统计信息
		totalCount := color.New(color.FgHiBlue)
		completedCount := color.New(color.FgHiGreen)
		totalCount.Printf("\n  总计：")
		completedCount.Printf("%d", len(notes))
		totalCount.Printf(" 个任务（")
		completedCount.Printf("%d", len(finishedNotes))
		totalCount.Printf(" 已完成）\n")
	}
}

func ShowProgressBar(total, completed int) {
	width := 21 // 将宽度从30改为20，使显示更加紧凑
	filled := int(float64(completed) / float64(total) * float64(width))
	percent := int(float64(completed) / float64(total) * 100)

	// 定义新的颜色方案
	progress := color.New(color.FgHiCyan)
	remaining := color.New(color.FgHiBlack)
	percentage := color.New(color.FgHiCyan, color.Bold)

	fmt.Print("  ") // 保持缩进
	for i := 0; i < width; i++ {
		if i < filled {
			progress.Print("█")
		} else {
			remaining.Print("░")
		}
	}
	fmt.Print(" ") // 在进度条和百分比之间添加空格
	percentage.Printf("%d%%", percent)
}
