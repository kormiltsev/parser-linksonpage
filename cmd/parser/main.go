package main

import (
	"fmt"

	"github.com/kormiltsev/parser-linksonpage/internal/finder"
	"github.com/kormiltsev/parser-linksonpage/internal/urls"
)

func main() {
	// для ввода через панель https://ya.ru
	//for _, url := range os.Args[1:] {

	// слайс линков из заданного списка
	links := urls.Catalog()     //-----
	for _, url := range links { //-----

		fmt.Println("links on page = ", url)

		// поиск линков на странице
		var answerlinks []string
		answerlinks = finder.Urlfinder(url)
		for _, link := range answerlinks {

			fmt.Println(link)
		}
	}
}
