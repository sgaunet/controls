package results

import (
	"fmt"

	"github.com/fatih/color"
)

type Result struct {
	Title  string
	Result string
	Pass   bool
}

func (r *Result) GetResultHtml() string {
	if r.Pass {
		return fmt.Sprintf("<span style=\"color:green\">%s</span>", r.Result)
	} else {
		return fmt.Sprintf("<span style=\"color:red\">%s</span>", r.Result)
	}
}

func (r *Result) PrintToStdout() {
	var std *color.Color
	if r.Pass {
		std = color.New(color.FgGreen, color.Bold)
	} else {
		std = color.New(color.FgRed, color.Bold)
	}

	std.Printf("Title  : %s\n", r.Title)
	std.Printf("Result : %s\n", r.Result)
	fmt.Println("                             **** ")
}
