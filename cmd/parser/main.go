package main

import (
	"fmt"

	"github.com/kormiltsev/parser-linksonpage/internal/finder"
	"github.com/kormiltsev/parser-linksonpage/internal/urls"
)

func main() {
	// for input via terminal
	//for _, url := range os.Args[1:] {

	// for using list of urls
	links := urls.Catalog()     //-----
	for _, url := range links { //-----

		fmt.Println("links on page = ", url)

		// looking for links on page
		var answerlinks []string
		answerlinks = finder.Urlfinder(url)
		for _, link := range answerlinks {

			fmt.Println(link)
		}
	}
}
