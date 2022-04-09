package main

import (
	"log"
	"os"

	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

func main() {

	cfg, err := LoadConfig("config.toml")
	if err != nil {
		log.Fatal(err)
	}

	logFile, err := os.OpenFile(cfg.Log, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer logFile.Close()
	log.SetOutput(logFile)

	store := SQLLiteStore{}
	if err = store.Connect(cfg.Db); err != nil {
		log.Fatal(err)
	}

	s := Server{}
	if err = s.Init(cfg.Srv, &store); err != nil {
		log.Fatal(err)
	}

	go s.Run()

	a := app.New()

	w := a.NewWindow("Server")

	btnSave := widget.NewButton("Выгрузка", func() {

	})

	btnExit := widget.NewButton("Выход", func() {
		a.Quit()
	})

	w.SetContent(
		widget.NewVBox(
			widget.NewHBox(layout.NewSpacer(), btnSave, btnExit),
		))

	a.Settings().SetTheme(theme.LightTheme())

	w.ShowAndRun()
}
