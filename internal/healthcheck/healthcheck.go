package healthcheck

import (
	"net/http"
	"sync"
	"time"
)

type Service struct {
}

func (service Service) CheckWebsites(sites []WebsiteHealthCheck) []WebsiteHealthCheck {
	var wg sync.WaitGroup
	wg.Add(len(sites))
	for idx := range sites {
		pointer := &sites[idx]
		go CheckWebsite(pointer, &wg)
	}
	wg.Wait()
	return sites
}

func CheckWebsite(site *WebsiteHealthCheck, s *sync.WaitGroup) {
	site.HTTPStatusCode = SendRequestForStatusCode(site.WebsiteURL)
	s.Done()
}

func SendRequestForStatusCode(url string) int {
	client := http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := client.Get(string(url))
	if err != nil {
		return 0
	}
	return resp.StatusCode
}
