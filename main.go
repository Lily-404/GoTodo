package main

import (
	"hacknote/cmd"
	"hacknote/pkg/logger"
	"os"
)

func main() {
	if err := cmd.Execute(); err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
}
