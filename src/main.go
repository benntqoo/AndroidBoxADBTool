package src

import (
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func main() {
	application := app.New()

	w := application.NewWindow("Android Box ADB 工具")
	w.SetContent(widget.NewVBox(
		widget.NewLabel("Hello Fyne!"),
		widget.NewButton("Quit", func() {
			application.Quit()
		}),
	))
	w.ShowAndRun()
}
