package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"server/config"
	"strings"
	"time"
)

// CompilerClient handles POST /compiler
type CompilerClient struct {
	ClientID     string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
	Script       string `json:"script"`
	Language     string `json:"language"`
	VersionIndex string `json:"versionIndex"`
}

// Response is the response from the compiler
type Response struct {
	Status     string
	StatusCode int
	Header     http.Header
	Body       interface{}
}

// CompilerClient does a POST request to the compiler and returns the response
func (c *CompilerClient) compilerClient(config *config.CompilerConfig) (*Response, error) {

	URL := config.URL
	c.ClientID = config.ClientId
	c.ClientSecret = config.ClientSecret

	// Check all data is filled in using reflect
	v := reflect.ValueOf(*c)
	for i := 0; i < v.NumField(); i++ {

		// Checks if the field value equals its zero value
		if reflect.DeepEqual(v.Field(i).Interface(), reflect.Zero(v.Field(i).Type()).Interface()) {

			tag := fmt.Sprintf("%v", v.Type().Field(i).Tag) // get the field tag
			field := strings.Split(tag, ":")[1]             // get the field name, e.g., `json:"nameTag"` returns: "nameTag"
			err := fmt.Sprintf("bad request: %s field empty", field)

			return nil, errors.New(err)
		}
	}

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

	// Send body as text/plain
	if response.Header.Get("Content-Type") == "text/plain" {
		bodyB, err := io.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}
		body := string(bodyB)
		res.Body = body
		return res, nil
	}

	// Send body as JSON
	var body map[string]interface{}
	err = json.NewDecoder(response.Body).Decode(&body)
	if err != nil {
		return nil, err
	}
	res.Body = body
	return res, nil
}
