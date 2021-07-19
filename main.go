package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
	"strings"
	"time"
)

func scrape_suzies() {
	c := colly.NewCollector(
		colly.MaxDepth(1),
	)
	var result []string
	c.SetRequestTimeout(55 * 1e9)
	c.OnHTML("div.food-menu", func(e *colly.HTMLElement) {
		today := time.Now().Format("2. 1.")
		e.DOM.Find("div.uk-card-body").Each(func(_ int, day_menu *goquery.Selection) {
			dow_date := strings.Split(day_menu.Find("h2").Text(), " ")
			dow := dow_date[0]
			date := dow_date[1] + " " + dow_date[2]
			if date == today {
				day_menu.Find("div.uk-grid-small").Each(func(_ int, s *goquery.Selection) { result = append(result, s.Text()) })
				fmt.Println("SUZIES daily menu @", dow, date)
				print_string_list(result)
			}
		})
	})
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("===============================")
		fmt.Println("Daily menu from: ", r.URL.String())
	})
	c.Visit("http://www.suzies.cz/poledni-menu")
}

func scrape_u_capa() {
	c := colly.NewCollector(
		colly.MaxDepth(1),
	)
	var result []string
	c.SetRequestTimeout(55 * time.Second)
	c.OnHTML("div.listek", func(e *colly.HTMLElement) {
		today := time.Now().Format("2. 1. 2006")
		e.ForEach("div.row", func(_ int, e1 *colly.HTMLElement) {
			dow := e1.ChildText("div.day")
			date := e1.ChildText("div.date")
			if date == today {
				e1.DOM.Find("div.row").Each(func(_ int, s *goquery.Selection) { result = append(result, s.Text()) })
				fmt.Println("U CAPA daily menu @", dow, date)
				print_string_list(result)
			}
		})
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("===============================")
		fmt.Println("Daily menu from: ", r.URL.String())
	})
	c.Visit("https://www.pivnice-ucapa.cz/denni-menu.php")
}

func scrape_veroni() {
	c := colly.NewCollector(
		colly.MaxDepth(1),
	)
	var result []string
	c.SetRequestTimeout(55 * 1e9)
	c.OnHTML(".obsah", func(e *colly.HTMLElement) {
		today := time.Now().Format("2.1.2006")
		e.ForEach("div.menicka", func(_ int, e1 *colly.HTMLElement) {
			dow_date := strings.Split(e1.ChildText("div.nadpis"), " ")
			if dow_date[1] == today {
				ul := e1.DOM.Find("ul").First()
				ul.Find("li").Each(func(_ int, s *goquery.Selection) {
					result = append(result, s.Find("div.polozka").First().Text()+" "+s.Find("div.cena").First().Text())
				})
				fmt.Println("VERONI daily menu @", today)
				print_string_list(result)
			}
		})
	})
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("===============================")
		fmt.Println("Daily menu from: ", r.URL.String())
	})
	c.Visit("https://www.menicka.cz/4921-veroni-coffee--chocolate.html")
}

func print_string_list(lst []string) {
	for i, s := range lst {
		fmt.Printf("'(%d)%s'\n", i, s)
	}
}

func main() {
	scrape_suzies()
	scrape_u_capa()
	scrape_veroni()
}
