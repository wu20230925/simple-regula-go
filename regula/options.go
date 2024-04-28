package regula

import "time"

// Option is the type of constructor options for NewClient(...).
type Option func(*Client)

// WithLicense configures a Maps API client with an API Key
func WithLicense(license string) Option {
	return func(c *Client) {
		c.license = license
	}
}

// WithBaseURL configures a Maps API client with a custom base url
func WithBaseURL(baseURL string) Option {
	return func(c *Client) {
		c.cli.SetBaseURL(baseURL)
	}
}

// WithRetryCount configures a retry count
func WithRetryCount(retryCount int) Option {
	return func(c *Client) {
		c.cli.SetRetryCount(retryCount)
	}
}

// WithRetryMaxWaitTime configures max wait time to sleep before retrying
func WithRetryMaxWaitTime(maxWaitTime time.Duration) Option {
	return func(c *Client) {
		c.cli.SetRetryMaxWaitTime(maxWaitTime)
	}
}
