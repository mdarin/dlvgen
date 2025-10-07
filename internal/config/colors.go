package config

import "github.com/fatih/color"

var (
	Blue    = color.New(color.FgBlue).SprintFunc()
	Green   = color.New(color.FgGreen).SprintFunc()
	Red     = color.New(color.FgRed).SprintFunc()
	Yellow  = color.New(color.FgYellow).SprintFunc()
	Magenta = color.New(color.FgMagenta).SprintFunc()
)
