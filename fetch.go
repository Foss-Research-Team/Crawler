// Fetch prints the content found at a URL.

package main


import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"bytes"
)

func getPage(a string)  []byte {
	
	resp, err := http.Get(a)

	if err != nil {
		
		fmt.Fprintf(os.Stderr,"fetch: %v\n",err)
		
		return nil
	}

	b, err := ioutil.ReadAll(resp.Body)
	
	resp.Body.Close()

	if err != nil {
		
		fmt.Fprintf(os.Stderr,"fetch: reading HTML contents: %v\n",err)
		
		return nil
	}

	return b

}


func extract_urls(html_page []byte) [1024][] byte {
	
	search_str := []byte("https://")

	search_sub_domain := []byte("href=\"/")

	search_dif_domain := []byte("href=\"//")

	var urls [1024][]byte

	i := 0

	url_index := 0
	
	len_html_page := len(html_page)


	for ( (i < len_html_page) && (url_index < 1024) ) {

		i = bytes.Index(html_page[i:],search_str) + i

		if (i < 0) {
			
			break
		}

		for html_page[i] != 0x22 {
			
			urls[url_index] = append(urls[url_index],html_page[i])
			
			i++
		}

		url_index++

		i++
		
	}

	i = 0

	for ( (i < len_html_page) && (url_index < 1024) ) {

		i = bytes.Index(html_page[i:],search_dif_domain) + i

		if (i < 0) {
			
			break
		}

		urls[url_index] = append(urls[url_index],[]byte("https:")...)


		for html_page[i] != 0x22 {

			i++
		}

		i++

		for html_page[i] != 0x22 {

			urls[url_index] = append(urls[url_index],html_page[i])
			
			i++
		}

		url_index++

		i++
		
	}

	i = 0

	for ( (i < len_html_page) && (url_index < 1024) ) {

		i = bytes.Index(html_page[i:],search_sub_domain) + i

		if (i < 0) {
			
			break
		}

		for html_page[i] != 0x22 {
			
			i++
		}

		i++

		if ( bytes.Equal( html_page[i+1:i+2],[]byte("/") ) ) {
			

			continue
		}


		urls[url_index] = append(urls[url_index],[]byte(os.Args[1])...)

		for html_page[i] != 0x22 {

			
			urls[url_index] = append(urls[url_index],html_page[i])
			
			i++
		}

		url_index++

		i++
		
	}
	
	i = 0
	
	return urls
	
	
}

func crawler(url string) {
	
	var c []byte 
	
	c = getPage(url)

	if c == nil {
		return 
	}

	var url_list [1024][] byte
	
	url_list = extract_urls(c)

	i := 0

	for (i < len(url_list)) && (url_list[i] != nil) {
	
		fmt.Printf("%s\n",url_list[i])

		i++

	}
	
	i = 0
	
	for (i < len(url_list)) && (url_list[i] != nil) {
		
		crawler(string(url_list[i]))

		i++
	}

}


//Problem: We need to get the computer to detect if the URL leads to an HTML file

func main() {
	
	crawler(os.Args[1])
}

