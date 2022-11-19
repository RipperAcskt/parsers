package phonenumbers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/RipperAcskt/prom.ua/proxy"
	"github.com/RipperAcskt/prom.ua/queue"
	"github.com/xuri/excelize/v2"
)

type phone struct {
	name        string
	description []string
	number      []string
}

func request(id, ip string, client *http.Client) (*http.Response, error) {
	var err error
	client.Transport, err = proxy.CreateTransport(ip)
	if err != nil {
		return nil, fmt.Errorf("create transport failed: %v", err)
	}

	body := requestFirstPart + id + requestSecondPart

	request, err := http.NewRequest("POST", "https://prom.ua/graphql", bytes.NewBuffer([]byte(body)))
	if err != nil {
		return nil, fmt.Errorf("new request failed: %v", err)
	}

	request.Header.Set("Content-type", "application/json")

	resp, _ := client.Do(request)
	return resp, nil
}

func GetNumbers(componyUrl []string, proxy *queue.Queue, topic string) ([]phone, error) {
	client := &http.Client{
		Timeout: 15 * time.Second,
	}

	var phoneNumbers []phone
	var wg sync.WaitGroup
	for _, id := range componyUrl {
		wg.Add(1)
		go func(id string) {
			defer wg.Done()
			var resp *http.Response
			var err error
			var ip string
			for {
				ip = proxy.Get()
				// fmt.Printf("phone\t%v\n", ip)
				resp, err = request(id, ip, client)
				if err == nil && resp != nil && resp.StatusCode != 429 {
					break
				}
			}
			var info []CompanyInfo

			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Fatalf("read all failed: %v", err)
			}

			defer resp.Body.Close()

			// fmt.Println(ip, string(body))
			err = json.Unmarshal(body, &info)
			if err != nil {
				log.Fatalf("unmarshal failed: %v", err)
			}

			var phoneNumber phone

			phoneNumber.name = info[0].Data.Company.Name

			for _, phone := range info[0].Data.Company.Phones {
				phoneNumber.description = append(phoneNumber.description, phone.Description)
				phoneNumber.number = append(phoneNumber.number, phone.Number)

			}
			phoneNumbers = append(phoneNumbers, phoneNumber)

		}(id)
	}
	wg.Wait()
	return phoneNumbers, nil
}

func CreateFile(p []phone, name string) error {
	f := excelize.NewFile()
	rowIndex := 1

	for _, phone := range p {
		err := f.SetCellValue("Sheet1", fmt.Sprintf("A%d", rowIndex), phone.name)
		if err != nil {
			return fmt.Errorf("set cell value name: %v", err)
		}

		for i := 0; i < len(phone.description); i++ {
			err = f.SetCellValue("Sheet1", fmt.Sprintf("B%d", rowIndex), phone.description[i])
			if err != nil {
				return fmt.Errorf("set cell value description: %v", err)
			}

			err = f.SetCellValue("Sheet1", fmt.Sprintf("C%d", rowIndex), phone.number[i])
			if err != nil {
				return fmt.Errorf("set cell value number: %v", err)
			}

			rowIndex++
		}
	}

	if err := f.SaveAs(name + ".xlsx"); err != nil {
		return fmt.Errorf("save as: %v", err)
	}
	return nil
}
