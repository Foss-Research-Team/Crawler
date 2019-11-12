package main

import(
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/proxy"
)

func main() {
	c := colly.NewCollector()

	if p, err := proxy.RoundRobinProxySwitcher(
		"socks5://127.0.0.1:1337",
		"socks5://127.0.0.1:1338",
		"http://127.0.0.1:8080",

	); err == nil {
		c.SetProxyFunc(p)
	}

	var proxies [] *url.URL = []*url.URL{
		&url.URL{Host:"127.0.0.1:8080"},
		&url.URL{Host: "127.0.0.1:8081"},

	}

	func randomProxySwitcher(_ *http.Request) (*url.URL, error) {
		return proxies[random.Intn(len(proxies))], nil
	}

	// ...

	c.SetProxyFunc(randomProxySwitcher)
}


