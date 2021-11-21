package main

import (
	"io/ioutil"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)

var count int = 1 // globally declared
func main() {
	a := app.New()
	w := a.NewWindow("Text Editor")
	w.Resize(fyne.NewSize(600, 600))

	content := container.NewVBox(
		container.NewHBox(
			widget.NewLabel("Text Editor"),
			widget.NewLabel("Just For Fun"),
		),
	)
	content.Add(widget.NewButton("Add New File", func() {
		content.Add(widget.NewLabel("New File" + strconv.Itoa(count)))
		count++
	}))

	// input := widget.NewEntry() // only one line entry
	input := widget.NewMultiLineEntry()
	input.SetPlaceHolder("Enter text..")

	saveBtn := widget.NewButton("Save File", func() {
		saveFileDialog := dialog.NewFileSave(
			func(uc fyne.URIWriteCloser, _ error) {
				textData := []byte(input.Text) // byte array is created that stores the input text
				uc.Write(textData)
			}, w)
		saveFileDialog.SetFileName("New File" + strconv.Itoa((count)) + ".txt")
		saveFileDialog.Show()
	})

	openBtn := widget.NewButton("Open File", func() {
		openFileDialog := dialog.NewFileOpen(
			func(r fyne.URIReadCloser, _ error) {
				readData, _ := ioutil.ReadAll(r)

				output := fyne.NewStaticResource("New File", readData)

				viewData := widget.NewMultiLineEntry()
				viewData.SetText(string(output.StaticContent))
				////////////////////////////////
				// homework
				// saveBtn := widget.NewButton("Save File", func() {
				// 	saveFileDialog := dialog.NewFileSave(
				// 		func(uc fyne.URIWriteCloser, _ error) {
				// 			textData := []byte(input.Text) // byte array is created that stores the input text
				// 			uc.Write(textData)
				// 		}, w)
				// 	saveFileDialog.SetFileName("New File" + strconv.Itoa((count)) + ".txt")
				// 	saveFileDialog.Show()
				// })
				/////////////////////////////////////
				w := fyne.CurrentApp().NewWindow(
					string(output.StaticName))
				w.SetContent(container.NewScroll(viewData))
				w.Resize(fyne.NewSize(400, 400))
				w.Show()
			}, w)
		openFileDialog.SetFilter(storage.NewExtensionFileFilter([]string{".txt"}))
		openFileDialog.Show()
	})
	w.SetContent(
		container.NewVBox(
			content,
			input,
			container.NewHBox(
				saveBtn,
				openBtn,
			),
		),
	)
	// api>dialoge>file dialog documentataion
	w.ShowAndRun()
}
