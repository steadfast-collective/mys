package main

import (
	"fmt"
	"mys/internal"
	"os"
)

func main() {
	parseCommand()
}

func init() {
	err := internal.ScaffoldConfig()

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	if len(os.Args) < 2 {
		fmt.Println("Expected a valid command, run -h | --help for usage")
		os.Exit(1)
	}
}

func parseCommand() {
	internal.RunCmd(os.Args[1])
}
