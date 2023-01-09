package main

import (
	"net/http/httptest"
	"testing"

	"github.com/MarcReetz/fail-end/testrequest"
)

func Test(t *testing.T) {
	r := SetUpRoutes()
	ts := httptest.NewServer(r)
	defer ts.Close()

	if _, body := testrequest.TestRequest(t, ts, "GET", "/", nil); body != "404 page not found\n" {
		t.Fatalf(body)
	}
}
