package main

import (
	// "bytes"
	"github.com/Akkshatt/go_snippet_box/internals/assert"
	// "io"
	// "log"
	"net/http"
	// "net/http/httptest"
	"testing"
)
func TestPing(t *testing.T) {
 app := newTestApplication(t)
 ts := newTestServer(t, app.routes())
 defer ts.Close()
 code, _, body := ts.get(t, "/ping")
 assert.Equal(t, code, http.StatusOK)
 assert.Equal(t, body, "OK")
}
