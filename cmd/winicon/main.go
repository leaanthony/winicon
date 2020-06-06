package main

import (
	"os"

	"github.com/leaanthony/winicon/cmd/winicon/internal/app"
)

func main() {
	err := app.Run(os.Args, os.Stdout)
	if err != nil {
		println("ERROR:", err.Error())
		os.Exit(1)
	}
}
