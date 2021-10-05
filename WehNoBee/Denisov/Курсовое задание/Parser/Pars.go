package main

import (
	"fmt"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/gocolly/colly"

	"github.com/maxence-charriere/go-app/v8/pkg/app"
)

type Word struct {
	En string `excel:"en"`
	Ru string `excel:"ru"`
}

var wordCollection = []Word{}
var cnt int

func main() {
	scrapPage("https://www.en365.ru/top1000.htm") //НУ здесь вписать сайт
	scrapPage("https://www.en365.ru/top1000a.htm")
	scrapPage("https://www.en365.ru/top1000b.htm")
	//fmt.Printf("%s \n", wordCollection)
	fmt.Printf("cnt %v \n", cnt)
	writeResultXls()

}

func scrapPage(url string) {
	c := colly.NewCollector()

	c.OnHTML("div > table tr", func(e *colly.HTMLElement) { //Сюда писать элемент (Только бог поможет понять хтмл код другого)
		enWord := e.DOM.Find("td:nth-child(2)").Text() //Это если надо именно 2 или другой элемент в в дереве
		ruWord := e.DOM.Find("td:nth-child(3)").Text() //Это если больше надо пройтись по строке ещё дальше
		if !strings.Contains(enWord, "Английское") {
			wordCollection = append(wordCollection, Word{enWord, ruWord})
			cnt++
		}
	})

	c.Visit(url)
}

func writeResultXls() {
	xlsx := excelize.NewFile()

	xlsx.NewSheet("Запись")

	for i, word := range wordCollection {
		xlsx.SetCellValue("Запись", fmt.Sprintf("A%v", i+1), word.En)
		xlsx.SetCellValue("Запись", fmt.Sprintf("B%v", i+1), word.Ru)
	}

	err := xlsx.SaveAs("./Slova.xlsx")
	if err != nil {
		fmt.Println(err)
	}
}

type hello struct {
	app.Compo

	name string
}

func (h *hello) Render() app.UI {
	return app.Div().Body(
		app.H1().Body(
			app.Text("Приложение "),
			app.If(h.name != "",
				app.Text(h.name),
			).Else(
				app.Text("Тест"),
			),
		),
		app.P().Body(
			app.Input().
				Type("text").
				Value(h.name).
				Placeholder("Введите адрес").
				AutoFocus(true).
				OnChange(h.ValueTo(&h.name)),
		),
	)
}
