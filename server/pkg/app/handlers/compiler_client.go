package handlers

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

type CompilerClient struct {
	ClientID     string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
	Script       string `json:"script"`
	Language     string `json:"language"`
	VersionIndex string `json:"versionIndex"`
}

type Response struct {
	Status     string
	StatusCode int
	Header     http.Header
	Body       interface{}
}

//TODO: add to config
var (
	URL string
)

//TODO: add check for empty fields
func (c *CompilerClient) CompilerClient() (*Response, error) {
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

	res := &Response{
		Status:     response.Status,
		StatusCode: response.StatusCode,
		Header:     response.Header,
	}

	if response.Header.Get("Content-Type") == "application/json" {
		var body map[string]string
		err = json.NewDecoder(response.Body).Decode(&body)
		if err != nil {
			return nil, err
		}
		res.Body = body
	} else {
		//send body as text/plain
		bodyB, err := io.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}
		body := string(bodyB)
		res.Body = body
	}
	return res, nil
}
