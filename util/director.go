package util

import (
	"github.com/Srijan-Sengupta/Purulia-Pollution-Check-Admin/ui"
)
//Submit function for ui.
func Submit(fWindow *ui.FirstWindow){
	option := fWindow.Option.Selected
	println(option)
	if option=="Post" {
		post(fWindow)
	}else if option == "Put" {
		put(fWindow)
	}else if option == "Delete" {
		delete(fWindow)
	}else if option == "Post-Init" {
		postInit(fWindow)
	}
}
