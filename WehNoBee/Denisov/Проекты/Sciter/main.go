package main

import (
	"fmt"

	"github.com/fatih/color"
	sciter "github.com/sciter-sdk/go-sciter"
	"github.com/sciter-sdk/go-sciter/window"
)

var root *sciter.Element
var rootSelectorErr error
var w *window.Window
var windowErr error

func init() {


	rect := sciter.NewRect(0, 100, 300, 350)
	w, windowErr = window.New(sciter.SW_TITLEBAR|
		sciter.SW_CONTROLS|
		sciter.SW_MAIN|
		sciter.SW_GLASSY,
		rect)

	if windowErr != nil {
		fmt.Println("Can not create new window")
		return
	}
	htloadErr := w.LoadFile("./simpleGuiApp.html")
	if htloadErr != nil {
		fmt.Println("Can not load html in the screen", htloadErr.Error())
		return
	}

	root, rootSelectorErr = w.GetRootElement()
	if rootSelectorErr != nil {
		fmt.Println("Can not select root element")
		return
	}

	w.SetTitle("Simple Calc")

}

func main() {

	addbutton, _ := root.SelectById("add")

	out1, errout1 := root.SelectById("output1")
	if errout1 != nil {
		color.Red("failed to bound output 1 ", errout1.Error())
	}
	addbutton.OnClick(func() {
		output := add()
		out1.SetText(fmt.Sprint(output))
	})

	w.Show()
	w.Run()

}


func add() int {

	
	in1, errin1 := root.SelectById("input1")
	if errin1 != nil {
		color.Red("failed to bound input 1 ", errin1.Error())
	}
	in2, errin2 := root.SelectById("input2")
	if errin2 != nil {
		color.Red("failed to bound input 2 ", errin2.Error())
	}

	in1val, errv1 := in1.GetValue()
	color.Green(in1val.String())

	if errv1 != nil {
		color.Red(errv1.Error())
	}
	in2val, errv2 := in2.GetValue()
	if errv2 != nil {
		color.Red(errv2.Error())
	}
	color.Green(in2val.String())

	return in1val.Int() + in2val.Int()
}
