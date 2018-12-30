package log

import "github.com/fatih/color"

func Info(s string, a ...interface{}) {
	grey := color.New(color.FgHiBlack).PrintfFunc()
	grey(s+"\n", a...)
}

func Warn(s string, a ...interface{}) {
	yellow := color.New(color.FgYellow).PrintfFunc()
	yellow(s+"\n", a...)
}

func Succ(s string, a ...interface{}) {
	cyan := color.New(color.FgGreen).PrintfFunc()
	cyan(s+"\n", a...)
}

func Err(s string, a ...interface{}) {
	red := color.New(color.FgRed).PrintfFunc()
	red(s+"\n", a...)
}

func Tx(s string, a ...interface{}) {
	cyan := color.New(color.FgHiCyan).PrintfFunc()
	cyan(s+"\n", a...)
}
