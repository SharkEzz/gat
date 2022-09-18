package printer

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/alecthomas/chroma/quick"
	"golang.org/x/term"
)

type tableParams struct {
	maxFirstColumnContentSize   int
	firstColumnGap              int
	firstColumnSize             int
	fistColumnSeparatorPosition int
	secondColumnStart           int
	extension                   string
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

func getHeader(filePath string, params *tableParams) (string, error) {
	buffer := &strings.Builder{}

	fmt.Fprint(buffer, black)

	termWidth, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		return "", err
	}

	for i := 0; i < termWidth; i++ {
		if i == params.firstColumnSize {
			fmt.Fprint(buffer, corner)
		} else {
			fmt.Fprint(buffer, line)
		}
	}

	for i := 0; i < params.firstColumnSize; i++ {
		fmt.Fprint(buffer, space)
	}

	fmt.Fprint(buffer, vertical)

	fmt.Fprint(buffer, reset)
	fmt.Fprint(buffer, " File: ")
	fmt.Fprintln(buffer, bold(filePath))
	fmt.Fprint(buffer, black)
	for i := 0; i < termWidth; i++ {
		if i == params.firstColumnSize {
			fmt.Fprint(buffer, middle)
		} else {
			fmt.Fprint(buffer, line)
		}
	}

	fmt.Fprint(buffer, reset)

	return buffer.String(), nil
}

func getFooter(params *tableParams) (string, error) {
	buffer := &strings.Builder{}

	fmt.Fprintln(buffer)

	fmt.Fprint(buffer, black)

	termWidth, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		return "", err
	}

	for i := 0; i < termWidth; i++ {
		if i == params.firstColumnSize {
			fmt.Fprint(buffer, cornerTop)
		} else {
			fmt.Fprint(buffer, line)
		}
	}

	fmt.Fprint(buffer, reset)

	return buffer.String(), nil
}

func printFileLine(lineNumber int, line string, params *tableParams) string {
	buffer := &strings.Builder{}

	fmt.Fprint(buffer, black)

	lineNumberCharLength := len(strconv.FormatInt(int64(lineNumber), 10))

	correctedGap := 0

	if lineNumberCharLength < params.maxFirstColumnContentSize {
		for i := 0; i < params.maxFirstColumnContentSize-lineNumberCharLength; i++ {
			correctedGap++
		}
	}

	for i := 0; i < (params.firstColumnGap + correctedGap); i++ {
		fmt.Fprint(buffer, space)
	}

	fmt.Fprint(buffer, lineNumber)
	for i := 0; i < params.firstColumnGap; i++ {
		fmt.Fprint(buffer, space)
	}
	fmt.Fprintf(buffer, "%s ", vertical)

	fmt.Fprint(buffer, reset)

	if params.extension != "" {
		quick.Highlight(buffer, line, params.extension, "terminal16m", "monokai")
	} else {
		fmt.Fprint(buffer, line)
	}

	return buffer.String()
}

func Print(content *string, filePath, extension string) (string, error) {
	maxFirstColumnContentSize := len(strconv.FormatInt(int64(strings.Count(*content, "\n")), 10))
	firstColumnGap := 3
	firstColumnSize := (2 * firstColumnGap) + maxFirstColumnContentSize
	separatorPosition := firstColumnSize
	secondColumnStart := separatorPosition + 2

	lineCount := strings.Count(*content, "\n")
	if lineCount > 10000 {
		extension = "" // skip coloration if file is too big
	}

	params := tableParams{
		maxFirstColumnContentSize,
		firstColumnGap,
		firstColumnSize,
		separatorPosition,
		secondColumnStart,
		extension,
	}

	buffer := &strings.Builder{}

	header, err := getHeader(filePath, &params)
	if err != nil {
		return "", nil
	}
	fmt.Fprint(buffer, header)

	for lineNumber, line := range strings.SplitAfter(*content, "\n") {
		fmt.Fprint(buffer, printFileLine(lineNumber+1, line, &params))
	}

	footer, err := getFooter(&params)
	if err != nil {
		return "", err
	}
	fmt.Fprint(buffer, footer)

	return buffer.String(), nil
}
