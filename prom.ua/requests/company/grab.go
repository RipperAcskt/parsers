package company

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/RipperAcskt/prom.ua/proxy"
	"github.com/RipperAcskt/prom.ua/queue"
)

func New(topic string) string {
	return requestFirstPart + topic + requestSecondPart
}

func Request(body string, proxyAddresses *queue.Queue) (*http.Response, error) {
	client := &http.Client{
		Timeout: 15 * time.Second,
	}
	var resp *http.Response
	var err error
	for {
		ip := proxyAddresses.Get()
		// fmt.Printf("company\t%v\n", ip)
		client.Transport, err = proxy.CreateTransport(ip)
		if err != nil {
			return nil, fmt.Errorf("create transport failed: %v", err)
		}

		request, err := http.NewRequest("POST", "https://prom.ua/graphql", bytes.NewBuffer([]byte(body)))
		if err != nil {
			return nil, fmt.Errorf("new request failed: %v", err)
		}

		request.Header.Set("Content-type", "application/json")

		resp, _ = client.Do(request)
		if err == nil && resp != nil && resp.StatusCode != 429 {
			return resp, nil
		}
	}

}

func GetUrl(resp *http.Response) ([]string, error) {
	var all []ResponseJson

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read all failed: %v", err)
	}
	defer resp.Body.Close()
	// fmt.Println(string(body))
	err = json.Unmarshal(body, &all)
	if err != nil {
		return nil, fmt.Errorf("unmarshal failed: %v", err)
	}

	var id []int

	for i := range all[2].Data.Listing.Filters.CompanyFilter.Values {
		id = append(id, all[2].Data.Listing.Filters.CompanyFilter.Values[i].Value)
	}

	var idUrl []string

	for _, id := range id {
		idUrl = append(idUrl, url.QueryEscape(fmt.Sprint(id)))
	}

	return idUrl, nil
}
