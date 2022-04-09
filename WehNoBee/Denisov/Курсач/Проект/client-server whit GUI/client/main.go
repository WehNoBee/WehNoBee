package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"

	"fyne.io/fyne"
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

	log.Println("Server addr: ", cfg.Srv)

	a := app.New()

	w := a.NewWindow("Client")

	planEvent := widget.NewEntry()
	executor := widget.NewEntry()
	count := widget.NewEntry()
	stateLabel := widget.NewLabel("")
	form := widget.NewForm(
		widget.NewFormItem("Добавить", planEvent),
		widget.NewFormItem("Назад", executor),
		widget.NewFormItem("Количество обуч.", count),
		widget.NewFormItem("Excel", stateLabel),
	)

	btnSave := widget.NewButton("OK", func() {
		dataSend := fmt.Sprintf("%s %s %s", planEvent.Text, executor.Text, count.Text)
		if err := sendData(cfg.Srv, dataSend); err != nil {
			log.Println(err)
			stateLabel.SetText(err.Error())
		} else {
			stateLabel.SetText("данные сохранены")
		}
	})

	btnClear := widget.NewButton("Очистить", func() {
		planEvent.SetText("")
		executor.SetText("")
		count.SetText("")
		stateLabel.SetText("")
	})

	btnExit := widget.NewButton("Выход", func() {
		a.Quit()
	})

	w.SetContent(
		widget.NewVBox(
			form,
			layout.NewSpacer(),
			widget.NewHBox(layout.NewSpacer(), btnClear, btnSave, btnExit),
		))

	w.Resize(fyne.NewSize(480, 200))
	a.Settings().SetTheme(theme.LightTheme())

	w.ShowAndRun()
}

func sendData(addr, data string) error {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return fmt.Errorf("Error connect to server: %v\n", err)
	}

	if _, err := conn.Write([]byte(data)); err != nil {
		return fmt.Errorf("Error send data to server: %v\n", err)
	}

	log.Printf("Send data: %v\n", data)

	resp := make([]byte, 512)
	respLen, err := conn.Read(resp)
	if err != nil && err != io.EOF {
		return fmt.Errorf("Response error: %v\n", err)
	}

	log.Println("Packet processing: ", string(resp[:respLen]))

	return conn.Close()
}
