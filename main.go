package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"flag"
	"github.com/PuerkitoBio/goquery"
)

func main() {
	var vpn = flag.String("vpn", "JJJ895AX", "Vehicle Plate Number")
	flag.Parse()
	resp, err := http.Get(fmt.Sprintf("%s%s", "http://www.lsmvaapvs.org/search.php?vpn=", strings.TrimSpace(*vpn)))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		fmt.Println("Registration Number is Invalid...!")
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	html := string(body)
	sanitized := strings.Replace(html, "<!--", "", 9)
	sanitized = strings.Replace(sanitized, "-->", "", 9)
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(sanitized))
	if err != nil {
		log.Fatal(err)
	}
	doc.Find(".form-group p").Each(func(i int, s *goquery.Selection) {
		fmt.Printf("%s ", s.Text())
		if i%2 != 0 {
			fmt.Println()
		}
	})
}
