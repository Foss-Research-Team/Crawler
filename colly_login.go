package main

import(
	"log"
	"github.com/gocolly/colly"
)

func main() {
	// create a new collector

	c := colly.NewCollector()

	//authenticate

	err := c.Post("https://stallman.org",map[string]string{"username":"admin","password":"admin"})

	if err != nil {
		log.Fatal(err)
	}

	// attach callbacks after login

	c.OnResponse(func(r *colly.Response) {

		log.Println("response received",r.StatusCode)
	})

	// start scraping

	c.Visit("https://lastpass.com")
}
