package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
	"regexp"
	"strings"
	"sync"
	"time"
)

func scrape_suzies(wg *sync.WaitGroup, strchan chan string, mutex *sync.Mutex) {
	c := colly.NewCollector(
		colly.MaxDepth(1),
	)
	var result []string
	var dow_date []string
	c.SetRequestTimeout(55 * time.Second)
	c.OnHTML("div.food-menu", func(e *colly.HTMLElement) {
		today := time.Now().Format("2.1.2006")
		e.DOM.Find("div.uk-card-body").Each(func(_ int, day_menu *goquery.Selection) {
			dow_date_act := strings.Split(day_menu.Find("h2").Text(), " ")
			if dow_date_act[1] == today {
				dow_date = dow_date_act
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
			res1 = append(res1, result[i]+"::"+result[i+len3]+" "+result[i+len3*2-1]) // third -1 `cos soup has no price
		}
		var res2 []string
		res2 = append(res2, fmt.Sprintf("======================================="))
		res2 = append(res2, fmt.Sprintf("SUZIES daily menu @ %s %s", dow_date[0], today))
		result = append(res2, res1...)
		mutex.Lock()
		for _, v := range result {
			strchan <- v
		}
		mutex.Unlock()
		wg.Done()
	})
	c.Visit("http://www.suzies.cz/poledni-menu")
}

func scrape_u_capa(wg *sync.WaitGroup, strchan chan string, mutex *sync.Mutex) {
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
		res2 = append(res2, fmt.Sprintf("======================================="))
		res2 = append(res2, fmt.Sprintf("U CAPA daily menu @ %s", dow_date))
		result = append(res2, result...)
		mutex.Lock()
		for _, v := range result {
			strchan <- v
		}
		mutex.Unlock()
		wg.Done()
	})
	c.Visit("https://www.pivnice-ucapa.cz/denni-menu.php")
}

func scrape_veroni(wg *sync.WaitGroup, strchan chan string, mutex *sync.Mutex) {
	c := colly.NewCollector(
		colly.MaxDepth(1),
	)
	var result []string
	var dow_date []string
	c.SetRequestTimeout(55 * time.Second)
	c.OnHTML(".obsah", func(e *colly.HTMLElement) {
		today := time.Now().Format("2.1.2006")
		e.ForEach("div.menicka", func(_ int, e1 *colly.HTMLElement) {
			dow_date_act := strings.Split(e1.ChildText("div.nadpis"), " ")
			if dow_date_act[1] == today {
				dow_date = dow_date_act
				ul := e1.DOM.Find("ul").First()
				ul.Find("li").Each(func(_ int, s *goquery.Selection) {
					result = append(result, s.Find("div.polozka").First().Text()+" "+s.Find("div.cena").First().Text())
				})
			}
		})
		var res2 []string
		res2 = append(res2, fmt.Sprintf("======================================="))
		res2 = append(res2, fmt.Sprintf("VERONI daily menu @ %s %s", dow_date[0], dow_date[1]))
		result = append(res2, result...)
		mutex.Lock()
		for _, v := range result {
			strchan <- v
		}
		mutex.Unlock()
		wg.Done()
	})
	c.Visit("https://www.menicka.cz/4921-veroni-coffee--chocolate.html")
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
	var (
		mutex sync.Mutex
		wg    sync.WaitGroup
	)
	strchan := make(chan string, 40)

	wg.Add(3)
	go scrape_suzies(&wg, strchan, &mutex)
	go scrape_u_capa(&wg, strchan, &mutex)
	go scrape_veroni(&wg, strchan, &mutex)
	wg.Wait()

	close(strchan)

	for s := range strchan {
		fmt.Println(s)
	}

}
