package cli

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

const (
	format = "%s: %s\n"
)

var log *cliLogger

func init() {
	log = NewCliLogger()
}

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
		fmt.Printf(format, red("ERR "), fmt.Sprintf(message, args...))
		return
	}
	fmt.Printf(format, red("ERR "), message)
}

func (l *cliLogger) Fatalf(message string, args ...interface{}) {
	red := color.New(color.FgWhite, color.BgRed).SprintfFunc()
	if len(args) > 0 {
		fmt.Printf(format, red("FATA"), fmt.Sprintf(message, args...))
		os.Exit(1)
	}
	fmt.Printf(format, red("FATA"), message)
	os.Exit(1)
}


func Infof(message string, args ...interface{}) {
	log.Infof(message, args...)
}

func Successf(message string, args ...interface{}) {
	log.Successf(message, args...)
}

func Warningf(message string, args ...interface{}) {
	log.Warningf(message, args...)
}

func Errorf(message string, args ...interface{}) {
	log.Errorf(message, args...)
}

func Fatalf(message string, args ...interface{}) {
	log.Fatalf(message, args...)
}
