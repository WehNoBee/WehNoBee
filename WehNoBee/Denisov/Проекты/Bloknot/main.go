package main

import (
	"os"

	"github.com/mattn/go-gtk/gtk"
)

var last int

func main() {
	gtk.Init(&os.Args)

	window := gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
	window.SetTitle("Notes")
	window.SetIconName("gtk-about")
	window.Connect("destroy", func() {
		gtk.MainQuit()
	})

	notebook := gtk.NewNotebook()
	notebook.AppendPage(gtk.NewVBox(false, 1),
		gtk.NewImageFromStock(gtk.STOCK_ADD, gtk.ICON_SIZE_MENU))

	notebook.Connect("button-release-event", func() bool {
		n := notebook.GetCurrentPage()
		if n == last {
			addPage(notebook)
		}
		return false
	})

	window.Add(notebook)
	window.SetSizeRequest(500, 300)
	window.ShowAll()

	gtk.Main()
}

func addPage(notebook *gtk.Notebook) {
	dialog := gtk.NewDialog()
	dialog.SetTitle("Title?")
	dVbox := dialog.GetVBox()

	input := gtk.NewEntry()
	input.SetEditable(true)
	vbox := gtk.NewVBox(false, 1)
	input.Connect("activate", func() {
		s := input.GetText()
		if s != "" {
			last = notebook.InsertPage(vbox, gtk.NewLabel(s), last) + 1
			notebook.ShowAll()
		}
		notebook.PrevPage()
		dialog.Destroy()
	})

	dVbox.Add(input)
	button := gtk.NewButtonWithLabel("OK")
	button.Connect("clicked", func() {
		input.Emit("activate")
	})
	dVbox.Add(button)
	dialog.SetModal(true)
	dialog.ShowAll()

	hbox := gtk.NewHBox(false, 1)
	swin := gtk.NewScrolledWindow(nil, nil)
	swin.SetPolicy(gtk.POLICY_AUTOMATIC, gtk.POLICY_AUTOMATIC)
	textview := gtk.NewTextView()
	swin.Add(textview)

	butBold := gtk.NewToolButtonFromStock(gtk.STOCK_BOLD)
	butIta := gtk.NewToolButtonFromStock(gtk.STOCK_ITALIC)

	butFont := gtk.NewFontButton()
	butFont.Connect("font-set", func() {
		textview.ModifyFontEasy(butFont.GetFontName())
	})

	butClose := gtk.NewToolButtonFromStock(gtk.STOCK_DELETE)
	butClose.Connect("clicked", func() {
		n := notebook.GetCurrentPage()
		notebook.RemovePage(notebook, n)
		last--
		notebook.PrevPage()
	})

	hbox.PackStart(butBold, false, false, 0)
	hbox.PackStart(butIta, false, false, 0)
	hbox.PackStart(butFont, false, false, 0)
	hbox.PackEnd(butClose, false, false, 0)
	vbox.PackStart(hbox, false, false, 0)
	vbox.Add(swin)
	notebook.ShowAll()
}
