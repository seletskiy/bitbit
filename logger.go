package main

import "log"

var logger Logger = Nothing

type Logger int

const (
	Nothing = iota
	Debug
)

func (currentLevel Logger) Log(level int, format string, data ...interface{}) {
	if int(currentLevel) < level {
		return
	}

	log.Printf(format, data...)
}
