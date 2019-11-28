// Fetch prints the content found at a URL.

package main


import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"bytes"
	"crypto/sha256"
)

var sha_map = make(map[string][]byte)

var url_map = make(map[string] int)

var blacklist_domain_map = make(map[string] int)

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
	
	var search_str []byte = []byte("\"https://")

	//var url_str []byte = []byte("(https://") //use this for background images only

	search_sub_domain := []byte("href=\"/")

	search_dif_domain := []byte("href=\"//")
	
	var url []byte

	var urls [1024][]byte

	var html_of_url []byte

	var shasum [32]byte

	i := 0

	i_0 := 0

	url_index := 0
	
	len_html_page := len(html_page)

	fmt.Printf("len_html_page: %d\n",len_html_page)

	fmt.Printf("%s\n\n","Searching complete URLs")

	for ( (i < len_html_page) && (url_index < 1024) ) {

		i_0 = i

		i = bytes.Index(html_page[i_0:],search_str) + i_0

		if ( (i < 0) || (i < i_0) ) {

			url = []byte{}

			fmt.Printf("Last index for complete URLS found: %d\n\n",i)

			i += 2

			break
		}

		i++ //Have the character position at the beginning of the https URL

		for ( html_page[i] != 0x22 ) {

			url = append(url,html_page[i])
			
			i++
		}

		i++
		
		html_of_url = getPage( string(url) )

		if ( html_of_url == nil ) {


			fmt.Printf("Failed to get HTML of page (%s) at index %d\n\n",url,i)
			url = []byte{}

			i++

			continue
			
		}

		shasum = sha256.Sum256(html_of_url)
		
		if ( ( len(sha_map[string(shasum[0:])]) == 0 ) && ( url_map[string(url)] == 0 ) ) {

			fmt.Printf("%s\n",url)

			fmt.Printf("Sha256: %x\n\n",shasum)

			urls[url_index] = make([]byte,len(url))

			copy(urls[url_index],url)

			sha_map[string(shasum[0:])] = urls[url_index]
			
			fmt.Printf("url_map of %s beforehand: %d\n",urls[url_index],url_map[string(urls[url_index])])

			url_map[string(urls[url_index])] = 1

			url_index++
		
		} 

		url = []byte{}

		fmt.Println(i)

		i++
		
	}

	i = 0

	i_0 = 0

	url = []byte{}

	fmt.Printf("%s\n\n","Searching URLs with different domains")

	for ( (i < len_html_page) && (url_index < 1024) ) {
		
		i_0 = i

		i = bytes.Index(html_page[i_0:],search_dif_domain) + i_0

		if ( (i < 0) || ( i < i_0 ) ) {
			
			url = []byte{}

			fmt.Printf("Last index for different domains found: %d\n\n",i)
			i += 2

			break
		}

		url = append(url,[]byte("https:")...)


		for html_page[i] != 0x22 {

			i++
		}

		i++

		for ( html_page[i] != 0x22 ) {

			url = append(url,html_page[i])
			
			i++
		}

		i++
		
		html_of_url = getPage( string(url) )

		if ( html_of_url == nil ) {
			

			fmt.Printf("Failed to get HTML of page (%s) at index %d\n\n",url,i)
			url = []byte{}
			
			fmt.Println(i)

			i++

			continue
			
		}

		shasum = sha256.Sum256(html_of_url)
		
		if ( (len(sha_map[string(shasum[0:])]) == 0) && ( url_map[string(url)] == 0 ) ) {
			
			urls[url_index] = make([]byte,len(url))

			copy(urls[url_index],url)
			
			fmt.Printf("%s\n",urls[url_index])

			fmt.Printf("Sha256: %x\n\n",shasum)

			sha_map[string(shasum[0:])] = urls[url_index]

			fmt.Printf("url_map of %s beforehand: %d\n",urls[url_index],url_map[string(urls[url_index])])

			url_map[string(urls[url_index])] = 1

			url_index++
		
		} 

		url = []byte{}

		fmt.Println(i)

		i++
		
	}

	url = []byte{}
	
	i_0 = 0

	i = 0

	fmt.Println("Searching sub domain URLs\n\n")

	for ( (i < len_html_page) && (url_index < 1024) ) {

		i_0 = i	

		i = bytes.Index(html_page[i_0:],search_sub_domain) + i_0

		if ( (i < 0) || (i < i_0) ) {
			
			url = []byte{}

			fmt.Printf("Last index for sub domains found:%d\n\n",i)

			i += 2

			break
		}

		for html_page[i] != 0x22 {
			
			i++
		}

		i++

		if ( bytes.Equal( html_page[i+1:i+2],[]byte("/") ) ) {
			
			url = []byte{}

			fmt.Println(i)

			i++

			continue
		}


		url = append(url,[]byte(os.Args[1])...)

		for html_page[i] != 0x22 {

			url = append(url,html_page[i])
			
			i++
		}

		i++
		
		html_of_url = getPage( string(url) )

		if ( html_of_url == nil ) {
			
			url = []byte{}

			fmt.Println(i)

			i++

			continue
			
		}

		shasum = sha256.Sum256(html_of_url)
		
		if ( ( len(sha_map[string(shasum[0:])]) == 0 ) && ( url_map[string(url)] == 0 ) ) {
			
			urls[url_index] = make([]byte,len(url))

			copy(urls[url_index],url)
			
			fmt.Printf("%s\n",urls[url_index])

			fmt.Printf("Sha256: %x\n\n",shasum)

			sha_map[string(shasum[0:])] = urls[url_index]

			fmt.Printf("url_map of %s beforehand: %d\n",urls[url_index],url_map[string(urls[url_index])])

			url_map[string(urls[url_index])] = 1

			url_index++
		
		} 

		url = []byte{}
		
		fmt.Println(i)

		i++
		
	}
	
	i_0 = 0

	i = 0
	
	return urls
	
	
}

func crawler(url string) {
	
	shasum_base_url := sha256.Sum256([]byte(url))

	fmt.Printf("Checksum of HMTL page of base url is:\n%x\n\n",shasum_base_url)

	sha_map[string(shasum_base_url[0:])] = []byte(url)

	url_map[string(sha_map[string(shasum_base_url[0:])])] = 1

	fmt.Printf("%s\n",sha_map[string(shasum_base_url[0:])])
	
	var c []byte 
	
	c = getPage(url)

	if c == nil {
		return 
	}

	var url_list [1024][] byte
	
	url_list = extract_urls(c)

	i := 0
	
	for ( (i < len(url_list)) && (url_list[i] != nil) ) {
		
		crawler(string(url_list[i]))

		i++
	}

}

func blacklist(black_url []string) {

	var i int = 0

	for i < len(black_url) {

	url_map[black_url[i]] = 1

	blacklist_domain_map[black_url[i]] = 1

	i++

	}
	

}

func main() {
	
//	fmt.Printf("%s\n",getPage(os.Args[1]))

	crawler(os.Args[1])
}

