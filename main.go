package main

import (
	"os"
	"strings"

	"github.com/SharkEzz/gat/pkg/printer"
	"github.com/SharkEzz/gat/utils"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		utils.ExitWithError("please provide a file")
	}

	fileContent, err := os.ReadFile(args[0])
	if err != nil {
		utils.ExitWithError(err.Error())
	}

	contentStr := strings.TrimSuffix(string(fileContent), "\n")

	err = printer.Print(&contentStr, args[0])
	if err != nil {
		utils.ExitWithError(err.Error())
	}
}
