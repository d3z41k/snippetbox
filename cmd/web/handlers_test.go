package main

import (
	"net/http"
	"testing"
)

func TestPing(t *testing.T) {
	app := newTestApplication(t)
	ts := newTestServer(t, app.routes())
	defer ts.Close()

	rs, err := ts.Client().Get(ts.URL + "/ping")
	if err != nil {
		t.Fatal(err)
	}

	code, _, body := ts.get(t, "/ping")

	if code != http.StatusOK {
		t.Errorf("want %d; got %d", http.StatusOK, rs.StatusCode)
	}

	if string(body) != "OK" {
		t.Errorf("want body to equal %q", "OK")
	}
}
