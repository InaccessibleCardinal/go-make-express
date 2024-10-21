package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type RequestConfig struct {
	Body string
	Method string
	Url string
	Headers map[string]string
}

func MakeHttpRequest[T any](cfg RequestConfig) (*T, error) {
	req, err := http.NewRequest(cfg.Method, cfg.Url, getBodyReader(cfg.Body))
	if err != nil {
		return nil, err
	}
	addRequestHeaders(req, cfg.Headers)

	res, err := getResponse(http.DefaultClient, req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	return handleBody[T](res.Body)
}

func getBodyReader(b string) io.Reader {
	if b != "" {
		return strings.NewReader(b)
	}
	return nil
}

func addRequestHeaders(req *http.Request, headers map[string]string) {
	if len(headers) == 0 {
		return
	}
	for k, v := range headers {
		req.Header.Add(k, v)
	}
}

func getResponse(client *http.Client, req *http.Request) (*http.Response, error) {
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode > 299 {
		return nil, fmt.Errorf("bad status %d", res.StatusCode)
	}
	return res, nil
}


func handleBody[T any](body io.ReadCloser) (*T, error) {
	bts, err := io.ReadAll(body)
	if err != nil {
		return nil, err
	}
	return parseType[T](bts)
}

func parseType[T any](bts []byte) (*T, error) {
	var entity T
	err := json.Unmarshal(bts, &entity)
	if err != nil {
		return nil, err
	}
	return &entity, nil
}