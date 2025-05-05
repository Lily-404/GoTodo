package main

import (
	"os"

	"github.com/Lily-404/todo/cmd"
	"github.com/Lily-404/todo/pkg/logger"
)

func main() {
	if err := cmd.Execute(); err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
}
