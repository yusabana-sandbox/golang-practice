package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMyHandler_ServeHTTP(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/handle_test", nil)
	handleTest(w, r)
	rw := w.Result()
	defer rw.Body.Close()

	if rw.StatusCode != http.StatusOK {
		t.Fatal("Unexpected status code")
	}
	b, err := ioutil.ReadAll(rw.Body)

	if err != nil {
		t.Fatal("unexpected error")
	}
	const expected = "Hello, net/http!!"
	if s := string(b); s != expected {
		t.Fatalf("unexpected response: %s", s)
	}
}
