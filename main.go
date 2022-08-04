//package main
//
//import (
//	"fmt"
//	"github.com/gocolly/colly"
//	"strings"
//)
//
//func main() {
//	c := colly.NewCollector()
//	//a := []string{}
//	var a string
//
//	c.OnHTML(".classnav", func(e *colly.HTMLElement) {
//		if (e.Attr()
//			a = e.Text
//		fmt.Println(e.Text)
//		//fmt.Println(e.ChildText("возраст"))
//
//	})
//
//	//}
//
//	c.Visit("https://tesera.ru/game/maple-valley/")
//	b := strings.Split(a, "        ")
//	fmt.Println(strings.Split(a, ""), len(b))
//}

package main

import (
	"fmt"
	googlesearch "github.com/rocketlaunchr/google-search"
)

//
//import (
//	"github.com/PuerkitoBio/goquery"
//	"net/http"
//)
//import "log"
//import "fmt"
//
//func main() {
//	webUrl := "https://tesera.ru/game/maple-valley/"
//	response, err := http.Get(webUrl)
//
//	if err != nil {
//		log.Fatalln(err)
//	} else if response.StatusCode == 200 {
//		fmt.Println("We can scrape this")
//	} else {
//		log.Fatalln("Do not scrape this")
//	}
//
//	document, err := goquery.NewDocumentFromReader(response.Body)
//
//	if err != nil {
//		log.Fatalln(err)
//	}
//
//	//var a string
//	document.Find("li").Each(func(index int, selector *goquery.Selection) {
//		if selector.Find("img").AttrOr("title", "") == "возраст" {
//			fmt.Println(selector.Text())
//		}
//
//		if selector.Find("img").AttrOr("title", "") == "число игроков" {
//			fmt.Println(selector.Text())
//		}
//
//		if selector.Find("img").AttrOr("title", "") == "рекомендуемое число игроков" {
//			fmt.Println(selector.Text())
//		}
//
//		if selector.Find("img").AttrOr("title", "") == "время партии" {
//			fmt.Println(selector.Text())
//		}
//
//		//b := strings.Split(a, "")
//		//
//		//fmt.Println(b)
//	})
//}

func main() {
	result, _ := googlesearch.Search(nil, "https://www.google.com/search?q=Манчкин+Тесера")
	url := result[0].URL

	fmt.Println(url)
}
