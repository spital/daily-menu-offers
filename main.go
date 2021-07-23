package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
	"regexp"
	"strings"
	"time"
)

func scrape_suzies() {
	c := colly.NewCollector(
		colly.MaxDepth(1),
	)
	var result []string
	var dow_date []string
	c.SetRequestTimeout(55 * time.Second)
	c.OnHTML("div.food-menu", func(e *colly.HTMLElement) {
		today := time.Now().Format("2. 1.")
		e.DOM.Find("div.uk-card-body").Each(func(_ int, day_menu *goquery.Selection) {
			dow_date = strings.Split(day_menu.Find("h2").Text(), " ")
			if dow_date[1]+" "+dow_date[2] == today {
				day_menu.Find("h3").Each(func(_ int, s *goquery.Selection) {
					result = append(result, trimEveryLine(s.Text()))
				})
				day_menu.Find("div.uk-width-expand").Each(func(_ int, s *goquery.Selection) {
					result = append(result, trimEveryLine(s.Text()))
				})
				day_menu.Find("div.price").Each(func(_ int, s *goquery.Selection) {
					result = append(result, trimEveryLine(s.Text()))
				})
			}
		})
		var res1 []string
		len3 := len(result)/3 + 1
		res1 = append(res1, result[0]+"::"+result[len3]) //soup
		for i := 2; i < len3; i++ {
			res1 = append(res1, result[i]+"::"+result[i+len3]+" "+result[i+len3*2-1])
		}
		//str1 := fmt.Sprintf("SUZIES daily menu @ %s %s", dow_date[0], today)
		var res2 []string
		res2 = append(res2, fmt.Sprintf("==============================="))
		res2 = append(res2, fmt.Sprintf("SUZIES daily menu @ %s %s", dow_date[0], today))
		result = append(res2, res1...)
		print_string_list(result)
	})
	c.Visit("http://www.suzies.cz/poledni-menu")
}

func scrape_u_capa() {
	c := colly.NewCollector(
		colly.MaxDepth(1),
	)
	var result []string
	var dow, date, dow_date string
	c.SetRequestTimeout(55 * time.Second)
	c.OnHTML("div.listek", func(e *colly.HTMLElement) {
		today := time.Now().Format("2. 1. 2006")
		e.DOM.Find("div.row").Each(func(_ int, daily_menu *goquery.Selection) {
			date = daily_menu.Find("div.date").Text()
			dow = daily_menu.Find("div.day").Text()
			if date == today {
			dow_date = dow + " " + date
				daily_menu.Find("div.row").Each(func(_ int, s *goquery.Selection) { result = append(result, trimEveryLine(s.Text())) })
			}
		})
		var res2 []string
		fmt.Println("dow date 2",dow,date)
		res2 = append(res2, fmt.Sprintf("==============================="))
		res2 = append(res2, fmt.Sprintf("U CAPA daily menu @ %s", dow_date))
		result = append(res2, result...)
		print_string_list(result)
	})
	c.Visit("https://www.pivnice-ucapa.cz/denni-menu.php")
}

func scrape_veroni() {
	c := colly.NewCollector(
		colly.MaxDepth(1),
	)
	var result []string
	var dow_date []string
	c.SetRequestTimeout(55 * time.Second)
	c.OnHTML(".obsah", func(e *colly.HTMLElement) {
		today := time.Now().Format("2.1.2006")
		e.ForEach("div.menicka", func(_ int, e1 *colly.HTMLElement) {
			dow_date = strings.Split(e1.ChildText("div.nadpis"), " ")
			if dow_date[1] == today {
				ul := e1.DOM.Find("ul").First()
				ul.Find("li").Each(func(_ int, s *goquery.Selection) {
					result = append(result, s.Find("div.polozka").First().Text()+" "+s.Find("div.cena").First().Text())
				})
			}
		})
		var res2 []string
		res2 = append(res2, fmt.Sprintf("==============================="))
		res2 = append(res2, fmt.Sprintf("VERONI daily menu @ %s %s", dow_date[0], dow_date[1]))
		result = append(res2, result...)
		print_string_list(result)
	})
	c.Visit("https://www.menicka.cz/4921-veroni-coffee--chocolate.html")
}

func print_string_list(lst []string) {
	for i, s := range lst {
		fmt.Printf("'(%d)%s'\n", i, s)
	}
}

func trimEveryLine(multiline string) string {
	var b strings.Builder
	for _, l := range strings.Split(multiline, "\n") {
		fmt.Fprintf(&b, " %v ", strings.TrimSpace(l))
	}
	space := regexp.MustCompile(`\s+`)
	ret := space.ReplaceAllString(b.String(), " ")
	return strings.TrimSpace(ret)
}

func main() {
	//scrape_suzies()
	scrape_u_capa()
	//scrape_veroni()
}
