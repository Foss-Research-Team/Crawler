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
}

	
