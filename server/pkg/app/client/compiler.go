package client

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"
)

type Compiler struct {
	ClientIDsssss string `json:"clientId"`
	ClientSecret  string `json:"clientSecret"`
	Script        string `json:"script"`
	Language      string `json:"language"`
	VersionIndex  string `json:"versionIndex"`
}

type Response struct {
	Status     string
	StatusCode int
	Header     http.Header
	Body       map[string]string
}

func (c *Compiler) CompilerClient() (*Response, error) {
	URL := "https://api.jdoodle.com/v1/execute"
	return c.CompilerClientURL(URL)
}

func (c *Compiler) CompilerClientURL(URL string) (*Response, error) {
	data, err := json.Marshal(c)
	if err != nil {
		return nil, err
	}

	client := http.Client{Timeout: time.Duration(5) * time.Second}

	request, err := http.NewRequest("POST", URL, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	request.Header.Add("Accept", "application/json")
	request.Header.Add("Content-Type", "application/json")
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	var body map[string]string
	err = json.NewDecoder(response.Body).Decode(&body)
	if err != nil {
		return nil, err
	}

	res := &Response{
		Status:     response.Status,
		StatusCode: response.StatusCode,
		Header:     response.Header,
		Body:       body,
	}
	return res, nil
}
