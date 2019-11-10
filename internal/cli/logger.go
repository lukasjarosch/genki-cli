package cli

import (
	"fmt"

	"github.com/fatih/color"
)

const (
	format = "%s: %s\n"
)

type cliLogger struct {
}

func NewCliLogger() *cliLogger {
	return &cliLogger{
	}
}

func (l *cliLogger) Infof(message string, args ...interface{}) {
	cyan := color.New(color.FgCyan).SprintfFunc()
	if len(args) > 0 {
		fmt.Printf(format, cyan("INFO"), fmt.Sprintf(message, args...))
		return
	}
	fmt.Printf(format, cyan("INFO"), message)
}

func (l *cliLogger) Successf(message string, args ...interface{}) {
	cyan := color.New(color.FgGreen).SprintfFunc()
	if len(args) > 0 {
		fmt.Printf(format, cyan("SUCC"), fmt.Sprintf(message, args...))
		return
	}
	fmt.Printf(format, cyan("SUCC"), message)
}

func (l *cliLogger) Warningf(message string, args ...interface{}) {
	yellow := color.New(color.FgHiYellow).SprintfFunc()
	if len(args) > 0 {
		fmt.Printf(format, yellow("WARN"), fmt.Sprintf(message, args...))
		return
	}
	fmt.Printf(format, yellow("WARN"), message)
}

func (l *cliLogger) Errorf(message string, args ...interface{}) {
	red := color.New(color.FgHiRed).SprintfFunc()
	if len(args) > 0 {
		fmt.Printf(format, red("ERRO"), fmt.Sprintf(message, args...))
		return
	}
	fmt.Printf(format, red("ERRO"), message)
}

func (l *cliLogger) Fatalf(message string, args ...interface{}) {
	red := color.New(color.FgWhite, color.BgRed).SprintfFunc()
	if len(args) > 0 {
		fmt.Printf(format, red("FATA"), fmt.Sprintf(message, args...))
		return
	}
	fmt.Printf(format, red("FATA"), message)
}

