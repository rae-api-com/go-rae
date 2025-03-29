package rae

import "time"

type ClientOption func(*Client)

func WithTimeout(timeout time.Duration) ClientOption {
	return func(c *Client) {
		c.timeout = timeout
	}
}

func WithVersion(version string) ClientOption {
	return func(c *Client) {
		c.version = version
	}
}
