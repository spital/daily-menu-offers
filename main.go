package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"strings"
	"time"
)

func scrape_u_capa() {
	c := colly.NewCollector(
		colly.MaxDepth(1),
	)
	c.SetRequestTimeout(55 * 1e9)
	c.OnHTML("div.listek", func(e *colly.HTMLElement) {
		today := time.Now().Format("2. 1. 2006")
		e.ForEach("div.row", func(_ int, e1 *colly.HTMLElement) {
			dow := e1.ChildText("div.day")
			date := e1.ChildText("div.date")
			if date == today {
				fmt.Println("today's menu", dow, date, e1)
			}
		})
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Daily menu from: ", r.URL.String())
	})
	c.Visit("https://www.pivnice-ucapa.cz/denni-menu.php")
}

func scrape_suzies() {
	c := colly.NewCollector(
		colly.MaxDepth(1),
	)
	c.SetRequestTimeout(55 * 1e9)
	c.OnHTML("#weekly-menu", func(e *colly.HTMLElement) {
		today := time.Now().Format("2.1.")
		e.ForEach("div.day", func(_ int, e1 *colly.HTMLElement) {
			dow_date := strings.Split(e1.ChildText("h4"), " ")
			if dow_date[1] == today {
				fmt.Println("today's menu", e1)
			}
		})
	})
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Daily menu from: ", r.URL.String())
	})
	c.Visit("http://www.suzies.cz/poledni-menu.html")
}

func scrape_veroni() {
	c := colly.NewCollector(
		colly.MaxDepth(1),
	)
	c.SetRequestTimeout(55 * 1e9)
	c.OnHTML(".obsah", func(e *colly.HTMLElement) {
		today := time.Now().Format("2.1.2006")
		e.ForEach("div.menicka", func(_ int, e1 *colly.HTMLElement) {
			dow_date := strings.Split(e1.ChildText("div.nadpis"), " ")
			if dow_date[1] == today {
				fmt.Println("today's menu", e1)
			}
		})
	})
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Daily menu from: ", r.URL.String())
	})
	c.Visit("https://www.menicka.cz/4921-veroni-coffee--chocolate.html")
}

func main() {
	scrape_suzies()
	scrape_u_capa()
	scrape_veroni()
}
