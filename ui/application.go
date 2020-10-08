package ui

import (
	"fyne.io/fyne"
	"fyne.io/fyne/widget"
)
// FirstWindow the first window view
type FirstWindow struct {
	 Application fyne.App
	 window fyne.Window

	 Option *widget.Select

	 DateLabel *widget.Label 
	 Date *widget.Entry

	 Pm2_5LevelLabel *widget.Label
	 Pm2_5Level *widget.Entry

	 HumidityLabel *widget.Label
	 Humidity *widget.Entry

	 TemperatureLabel *widget.Label
	 Temperature *widget.Entry

	 UsernameLabel *widget.Label
	 Username *widget.Entry
	 
	 PasswordLabel *widget.Label
	 Password *widget.Entry

	 SubmitButton *widget.Button

	 Log *widget.Label
}
// firstWindow declaration
var firstWindow FirstWindow

//Show hello
func Show(){
	firstWindow.window.ShowAndRun()
}
//InititializeUI :- This innitializes the ui with the provided parameters
func InititializeUI(a fyne.App, title string, submit func(*FirstWindow)){
	window := a.NewWindow(title);
	window.CenterOnScreen()
	window.FullScreen()
	Options := []string{"Post", "Put", "Delete","Post-Init"}

	firstWindow.Option = widget.NewSelect(Options,func(s string) {
		if s=="Post" {
			postSetup();
		}
		if s=="Put" {
			putSetup()
		}
		if s=="Delete" {
			deleteSetup()
		}
	})

	firstWindow.DateLabel = widget.NewLabel("Enter Date in format :-'dd-mm-yyyy' eg:12-05-2020 here : ")
	firstWindow.Date = widget.NewEntry()

	firstWindow.HumidityLabel = widget.NewLabel("Enter Humidity here : ")
	firstWindow.Humidity = widget.NewEntry()

	firstWindow.TemperatureLabel = widget.NewLabel("Enter Temperature here : ")
	firstWindow.Temperature = widget.NewEntry()

	firstWindow.Pm2_5LevelLabel = widget.NewLabel("Enter PM 2.5 level here : ")
	firstWindow.Pm2_5Level = widget.NewEntry()

	firstWindow.UsernameLabel = widget.NewLabel("Enter user name : ")
	firstWindow.Username = widget.NewEntry()

	firstWindow.PasswordLabel = widget.NewLabel("Enter Password : ")
	firstWindow.Password = widget.NewPasswordEntry()

	firstWindow.SubmitButton = widget.NewButton("Submit",func() {
		submit(&firstWindow)
	})

	firstWindow.Log = widget.NewLabel("")
	
	firstWindow.window = window
	firstWindow.Application = a

	firstWindow.window.SetContent(widget.NewVBox(firstWindow.Option,
		widget.NewHBox(firstWindow.DateLabel, firstWindow.Date),
		widget.NewHBox(firstWindow.TemperatureLabel,firstWindow.Temperature),
		widget.NewHBox(firstWindow.Pm2_5LevelLabel,firstWindow.Pm2_5Level),
		widget.NewHBox(firstWindow.HumidityLabel,firstWindow.Humidity),
		widget.NewHBox(firstWindow.UsernameLabel,firstWindow.Username),
		widget.NewHBox(firstWindow.PasswordLabel,firstWindow.Password),
		firstWindow.SubmitButton,
		firstWindow.Log))

}
func postSetup(){
	firstWindow.Option.Show()
	firstWindow.Date.Enable()
	firstWindow.Humidity.Enable()
	firstWindow.Temperature.Enable()
	firstWindow.Pm2_5Level.Enable()
	firstWindow.Username.Enable()
	firstWindow.Password.Enable()
	firstWindow.SubmitButton.Enable()

	firstWindow.Option.Show()
	firstWindow.Date.Show()
	firstWindow.Humidity.Show()
	firstWindow.Temperature.Show()
	firstWindow.Pm2_5Level.Show()
	firstWindow.Username.Show()
	firstWindow.Password.Show()
	firstWindow.SubmitButton.Show()
}
func putSetup(){
	postSetup()
}
func deleteSetup(){
	firstWindow.Option.Show()
	firstWindow.Date.Enable()
	firstWindow.Humidity.Hide()
	firstWindow.Temperature.Hide()
	firstWindow.Pm2_5Level.Hide()
	firstWindow.Username.Enable()
	firstWindow.Password.Enable()
	firstWindow.SubmitButton.Enable()

	firstWindow.Option.Show()
	firstWindow.Date.Show()
	firstWindow.Username.Show()
	firstWindow.Password.Show()
	firstWindow.SubmitButton.Show()
}

func postInitSetup(){
	firstWindow.Option.Show()
	firstWindow.Date.Hide()
	firstWindow.Humidity.Hide()
	firstWindow.Temperature.Hide()
	firstWindow.Pm2_5Level.Hide()
	firstWindow.Username.Enable()
	firstWindow.Password.Enable()
	firstWindow.SubmitButton.Enable()
}