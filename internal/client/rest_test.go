package client

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testThing struct {
	Id int `json:"id"`
	Name string `json:"name"`
}

type testServerMessage struct {
	Message string
}

var (
	bearerTest = "bearer test"
	getSuccessThing = testThing{Id: 1, Name: "ken"}
	putSuccess = testServerMessage{Message: "accepted"}
	testHeaders = map[string]string{"authorization": bearerTest}
)

func TestMakeHttpRequest_Get_BasicSuccess(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bts, _ := json.Marshal(&getSuccessThing)
		w.WriteHeader(http.StatusOK)
		w.Write(bts)
	}))
	cfg := RequestConfig{Url: ts.URL, Method: http.MethodGet}

	response, err := MakeHttpRequest[testThing](cfg)

	assert.Nil(t, err)
	assert.Equal(t, getSuccessThing, *response)
}

func TestMakeHttpRequest_Get_BasicErrorOnDo(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
	cfg := RequestConfig{Url: ts.URL[1:], Method: http.MethodGet}

	response, err := MakeHttpRequest[testThing](cfg)

	assert.NotNil(t, err)
	assert.Nil(t, response)
}

func TestMakeHttpRequest_Get_BasicErrorOnRequest(t *testing.T) {
	httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
	cfg := RequestConfig{Url: "\t", Method: http.MethodGet}

	response, err := MakeHttpRequest[testThing](cfg)

	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "invalid control character in URL")
	assert.Nil(t, response)
}

func TestMakeHttpRequest_Get_BasicErrorOnParse(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("<data>lol</data>"))
	}))
	cfg := RequestConfig{Url: ts.URL, Method: http.MethodGet}

	response, err := MakeHttpRequest[testThing](cfg)

	assert.NotNil(t, err)
	assert.Nil(t, response)
}

func TestMakeHttpRequest_Put_BasicSuccess(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusAccepted)
		w.Write([]byte(`{"message":"accepted"}`))
	}))
	cfg := RequestConfig{Url: ts.URL, Method: http.MethodPut, Body: `{"id":1,"name":"test"}`}

	response, err := MakeHttpRequest[testServerMessage](cfg)

	assert.Nil(t, err)
	assert.Equal(t, putSuccess, *response)
}

type testBadBody struct {}
func (t testBadBody) Read([]byte) (n int, err error) { 
	return 0, errors.New("lol")
}

func Test_handleBody_BadBody(t *testing.T) {
	_, err := handleBody[any](io.NopCloser(testBadBody{}))
	assert.NotNil(t, err)
}

func TestMakeHttpRequest_Get_AuthHeader(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("authorization") != "bearer test" {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("unauthorized"))
			return
		}
		bts, _ := json.Marshal(&getSuccessThing)
		w.WriteHeader(http.StatusOK)
		w.Write(bts)
	}))
	goodCfg := RequestConfig{Url: ts.URL, Method: http.MethodGet, Headers: testHeaders}
	response, err := MakeHttpRequest[testThing](goodCfg)

	assert.Nil(t, err)
	assert.Equal(t, getSuccessThing, *response)

	badCfg := RequestConfig{Url: ts.URL, Method: http.MethodGet}
	response, err = MakeHttpRequest[testThing](badCfg)

	assert.NotNil(t, err)
	assert.Nil(t, response)
}
