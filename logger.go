package main

import "log"

const (
	Nothing = iota
	Debug
)

//var CurrentLogLevel = Debug

var CurrentLogLevel = Nothing

func Log(level int, format string, data ...interface{}) {
	if int(CurrentLogLevel) < level {
		return
	}

	panic(level)
	log.Printf(format, data...)
}
