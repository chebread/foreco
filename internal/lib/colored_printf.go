package lib

import "github.com/fatih/color"

var BoldCyanPrintf = color.New(color.FgCyan, color.Bold).PrintfFunc()
var RedCyanPrintf = color.New(color.FgRed, color.Bold).PrintfFunc()
