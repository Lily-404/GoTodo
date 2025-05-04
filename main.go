package main

import (
	"github.com/Lily-404/todo/cmd"
	"github.com/Lily-404/todo/pkg/logger"
	"os"
)

func main() {
	if err := cmd.Execute(); err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
}
