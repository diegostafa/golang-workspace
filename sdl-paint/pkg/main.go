package main

import (
	"os"
	"strconv"
)

func grabArgs() (int, int) {
	args := os.Args[1:]
	width, err := strconv.Atoi(args[0])
	SmartPanic(err)
	height, err := strconv.Atoi(args[1])
	SmartPanic(err)
	return width, height
}

func main() {
	w, h := grabArgs()
	app := new(controller)
	app.Exec(w, h)
}
