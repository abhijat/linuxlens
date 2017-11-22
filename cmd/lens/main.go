package main

import (
	"linuxlens"
	"fmt"
)

func main() {
	processes := linuxlens.ListProcFiles()
	for _, process := range processes {
		fmt.Println(process)
	}
}
