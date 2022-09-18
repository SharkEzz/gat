package printer

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"golang.org/x/term"
)

type tableParams struct {
	maxFirstColumnContentSize   int
	firstColumnGap              int
	firstColumnSize             int
	fistColumnSeparatorPosition int
	secondColumnStart           int
}

const (
	line      = "─"
	corner    = "┬"
	cornerTop = "┴"
	vertical  = "│"
	middle    = "┼"
	space     = " "

	black = "\033[0;30m"
	reset = "\033[0m"
)

func bold(s string) string {
	return fmt.Sprintf("\033[1m%s\033[0m", s)
}

func printHeader(filePath string, params *tableParams) error {
	fmt.Print(black)

	termWidth, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		return err
	}

	for i := 0; i < termWidth; i++ {
		if i == params.firstColumnSize {
			fmt.Print(corner)
		} else {
			fmt.Print(line)
		}
	}

	for i := 0; i < params.firstColumnSize; i++ {
		fmt.Print(space)
	}

	fmt.Print(vertical)

	fmt.Print("\033[0m")
	fmt.Print(" File: ")
	fmt.Println(bold(filePath))
	fmt.Print("\033[0;30m")
	for i := 0; i < termWidth; i++ {
		if i == params.firstColumnSize {
			fmt.Print(middle)
		} else {
			fmt.Print(line)
		}
	}

	fmt.Print(reset)

	return nil
}

func printFooter(params *tableParams) error {
	fmt.Println()

	fmt.Print(black)

	termWidth, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		return err
	}

	for i := 0; i < termWidth; i++ {
		if i == params.firstColumnSize {
			fmt.Print(cornerTop)
		} else {
			fmt.Print(line)
		}
	}

	fmt.Print(reset)

	return nil
}

func printFileLine(lineNumber int, line string, params *tableParams) {
	fmt.Print(black)

	lineNumberCharLength := len(strconv.FormatInt(int64(lineNumber), 10))

	correctedGap := 0

	if lineNumberCharLength < params.maxFirstColumnContentSize {
		for i := 0; i < params.maxFirstColumnContentSize-lineNumberCharLength; i++ {
			correctedGap++
		}
	}

	for i := 0; i < (params.firstColumnGap + correctedGap); i++ {
		fmt.Print(space)
	}

	fmt.Print(lineNumber)
	for i := 0; i < params.firstColumnGap; i++ {
		fmt.Print(space)
	}
	fmt.Printf("%s ", vertical)

	fmt.Print(reset)
	fmt.Print(strings.ReplaceAll(line, "\t", "    "))
}

func Print(content *string, filePath string) error {
	maxFirstColumnContentSize := len(strconv.FormatInt(int64(strings.Count(*content, "\n")), 10))
	firstColumnGap := 3
	firstColumnSize := (2 * firstColumnGap) + maxFirstColumnContentSize
	separatorPosition := firstColumnSize
	secondColumnStart := separatorPosition + 2

	params := tableParams{
		maxFirstColumnContentSize,
		firstColumnGap,
		firstColumnSize,
		separatorPosition,
		secondColumnStart,
	}

	if err := printHeader(filePath, &params); err != nil {
		return err
	}

	for lineNumber, line := range strings.SplitAfter(*content, "\n") {
		printFileLine(lineNumber+1, line, &params)
	}

	if err := printFooter(&params); err != nil {
		return err
	}

	return nil
}
