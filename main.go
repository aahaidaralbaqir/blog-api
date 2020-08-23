package main

import (
	"go-crash-course/cmd"
)

const (
	PORT int = 8000
)

func main() {
	application := new(cmd.Application)
	application.ConfigureDatabase()
	application.Start(PORT)
}
