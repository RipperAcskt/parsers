package proxy

import (
	"log"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/RipperAcskt/prom.ua/queue"
)

func Scrap_proxy() (*queue.Queue, error) {
	resp, err := http.Get("https://raw.githubusercontent.com/TheSpeedX/SOCKS-List/master/http.txt")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	ipPort := strings.Split(doc.Text(), "\n")

	var wg sync.WaitGroup
	var cheked []string
	for i := 0; i < len(ipPort); i++ {

		client := &http.Client{
			Timeout: 15 * time.Second,
		}
		wg.Add(1)
		go func(ip string) {
			defer wg.Done()
			client.Transport, err = CreateTransport(ip)

			if err != nil {
				log.Fatal(err)
			}
			request, err := http.NewRequest("GET", "https://prom.ua", nil)
			if err != nil {
				log.Fatalf("new request failed: %v", err)
			}
			_, err = client.Do(request)
			if err == nil {
				cheked = append(cheked, ip)
			}
		}(ipPort[i])

	}
	wg.Wait()
	result := queue.New(cheked)
	return result, nil
}

func CreateTransport(ip string) (*http.Transport, error) {
	var proxy = "http://" + ip

	proxyURL, err := url.Parse(proxy)

	if err != nil {
		return nil, err
	}

	transport := &http.Transport{
		Proxy: http.ProxyURL(proxyURL),
	}
	return transport, nil

}
