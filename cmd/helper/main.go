package main

import (
	"flag"
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/unit"

	"github.com/trevinteacutter/mwo-helper/internal/cmd/helper"
)

func main() {
	flag.Parse()

	go func() {
		window := app.NewWindow(
			app.Title("MWO Helper"),
			app.Size(unit.Dp(1280), unit.Dp(720)),
		)

		if err := helper.Loop(window); err != nil {
			log.Panic(err)
		}

		os.Exit(0)
	}()

	app.Main()
}
