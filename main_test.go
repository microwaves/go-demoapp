package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

var writer *httptest.ResponseRecorder

func setup() {
	writer = httptest.NewRecorder()
}

func TestIndexGet200(t *testing.T) {
	setup()
	expected := http.StatusOK

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal("Creating 'GET /' request failed!")
	}

	indexHandler(writer, req)
	resp := writer.Result()

	if resp.StatusCode != expected {
		t.Fatal("Server error: Returned code:", resp.StatusCode, "instead of:", expected)
	}
}

func TestIndexBody(t *testing.T) {
	setup()
	expected := "Heey! This is supposed to be a test. :-)"

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal("Creating 'GET /' request failed!")
	}

	indexHandler(writer, req)

	resp := writer.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	if string(body) != expected {
		t.Fatal("Error: Returned body:", string(body), "instead of:", expected)
	}
}
