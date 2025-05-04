package logger

import "github.com/fatih/color"

var isVerbose bool

func SetVerbose(verbose bool) {
	isVerbose = verbose
}

func Info(msg string) {
	if isVerbose {
		color.Cyan("[INFO] %s", msg)
	}
}

func Error(msg string) {
	color.Red("[ERROR] %s", msg)
}

func Success(msg string) {
	color.Green("[SUCCESS] %s", msg)
}
