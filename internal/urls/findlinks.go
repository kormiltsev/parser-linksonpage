package urls

import (
	"fmt"

	"golang.org/x/net/html"
)

// 1/2 looking by block type
func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}

	return links
}

// 2/2 looking by block names
func visit2(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "pre" {
		for _, a := range n.Attr {
			if a.Val == "chordsBlock" {
				fmt.Println("itemprop = ", a)
				links = append(links, a.Val)
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit2(links, c)
	}
	return links
}
