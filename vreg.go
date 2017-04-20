package vreg

import (
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"net/http"
	"strings"
)

func Query(vpn string) error {
	resp, err := http.Get(fmt.Sprintf("%s%s", "http://www.lsmvaapvs.org/search.php?vpn=", strings.TrimSpace(vpn)))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return errors.New("Registration Number is Invalid...!")
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	html := string(body)
	//sanitized := strings.Replace(html, "<!--", "", 9)
	//sanitized = strings.Replace(sanitized, "-->", "", 9)
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return err
	}
	fmt.Println()
	doc.Find(".form-group p").Each(func(i int, s *goquery.Selection) {
		fmt.Printf("%s ", s.Text())
		if i%2 != 0 {
			fmt.Println()
		}
	})
	return nil
}
