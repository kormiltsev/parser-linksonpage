package finder

import (
	"fmt"
	"os"

	"github.com/kormiltsev/parser-linksonpage/internal/urls"
)

func Urlfinder(url string) []string {
	links, err := urls.FindLinks(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "parse: %v\n", err)
	}
	return links
}
