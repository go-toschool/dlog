package dlog

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-toschool/dlog"
)

// Client ...
type Client struct {
	httpClient http.Client
	baseURL    string
	service    string
}

// NewClient ...
func NewClient(service string) Client {
	return Client{
		httpClient: http.Client{},
		service:    service,
	}
}

// SetBaseURL ...
func (c *Client) SetBaseURL(url string) {
	c.baseURL = url
}

func (c *Client) do(req *http.Request) error {
	req.Header.Set("Accept", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return errors.New("something went wrong")
	}
	return nil
}

// Info ...
func (c *Client) Info(message string) error {
	msg := &dlog.Message{
		Service: c.service,
		Info:    message,
		Level:   "info",
	}
	body, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", c.baseURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	return c.do(req)
}

// Warn ...
func (c *Client) Warn(message string) error {
	msg := &dlog.Message{
		Service: c.service,
		Warn:    message,
		Level:   "warn",
	}
	body, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", c.baseURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	return c.do(req)
}

// Error ...
func (c *Client) Error(message string) error {
	msg := &dlog.Message{
		Service: c.service,
		Error:   message,
		Level:   "error",
	}
	body, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", c.baseURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	return c.do(req)
}
