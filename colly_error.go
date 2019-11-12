package main

import(
	"fmt"
	"github.com/gocolly/colly"
)

func main(){
	// Create a collector

	c := colly.NewCollector()

	// Set HTML callback
	//Will not be called if an error occurs
	
	c.OnHTML("*",func(e *colly.HTMLElement) {
		fmt.Println(e)
	})

	// Set error handler

	c.OnError(func(r *colly.Response,err error) {
		fmt.Println("Request URL:",r.Request.URL,"failed with response:",r,"\nError:",err)
	})
	
	// Set scraping

	c.Visit("https://duckduckgo.com")
	

	}

	
