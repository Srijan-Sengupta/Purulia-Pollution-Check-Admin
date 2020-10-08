package main

import (
	"github.com/Srijan-Sengupta/Purulia-Pollution-Check-Admin/ui"
	"github.com/Srijan-Sengupta/Purulia-Pollution-Check-Admin/util"

	"fyne.io/fyne/app"
)

func main() {
	a := app.New()
	ui.InititializeUI(a,"Purulia Pollution check",util.Submit)
	ui.Show()
}
