package main

import(
	"github.com/gocolly/colly"
)

func main(){
	c := colly.NewCollector(
		colly.UserAgent("myUserAgent"),
		colly.AllowedDomains("foo.com","bar.com"),
	)

	//Custom User Agent and allowed domains are cloned to c2

	c2 := c.Clone()

	c.OnResponse(func(r *colly.Response) {
		r.Ctx.Put(r.Headers.Get("Custom-Header"))
		c2.Request("GET", "https://foo.com/", nil, r.Ctx, nil)
	})
}
