package main

import (
	"encoding/csv"
	"fmt"
	"github.com/alexflint/go-arg"
	"github.com/jedib0t/go-pretty/v6/table"
	"io"
	"log"
	"os"
	"strings"
)

type args struct {
	Header bool    `arg:"-H" help:"Treats the first row as a header row"`
	Footer bool    `arg:"-F" help:"Treats the last row as a footer row"`
	Style  *string `arg:"-S" help:"Set the table style: DEFAULT/DOUBLE/LIGHT/ROUNDED/BOLD/COLORED_BRIGHT/COLORED_DARK"`
	File   *string `arg:"positional"`
}

var styles = map[string]table.Style{
	"DEFAULT":        table.StyleDefault,
	"DOUBLE":         table.StyleDouble,
	"LIGHT":          table.StyleLight,
	"ROUNDED":        table.StyleRounded,
	"BOLD":           table.StyleBold,
	"COLORED_BRIGHT": table.StyleColoredBright,
	"COLORED_DARK":   table.StyleColoredDark,
}

func (args) Description() string {
	return "Reads a CSV input and outputs it formatted as a table"
}

func main() {
	var args args
	p := arg.MustParse(&args)
	var inputReader io.Reader
	if args.File == nil {
		fi, _ := os.Stdin.Stat()
		if (fi.Mode() & os.ModeCharDevice) != 0 {
			p.Fail("No input given")
		}
		inputReader = os.Stdin
	} else {
		file, err := os.Open(*args.File)
		if err != nil {
			p.Fail(fmt.Sprint("Failed to read input file:\n\t", err))
		}
		inputReader = file
	}
	reader := csv.NewReader(inputReader)
	records, err := reader.ReadAll()
	log.SetFlags(0)
	if err != nil {
		log.Fatal("Failed to read CSV:\n\t", err)
	}

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	row := 0
	if args.Header {
		t.AppendHeader(toRow(records[0]))
		row++
	}

	if args.Style != nil {
		style, exists := styles[strings.ToUpper(*args.Style)]
		if exists {
			t.SetStyle(style)
		} else {
			p.Fail(fmt.Sprint("Unknown style: ", *args.Style))
		}
	}

	rowCount := len(records)
	for ; row < rowCount; row++ {
		appendFunc := t.AppendRow
		if row == rowCount-1 && args.Footer {
			appendFunc = t.AppendFooter
		}
		appendFunc(toRow(records[row]))
	}
	t.Render()
}

func toRow(rowStr []string) table.Row {
	r := make(table.Row, len(rowStr))
	for i, v := range rowStr {
		r[i] = v
	}
	return r
}
