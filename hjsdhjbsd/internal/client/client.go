// Package client implements functions to demonstrate godoc features
//
// The PostJson returns error and sen yr text to dest
//
// PostJson() error
package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

const (
	jsonContentType = "application/json"
)

var (
	ErrorIncorrectBodyFormat = errors.New("incorrect body format")
	ErrorSendRequest         = errors.New("couldn't send request")
)

type Rarer interface {
	Rare() string
}

type HttpStatusError struct {
	status int
}

func NewHttpStatusError(status int) error {
	return &HttpStatusError{status}
}

func (s *HttpStatusError) Error() string {
	return fmt.Sprintf("status code: %d", s.status)
}

func (s *HttpStatusError) Status() int {
	return s.status
}

// PostJson returns a error only
func PostJson(client *http.Client, url, body string) (err error) {
	
	js := make(map[string]interface{}, 10)

	err = json.Unmarshal([]byte(body), &js)
	if err != nil {
		return fmt.Errorf("%w: %s", ErrorIncorrectBodyFormat, err.Error())
	}

	res, err := client.Post(url, jsonContentType, bytes.NewReader([]byte(body)))
	if err != nil {
		return fmt.Errorf("%w: %s", ErrorSendRequest, err.Error())
	}

	if res.StatusCode != http.StatusOK {
		return NewHttpStatusError(res.StatusCode)
	}

	return
}

