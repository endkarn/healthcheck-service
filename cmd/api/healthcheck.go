package api

import (
	"encoding/csv"
	"fmt"
	"github.com/gin-gonic/gin"
	"healthcheck-service/internal/healthcheck"
	"net/http"
	"net/url"
	"strings"
)

type HealthCheckAPI struct {
	HealthCheckService healthcheck.Service
}

type HealthCheckResponse struct {
	TotalUp   int                              `json:"total_up"`
	TotalDown int                              `json:"total_down"`
	Sites     []healthcheck.WebsiteHealthCheck `json:"sites"`
}

func (api HealthCheckAPI) CheckSitesHandler(c *gin.Context) {
	file, _, err := c.Request.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("file err : %s", err.Error()))
		return
	}
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("file err : %s", err.Error()))
		return
	}
	sites, err := SitesMapper(records)
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("file err : %s", err.Error()))
		return
	}
	checkedResult := api.HealthCheckService.CheckWebsites(sites)
	upSites, downSites := countTotal(checkedResult)
	response := HealthCheckResponse{
		TotalUp:   upSites,
		TotalDown: downSites,
		Sites:     checkedResult,
	}
	c.JSON(http.StatusOK, response)
}

func SitesMapper(records [][]string) ([]healthcheck.WebsiteHealthCheck, error) {
	var sites []healthcheck.WebsiteHealthCheck
	if len(records) > 1 {
		for index, record := range records {
			site := healthcheck.WebsiteHealthCheck{
				Order: index,
			}
			siteUrl, err := url.ParseRequestURI(fillPrefixURL(record[0]))
			if err != nil {
				return nil, fmt.Errorf("error on parsing URL from CSV")
			}
			site.WebsiteURL = siteUrl.String()
			sites = append(sites, site)
		}
	} else {
		for index, record := range records[0] {
			site := healthcheck.WebsiteHealthCheck{
				Order: index,
			}
			siteUrl, err := url.ParseRequestURI(fillPrefixURL(record))
			if err != nil {
				return nil, fmt.Errorf("error on parsing URL from CSV")
			}
			site.WebsiteURL = siteUrl.String()
			sites = append(sites, site)
		}
	}
	return sites, nil
}

func countTotal(sites []healthcheck.WebsiteHealthCheck) (int, int) {
	upSites := 0
	downSites := 0
	for _, site := range sites {
		if site.HTTPStatusCode == 200 {
			upSites += 1
		} else {
			downSites += 1
		}
	}
	return upSites, downSites
}

func fillPrefixURL(url string) string {
	if len(url) == 0 {
		return ""
	}
	if strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://") {
		return url
	}
	return "https://" + url
}
