package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"golang.org/x/net/html"
)

func main() {
	// Get a web page
	resp, err := http.Get("https://en.wikipedia.org/wiki/Main_Page")
	CheckErr(err)

	// Find all the href attributes.
	//hrefs := GetHrefs(resp)

	body, err := ioutil.ReadAll(resp.Body)
	length := float64(len(body)) / 1e6
	fmt.Println(length)

	// for _, href := range *hrefs {
	// 	fmt.Println(string(href))
	// }

}

func GetHrefs(response *http.Response) *[]string {
	hrefs := new([]string)

	tokenizer := html.NewTokenizer(response.Body)

	for {
		tokens := tokenizer.Next()

		switch {
		case tokens == html.ErrorToken:
			//That's an error (or the end of the doc) for now, we're out of here.
			return hrefs

		case tokens == html.StartTagToken:
			token := tokenizer.Token()

			if token.Data == "a" {
				//it's an anchor!
				//what's the href?
				href := TagValue("href", &token.Attr)
				fmt.Println(href)
			}
		}
	}
}

// TagValue gets the value of a specific tag. Returns empty if tag not found
func TagValue(tag string, attrs *[]html.Attribute) string {
	for _, attr := range *attrs {
		if attr.Key == tag {
			return attr.Val
		}
	}
	return ""
}
