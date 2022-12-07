package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly"
)

// Information about the churches that will be stored
type Church struct {
	Name         string
	Address      string
	City         string
	State        string
	Zip          string
	Phone        string
	Website      string
	Email        string
	Denomination string
	Service      string
	Pastor       string
}

func main() {

	fName := "data.csv"
	file, err := os.Create(fName)
	if err != nil {
		log.Fatalf("Could not create the file. err: %q", err)
	}
	defer file.Close()

	// create a new CSV writer
	writer := csv.NewWriter(file)
	defer writer.Flush()

	c := colly.NewCollector(
		colly.AllowedDomains("www.usachurches.org"),
		colly.CacheDir("./cache"),
	)

	detailCollector := c.Clone()
	// churches := make([]Church, 0)

	c.OnHTML("strong", func(e *colly.HTMLElement) {

		link := e.ChildAttr("a[href]", "href")
		e.Request.Visit(link)

		c.OnRequest(func(r *colly.Request) {
			log.Println("visiting", r.URL.String())
		})

		// e.ForEach("a[href]", func(i int, elem *colly.HTMLElement) {
		// 	writer.Write([]string{
		// 		elem.Text,
		// 		elem.Attr("href")})
		// 	fmt.Println(elem.Attr("href"))
		// })
	})

	detailCollector.OnHTML("h1", func(e *colly.HTMLElement) {
		log.Println("Church found", e.Request.URL)
		title := e.Text
		if title == "" {
			log.Println("No title found", e.Request.URL)
		}
		fmt.Println(title)
	})
	// c.OnHTML("a[href]", func(e *colly.HTMLElement) {
	// 	e.Request.Visit(e.Attr("href"))
	// })
	// c.OnRequest(func(r *colly.Request) {
	// 	fmt.Println("Visiting", r.URL)
	// })

	c.Visit("https://www.usachurches.org/search/ga/atlanta/")
}
