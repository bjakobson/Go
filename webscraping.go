package main

import (
	"fmt"
	"log"

	"github.com/anaskhan96/soup"
)

func main() {
	resp, err := soup.Get("https://www.amazon.com/Nintendo-Switch-Gray-Joy/dp/B01LTHP2ZK?pf_rd_p=616cfad6-05f3-425d-8659-8a0037d112c9&pd_rd_wg=rM9sc&pf_rd_r=W37PEBBGVGX96ZK8TRH6&ref_=pd_gw_cr_simh&pd_rd_w=FzOZk&pd_rd_r=2744e1a8-4e48-49ba-9388-ab35566bb290")
	if err != nil {
		log.Fatal(err)
	}
	doc := soup.HTMLParse(resp)
	links := doc.Find("div", "id", "titleSection").FindAll("span")
	for _, link := range links {
		fmt.Println(link.Text())
	}
}
 GO
 GO
 GO
 Go
