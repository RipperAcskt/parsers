package main

import (
	"fmt"
	"log"
	"io/ioutil"
	"strings"
	"time"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"strconv"
)

var(	
	url = "https://icetrade.by/search/auctions?search_text=Б&zakup_type%5B1%5D=1&zakup_type%5B2%5D=1&establishment=0&t%5BTrade%5D=1&t%5BeTrade%5D=1&t%5BsocialOrder%5D=1&t%5BsingleSource%5D=1&t%5BAuction%5D=1&t%5BRequest%5D=1&t%5BcontractingTrades%5D=1&t%5Bnegotiations%5D=1&t%5BOther%5D=1&r%5B1%5D=1&r%5B2%5D=2&r%5B7%5D=7&r%5B3%5D=3&r%5B4%5D=4&r%5B6%5D=6&r%5B5%5D=5&sort=num%3Adesc&sbm=1&onPage=100"
	write = ""
	secondUrl = "https://ts.butb.by/ppt/ru/tree/demand/all/gpc_orders_json"
	thirdUrl = "https://goszakupki.by/tenders/posted?TendersSearch%5Btext%5D=Б&TendersSearch%5Bstatus%5D%5B%5D=Submission&page=1"
	catalog__token = ""
	csrftoken = ""
)

func initNoon() {
    t := time.Now()
    n := time.Date(t.Year(), t.Month(), t.Day(), 3, 0, 0, 0, t.Location())
    d := n.Sub(t)
    if d < 0 {
        n = n.Add(24 * time.Hour)
        d = n.Sub(t)
    }
    for {
        time.Sleep(d)
        d = 24 * time.Hour
        startAll()
    }
}

func MakeHTML(urll string, title string, WhatFound string){
	str := fmt.Sprintf(`<tr><th style="float: left;">%s<a href=%s>%s</a></th></tr>`, WhatFound, strings.Join(strings.Fields(urll), " "), strings.Join((strings.Fields(title)),  " "))
	write += str

}

func FirstSearch(ListOfFind []string){
	
	for i := 0; i < len(ListOfFind); i++{
		WhatFound := ListOfFind[i]
		urlForSearch := strings.Replace(url, "Б", WhatFound, -1)
		fmt.Println(urlForSearch)


		resp, err := http.Get(urlForSearch)

		if err != nil{
			log.Fatal(err)
		}
		defer resp.Body.Close()
		if resp.StatusCode != 200{
			log.Fatalf("fatal error: %d %s", resp.StatusCode, resp.Status)
		}

		//load html doc
		doc, err := goquery.NewDocumentFromReader(resp.Body)
		if err != nil{
			log.Fatal(err)
		}

		//find elements

		doc.Find(".lst.top").Each(func(i int, s *goquery.Selection){
			title := s.Find("a").Text()
			urll, _ := s.Find("a").Attr("href")

			if urll != "" && title != ""{

				MakeHTML(urll, title, WhatFound)
			}


		})

		time.Sleep(6*time.Second)
	}

}

func SecondSearch(ListOfFind []string){
	// for i := 0; i < len(ListOfFind); i++{
		client := http.DefaultClient


		req, err := http.NewRequest("GET", secondUrl, nil)
    	req.Header.Add("Accept", `text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8`)
    	req.Header.Add("User-Agent", `Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.89 Safari/537.36`)
    	resp, err := client.Do(req)
    	for{
			if err != nil{
				fmt.Println("error")
				// log.Fatal(err)
			}else{
				break
			}
			resp, err = client.Do(req)
		}
		html, err := ioutil.ReadAll(resp.Body)
		fmt.Printf("%s\n", html)
		defer resp.Body.Close()
		if resp.StatusCode != 200{
			log.Fatalf("fatal error: %d %s", resp.StatusCode, resp.Status)
		}

		//load html doc
		doc, err := goquery.NewDocumentFromReader(resp.Body)
		if err != nil{
			log.Fatal(err)
		}

		//find elements
		
		doc.Find("#catalog__token").Each(func(i int, s *goquery.Selection){
			catalog__token, _ = s.Attr("value")
			// urll, _ := s.Find("a").Attr("href")

			fmt.Println(catalog__token, "token")
		})
		doc.Find(".form-DS").Each(func(i int, s *goquery.Selection){
			csrftoken, _ = s.Find("input").Attr("name")
		})

	// }
}

func ThirdSearch(ListOfFind []string){
	
	for i := 0; i < len(ListOfFind); i++{
		pages := 0
		WhatFound := ListOfFind[i]
		urlForSearch := strings.Replace(thirdUrl, "Б", WhatFound, -1)
		// fmt.Println(urlForSearch)


		resp, err := http.Get(urlForSearch)

		for{
			if err != nil{
				log.Println("error")
				// log.Fatal(err)
			}else{
				break
			}
			resp, err = http.Get(urlForSearch)
		}
		defer resp.Body.Close()
		if resp.StatusCode != 200{
			log.Fatalf("fatal error: %d %s", resp.StatusCode, resp.Status)
		}

		//load html doc
		doc, err := goquery.NewDocumentFromReader(resp.Body)
		if err != nil{
			log.Fatal(err)
		}

		//find elements

		doc.Find(".summary").Each(func(i int, s *goquery.Selection){
			
			arr := strings.Fields(s.Text())
			st := arr[4]
			num, _ := strconv.Atoi(st[:len(st)-1])
			if num%20 == 0{
				pages = num/20
			}else{
				pages = num/20+1
			}
			// fmt.Println(num, "\n", pages)
		})

		for j := 1; j <= pages; j++{
			urlForSearch = strings.Replace(urlForSearch, "1", strconv.Itoa(j), -1)
			// fmt.Println(urlForSearch)


			resp, err = http.Get(urlForSearch)

			if err != nil{
				
				log.Fatal(err)
			}
			defer resp.Body.Close()
			if resp.StatusCode != 200{
				log.Fatalf("fatal error: %d %s", resp.StatusCode, resp.Status)
			}

			//load html doc
			doc, err := goquery.NewDocumentFromReader(resp.Body)
			if err != nil{
				log.Fatal(err)
			}

			//find elements

			doc.Find(".table.table-hover").Each(func(i int, s *goquery.Selection){
				tbody := s.Find("tbody")
				tbody.Find("a").Each(func(j int, sq *goquery.Selection){

					title := sq.Text()
					partOfUrl, _ := sq.Attr("href")
					urll := "https://goszakupki.by" + partOfUrl
					fmt.Println(title, "\n", urll)
					if urll != "" && title != ""{

						MakeHTML(urll, title, WhatFound)
					}
				})
				
				
			})
			time.Sleep(500)
		}
		
	}
}


func startAll(){
	data, err := ioutil.ReadFile("read.txt")
	s := strings.ReplaceAll(string(data), " ", "+")
	st := strings.Fields(string(s))

	if err != nil{
		log.Fatal(err)
		return
	}

	write += `<!DOCTYPE> <html><head><title>Parse</title></head><body>`
	write += `<h2>icetrade.by</h2>`
	write += `<table style="width=100%">`
	FirstSearch(st)
	write += `</table>`
	
	write += `<h2>goszakupki.by</h2>`
	write += `<table style="width=100%">`
	ThirdSearch(st)



	write += `</table></body></html>`
	ioutil.WriteFile("write.html", []byte(write), 0644)
}

func main() {
	// initNoon()
	startAll()
}