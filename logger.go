package main

import "log"

const (
	Nothing = iota
	Debug
)

var CurrentLogLevel = Nothing

func Log(level int, format string, data ...interface{}) {
	if int(CurrentLogLevel) < level {
		return
	}

	log.Printf(format, data...)
}
