package main

import (
	"net/http"
	"testing"

	"github.com/deferet/snippetbox/internal/assert"
)

func TestPing(t *testing.T) {
	app := newTestApplication()

	ts := newTestServer(t, app.routes())
	defer ts.Close()

	statusCode, _, body := ts.get(t, "/ping")

	assert.Equal(t, statusCode, http.StatusOK)

	assert.Equal(t, string(body), "OK")
}
