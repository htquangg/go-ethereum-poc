package visualize

import (
	"os"
	"strings"

	"github.com/jedib0t/go-pretty/table"
	"github.com/mattn/go-isatty"
	"github.com/muesli/ansi"
	"github.com/muesli/reflow/truncate"
	"github.com/rs/zerolog/log"
	"golang.org/x/term"
)

type TableOptions struct {
	Title string
}


const (
	// combined width of the table borders and padding
	borderWidths = 10
	// char to indicate that a string has been truncated
	ellipsis = "…"
)

func Table(headers [2]string, rows [][2]string) {
	shouldTruncate := isatty.IsTerminal(os.Stdout.Fd())

	// This will return an error if we're not in a terminal or
	// if the terminal is a cygwin terminal like Git Bash.
	width, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		if shouldTruncate {
			log.Error().Msgf("error getting terminal size: %s", err)
		} else {
			log.Debug().Err(err)
		}
	}

	longestSecretName, longestSecretType := getLongestValues(append(rows, headers))
	availableWidth := width - longestSecretName - longestSecretType - borderWidths
	if availableWidth < 0 {
		availableWidth = 0
	}

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetStyle(table.StyleLight)

	t.Style().Options.DrawBorder = true
	t.Style().Options.SeparateHeader = true
	t.Style().Options.SeparateColumns = true

	tableHeaders := table.Row{}
	for _, header := range headers {
		tableHeaders = append(tableHeaders, header)
	}

	t.AppendHeader(tableHeaders)
	for _, row := range rows {
		tableRow := table.Row{}
		for i, val := range row {
			if i == 1 && stringWidth(val) > availableWidth && shouldTruncate {
				val = truncate.StringWithTail(val, uint(availableWidth), ellipsis)
			}
			tableRow = append(tableRow, val)
		}
		t.AppendRow(tableRow)
	}

	t.Render()
}

func getLongestValues(rows [][2]string) (longestSecretName, longestSecretType int) {
	for _, row := range rows {
		if len(row[0]) > longestSecretName {
			longestSecretName = stringWidth(row[0])
		}
		if len(row[1]) > longestSecretType {
			longestSecretType = stringWidth(row[1])
		}
	}
	return
}

func stringWidth(str string) (width int) {
	for _, l := range strings.Split(str, "\n") {
		w := ansi.PrintableRuneWidth(l)
		if w > width {
			width = w
		}
	}
	return width
}
