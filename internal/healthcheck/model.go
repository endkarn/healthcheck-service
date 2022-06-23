package healthcheck

type WebsiteHealthCheck struct {
	Order          int
	WebsiteURL     string
	HTTPStatusCode int
}
