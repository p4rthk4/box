package logx

import (
	"fmt"
	"github.com/fatih/color"
	"gopkg.in/yaml.v3"
)

func LogError(message string, err error) {
	red := color.New(color.FgRed).SprintFunc()
	bold := color.New(color.Bold).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()
	blue := color.New(color.FgBlue).SprintFunc()

	fmt.Printf("%s %s: %v\n", bold(red("ERROR")), message, err.Error())
	fmt.Printf("%s:\n%s\n", yellow("Detailed error"), blue(detailedError(err)))
	// log.Printf("%s:\n%s", yellow("Stack trace"), blue(string(debug.Stack())))
}

func detailedError(err error) string {
	switch e := err.(type) {
	case *yaml.TypeError:
		return fmt.Sprintf("YAML parsing error: %s", e.Errors)
	default:
		return err.Error()
	}
}
