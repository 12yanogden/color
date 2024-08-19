package colors

import (
	"os"
	"regexp"
	"runtime"
	"strings"

	slices "github.com/12yanogden/go-slices"
	"golang.org/x/term"
)

var COLORS = map[string]string{
	"RESET":  "\033[0m",
	"RED":    "\033[31m",
	"GREEN":  "\033[32m",
	"YELLOW": "\033[33m",
	"BLUE":   "\033[34m",
	"PURPLE": "\033[35m",
	"CYAN":   "\033[36m",
	"GRAY":   "\033[37m",
	"WHITE":  "\033[97m",
}

// Return true if colors should be used, else false
func isColorable() bool {
	return runtime.GOOS != "windows" && term.IsTerminal(int(os.Stdout.Fd()))
}

// Colorize the message given to green
func Green(msg string) string {
	if isColorable() {
		msg = COLORS["GREEN"] + msg + COLORS["RESET"]
	}

	return msg
}

// Colorize the message given to red
func Red(msg string) string {
	if isColorable() {
		msg = COLORS["RED"] + msg + COLORS["RESET"]
	}

	return msg
}

func StrLen(str string) int {
	for _, color := range COLORS {
		str = strings.ReplaceAll(str, color, "")
	}

	return len(str)
}

func Left(str string, width int) string {
	padCount := 0

	if width > StrLen(str) {
		padCount = width - StrLen(str)
	}

	str = str + strings.Repeat(" ", padCount)

	return str
}

func LocateAll(str string) []Location {
	locations := []Location{}
	colorPattern := regexp.MustCompile(`\[[0-9]+m`)
	indexSlices := colorPattern.FindAllIndex([]byte(str), -1)

	for _, indexSlice := range indexSlices {
		locations = append(locations, Location{
			Start: indexSlice[0] - 1,
			End:   indexSlice[1] - 1,
		})
	}

	return locations
}

func Decolor(str string) string {
	locations := LocateAll(str)

	for _, location := range locations {
		str = str[:location.Start] + str[location.End:]
	}

	return str
}

// Automatically increment an index to skip a color location
func Skip(i *int, locations *[]Location) bool {
	if len(*locations) > 0 && *i == (*locations)[0].Start {
		*i = (*locations)[0].End
		*locations = slices.Shift(*locations)

		return true
	}

	return false
}
