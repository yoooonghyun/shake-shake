package main

import (
	"fmt"
	"shake-shake/src/app"
)

func main() {
	a, err := app.CreateApp()

	if err != nil {
		fmt.Print(err)
		return
	}

	a.Run()
}
