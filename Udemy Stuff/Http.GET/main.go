package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"golang.org/x/net/html"
)

func main() {
	GetHTTP()

}

// GetHTTP Webpage
func GetHTTP() {
	resp, _ := http.Get("https://bitdev.io")
	page, _ := ioutil.ReadAll(resp.Body)
	b := resp.Body
	defer b.Close() // close Body when the function returns

	z := html.NewTokenizer(b)
	t := z.Token()
	defer resp.Body.Close()
	fmt.Printf("%s", page)
	getHref(t)
}

func getHref(t html.Token) (ok bool, href string) {
	// Iterate over all of the Token's attributes until we find an "href"
	for _, a := range t.Attr {
		if a.Key == "href" {
			href = a.Val
			ok = true
		}
	}
	fmt.Printf("%s", href)
	return ok, href
}
