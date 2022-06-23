package api

import (
	"github.com/go-playground/assert/v2"
	"healthcheck-service/internal/healthcheck"
	"testing"
)

func Test_SitesMapper_Input_1st_CSV_Format_Should_Get_Correct_Array_Struct(t *testing.T) {
	expected := []healthcheck.WebsiteHealthCheck{{
		Order:          0,
		WebsiteURL:     "https://facebook.com",
		HTTPStatusCode: 0,
	},
		{
			Order:          1,
			WebsiteURL:     "https://google.com",
			HTTPStatusCode: 0,
		},
		{
			Order:          2,
			WebsiteURL:     "https://stackoverflow.com",
			HTTPStatusCode: 0,
		}}
	records := [][]string{{"https://facebook.com"}, {"https://google.com"}, {"https://stackoverflow.com"}}
	actual, _ := SitesMapper(records)
	assert.Equal(t, expected, actual)
}

func Test_SitesMapper_Input_2st_CSV_Format_Should_Get_Correct_Array_Struct(t *testing.T) {
	expected := []healthcheck.WebsiteHealthCheck{{
		Order:          0,
		WebsiteURL:     "https://facebook.com",
		HTTPStatusCode: 0,
	},
		{
			Order:          1,
			WebsiteURL:     "https://google.com",
			HTTPStatusCode: 0,
		},
		{
			Order:          2,
			WebsiteURL:     "https://stackoverflow.com",
			HTTPStatusCode: 0,
		}}
	records := [][]string{{"https://facebook.com", "https://google.com", "https://stackoverflow.com"}}
	actual, _ := SitesMapper(records)
	assert.Equal(t, expected, actual)
}

func Test_SitesMapper_Input_Null_CSV_Get_nil_Struct(t *testing.T) {
	records := [][]string{{}}
	actual, _ := SitesMapper(records)
	assert.Equal(t, nil, actual)
}

func Test_countTotal_Input_3_Green_Sites_0_Red_Sites(t *testing.T) {
	sites := []healthcheck.WebsiteHealthCheck{
		{
			HTTPStatusCode: 200,
		},
		{
			HTTPStatusCode: 200,
		},
		{
			HTTPStatusCode: 200,
		}}
	expectedGreen, expectedRed := 3, 0
	actualGreen, actualRed := countTotal(sites)
	assert.Equal(t, expectedGreen, actualGreen)
	assert.Equal(t, expectedRed, actualRed)
}

func Test_countTotal_Input_0_Green_Sites_5_Red_Sites(t *testing.T) {
	sites := []healthcheck.WebsiteHealthCheck{
		{
			HTTPStatusCode: 400,
		},
		{
			HTTPStatusCode: 500,
		},
		{
			HTTPStatusCode: 503,
		},
		{
			HTTPStatusCode: 0,
		},
		{
			HTTPStatusCode: 304,
		},
	}
	expectedGreen, expectedRed := 0, 5
	actualGreen, actualRed := countTotal(sites)
	assert.Equal(t, expectedGreen, actualGreen)
	assert.Equal(t, expectedRed, actualRed)

}

func Test_countTotal_Input_1_Green_Sites_1_Red_Sites(t *testing.T) {
	sites := []healthcheck.WebsiteHealthCheck{
		{
			HTTPStatusCode: 200,
		},
		{
			HTTPStatusCode: 404,
		}}
	expectedGreen, expectedRed := 1, 1
	actualGreen, actualRed := countTotal(sites)
	assert.Equal(t, expectedGreen, actualGreen)
	assert.Equal(t, expectedRed, actualRed)
}

func Test_fillPrefixURL_Input_No_HTTP_Prefix(t *testing.T) {
	expected := "https://www.twitter.com"
	actual := fillPrefixURL("www.twitter.com")
	assert.Equal(t, expected, actual)
}

func Test_fillPrefixURL_Input_Have_HTTP_Prefix(t *testing.T) {
	expected := "http://www.google.com"
	actual := fillPrefixURL("http://www.google.com")
	assert.Equal(t, expected, actual)
}

func Test_fillPrefixURL_Input_Empty_String(t *testing.T) {
	expected := ""
	actual := fillPrefixURL("")
	assert.Equal(t, expected, actual)
}
