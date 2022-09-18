package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"

	"github.com/SharkEzz/gat/pkg/printer"
	"github.com/SharkEzz/gat/utils"
)

func getExtension(filePath string) string {
	regex := regexp.MustCompile(`\.([a-zA-Z]+)`)

	result := regex.FindAllStringSubmatch(filePath, -1)

	if len(result) == 0 {
		return ""
	}

	return result[0][1]
}

func main() {
	noPaging := flag.Bool("no-paging", false, "Disable paging")

	flag.Parse()

	file := flag.Arg(0)
	if file == "" {
		utils.ExitWithError("please provide a file")
	}

	extension := getExtension(file)

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
