package html_link_parser

import (
	"fmt"
	"os"
	"strings"
	"golang.org/x/net/html"
)

var htmlFile, err = os.Open("index.html")
var z = html.NewTokenizer(htmlFile)

type Link struct {
	Href string
	Text string
}

var links []Link

func GetLinks() []Link {
	for {
		token := z.Next()
		if token == html.ErrorToken {
			break
		}
		tn, _ := z.TagName()
		tagName := string(tn)
		if token == html.StartTagToken {
			if tagName == "a" {
				for {
					// println(tagName)
					key, value, more := z.TagAttr()
					stringKey := string(key)
					stringValue := string(value)
					if stringKey == "href" {
						var link = Link{}
						link.Href = stringValue
						link.Text = getText(z)
						links = append(links, link)
					}

					if !more {
						break
					}
				}
			}
		}
	}
	for _, link := range links {
		fmt.Println("href", link.Href)
		fmt.Println("value", link.Text)
	}
	return links
}

func getText(z *html.Tokenizer) string {
	var depth int = 1
	var text string
outerloop:
	for {
		tokenType := z.Next()
		tn, _ := z.TagName()
		tagName := string(tn)
		switch tokenType {
		case html.StartTagToken:
			depth++
			text += " " + tagName + " "
		case html.EndTagToken:
			depth--
			if depth == 0 {
				break outerloop
			}
			text += " " + tagName + " "
		case html.TextToken:
			text += strings.TrimSpace(string(z.Text()))
		default:
			break outerloop
		}
	}
	return text
}
