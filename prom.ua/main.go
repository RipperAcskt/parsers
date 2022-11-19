package main

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/RipperAcskt/prom.ua/config"
	"github.com/RipperAcskt/prom.ua/proxy"
	"github.com/RipperAcskt/prom.ua/requests/company"
	phone "github.com/RipperAcskt/prom.ua/requests/phoneNumbers"
)

func main() {
	n := time.Now()
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("config new failed: %v", err)
	}

	fmt.Println("Getting proxy...")
	proxyAddresses, err := proxy.Scrap_proxy()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Getting proxy successfully!")
	fmt.Println()

	var wg sync.WaitGroup
	for _, topic := range cfg {
		wg.Add(1)
		go func(topic string) {
			defer func(topic string) {
				fmt.Printf("%v done\n", topic)
				fmt.Println()
				wg.Done()
			}(topic)

			fmt.Printf("%v start\n", topic)
			fmt.Println()
			if topic == "" {
				return
			}
			body := company.New(topic)

			resp, err := company.Request(body, proxyAddresses)
			if err != nil {
				log.Fatalf("info request failed: %v", err)
			}

			componyUrl, err := company.GetUrl(resp)
			if err != nil {
				log.Fatalf("get url failed: %v", err)
			}

			phoneNumbers, err := phone.GetNumbers(componyUrl, proxyAddresses, topic)
			if err != nil {
				log.Fatalf("get numbers failed: %v", err)
			}

			err = phone.CreateFile(phoneNumbers, topic)
			if err != nil {
				log.Fatalf("create file failed: %v", err)
			}
		}(topic)
	}
	wg.Wait()
	fmt.Printf("Time spent: %v\n", time.Since(n))
}
