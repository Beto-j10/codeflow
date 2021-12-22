package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"reflect"
	"strings"
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

func (c *CompilerClient) CompilerClient() (*Response, error) {

	URL = "https://api.jdoodle.com/v1/execute"
	// Check all data is filled in using reflect
	v := reflect.ValueOf(*c)
	for i := 0; i < v.NumField(); i++ {

		// checks if the field value equals its zero value
		if reflect.DeepEqual(v.Field(i).Interface(), reflect.Zero(v.Field(i).Type()).Interface()) {

			tag := fmt.Sprintf("%v", v.Type().Field(i).Tag)
			field := strings.Split(tag, ":")[1] // e.g, `json:"nameTag"` returns: "nameTag"
			err := fmt.Sprintf("bad request: %s field empty", field)

			return nil, errors.New(err)
		}
	}

	data, err := json.Marshal(c)
	if err != nil {
		log.Printf("error marshal %v", err)
		return nil, err
	}

	client := http.Client{Timeout: time.Duration(5) * time.Second}

	request, err := http.NewRequest("POST", URL, bytes.NewBuffer(data))
	if err != nil {
		log.Printf("error newRequest %v", err)
		return nil, err
	}

	request.Header.Add("Accept", "application/json")
	request.Header.Add("Content-Type", "application/json")

	response, err := client.Do(request)
	if err != nil {
		log.Printf("error client.Do request %v", err)
		return nil, err
	}

	defer response.Body.Close()

	res := &Response{
		Status:     response.Status,
		StatusCode: response.StatusCode,
		Header:     response.Header,
	}

	//send body as text/plain
	if response.Header.Get("Content-Type") == "text/plain" {
		bodyB, err := io.ReadAll(response.Body)
		if err != nil {
			log.Printf("error readAll body %v", err)
			return nil, err
		}
		body := string(bodyB)
		res.Body = body
		return res, nil
	}

	//send body as JSON
	var body map[string]interface{}
	err = json.NewDecoder(response.Body).Decode(&body)
	if err != nil {
		log.Printf("error newDecoder body %v", err)
		return nil, err
	}
	res.Body = body
	return res, nil
}
