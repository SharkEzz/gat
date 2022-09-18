package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/SharkEzz/gat/pkg/printer"
	"github.com/SharkEzz/gat/utils"
)

func main() {
	noPaging := flag.Bool("no-paging", false, "Disable paging")

	flag.Parse()

	file := flag.Arg(0)
	if file == "" {
		utils.ExitWithError("please provide a file")
	}

	extension := utils.GetExtension(file)

	fileContent, err := os.ReadFile(file)
	if err != nil {
		utils.ExitWithError(err.Error())
	}

	contentStr := string(fileContent)

	output, err := printer.Print(&contentStr, file, extension)
	if err != nil {
		utils.ExitWithError(err.Error())
	}

	if *noPaging {
		fmt.Println(output)
		os.Exit(0)
	}

	cmd := exec.Command("less", "-R")
	cmd.Stdin = strings.NewReader(output)
	cmd.Stdout = os.Stdout
	err = cmd.Run()
	if err != nil {
		utils.ExitWithError(fmt.Sprintf("failed to run less: %s", err.Error()))
	}
}
