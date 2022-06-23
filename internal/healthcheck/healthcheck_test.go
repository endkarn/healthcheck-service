package healthcheck

import (
	"github.com/go-playground/assert/v2"
	"gopkg.in/h2non/gock.v1"
	"testing"
	"time"
)

func Test_CheckWebsite_Get_3_Sites_All_Up_Website(t *testing.T) {
	service := Service{}
	defer gock.Off()
	gock.New("https://facebook.com").Get("").Reply(200)
	gock.New("https://google.com").Get("").Reply(200)
	gock.New("https://stackoverflow.com").Get("").Reply(200)
	sites := []WebsiteHealthCheck{{
		Order:          0,
		WebsiteURL:     "https://facebook.com",
		HTTPStatusCode: 0,
	}, {
		Order:          1,
		WebsiteURL:     "https://google.com",
		HTTPStatusCode: 0,
	}, {
		Order:          2,
		WebsiteURL:     "https://stackoverflow.com",
		HTTPStatusCode: 0,
	}}
	expected := []WebsiteHealthCheck{{
		Order:          0,
		WebsiteURL:     "https://facebook.com",
		HTTPStatusCode: 200,
	}, {
		Order:          1,
		WebsiteURL:     "https://google.com",
		HTTPStatusCode: 200,
	}, {
		Order:          2,
		WebsiteURL:     "https://stackoverflow.com",
		HTTPStatusCode: 200,
	}}
	actual := service.CheckWebsites(sites)
	assert.Equal(t, expected, actual)
}

func Test_CheckWebsite_Get_2_Sites_All_Up_Website(t *testing.T) {
	service := Service{}
	defer gock.Off()
	gock.New("https://youtube.com").Get("").Reply(200)
	gock.New("https://github.com").Get("").Reply(200)
	sites := []WebsiteHealthCheck{{
		Order:          0,
		WebsiteURL:     "https://youtube.com",
		HTTPStatusCode: 0,
	}, {
		Order:          1,
		WebsiteURL:     "https://github.com",
		HTTPStatusCode: 0,
	}}
	expected := []WebsiteHealthCheck{{
		Order:          0,
		WebsiteURL:     "https://youtube.com",
		HTTPStatusCode: 200,
	}, {
		Order:          1,
		WebsiteURL:     "https://github.com",
		HTTPStatusCode: 200,
	}}
	actual := service.CheckWebsites(sites)
	assert.Equal(t, expected, actual)
}

func Test_CheckWebsite_Get_3_Sites_2_Up_1_Down_Status_Website(t *testing.T) {
	service := Service{}
	defer gock.Off()
	gock.New("https://youtube.com").Get("").Reply(200)
	gock.New("https://github.com").Get("").Reply(200)
	gock.New("https://cloudflare.com").Get("").Reply(200).Delay(11 * time.Second)
	sites := []WebsiteHealthCheck{{
		Order:          0,
		WebsiteURL:     "https://youtube.com",
		HTTPStatusCode: 0,
	}, {
		Order:          1,
		WebsiteURL:     "https://github.com",
		HTTPStatusCode: 0,
	}, {
		Order:          2,
		WebsiteURL:     "https://www.cloudflare.com",
		HTTPStatusCode: 0,
	}}
	expected := []WebsiteHealthCheck{{
		Order:          0,
		WebsiteURL:     "https://youtube.com",
		HTTPStatusCode: 200,
	}, {
		Order:          1,
		WebsiteURL:     "https://github.com",
		HTTPStatusCode: 200,
	}, {
		Order:          2,
		WebsiteURL:     "https://www.cloudflare.com",
		HTTPStatusCode: 0,
	}}
	actual := service.CheckWebsites(sites)
	assert.Equal(t, expected, actual)
}

func Test_SendRequestForStatusCode_Get_404_From_URL(t *testing.T) {
	defer gock.Off()
	gock.New("https://google.com").Reply(404)
	actual := SendRequestForStatusCode("https://google.com")
	assert.Equal(t, 404, actual)
}

func Test_SendRequestForStatusCode_Get_200_From_URL(t *testing.T) {
	defer gock.Off()
	gock.New("https://youtube.com").Reply(200)
	actual := SendRequestForStatusCode("https://youtube.com")
	assert.Equal(t, 200, actual)
}

func Test_SendRequestForStatusCode_Get_Timeout_In_10sec_From_URL(t *testing.T) {
	defer gock.Off()
	gock.New("https://cloudflare.com").Reply(200).Delay(11 * time.Second)
	actual := SendRequestForStatusCode("https://cloudflare.com")
	assert.Equal(t, 0, actual)
}
