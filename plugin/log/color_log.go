package log

import (
	"time"

	"github.com/fatih/color"
)

func Info(s string, a ...interface{}) {
	grey := color.New(color.FgHiBlack).PrintfFunc()
	grey(time.Now().Format("2016-01-02 15:04:05")+" "+s+"\n", a...)
}

func Warn(s string, a ...interface{}) {
	yellow := color.New(color.FgYellow).PrintfFunc()
	yellow(time.Now().Format("2016-01-02 15:04:05")+" "+s+"\n", a...)
}

func Succ(s string, a ...interface{}) {
	cyan := color.New(color.FgGreen).PrintfFunc()
	cyan(time.Now().Format("2016-01-02 15:04:05")+" "+s+"\n", a...)
}

func Err(s string, a ...interface{}) {
	red := color.New(color.FgRed).PrintfFunc()
	red(time.Now().Format("2016-01-02 15:04:05")+" "+s+"\n", a...)
}

func Tx(s string, a ...interface{}) {
	cyan := color.New(color.FgHiCyan).PrintfFunc()
	cyan(time.Now().Format("2016-01-02 15:04:05")+" "+s+"\n", a...)
}
