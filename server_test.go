package main

import (
	"net/http/httptest"
	"testing"
)

func Test(t *testing.T) {
	r := SetUpRoutes()
	ts := httptest.NewServer(r)
	defer ts.Close()
}
