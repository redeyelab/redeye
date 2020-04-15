package http

import (
	"fmt"
	"net/http/httptest"
	"testing"
)

// TestServer
func TestServer(t *testing.T) {

	r := NewRouter()
	if r == nil {
		t.Error("Failed to get a new router")
	}

	req := httptest.NewRequest("GET", "http://foo.com/bar", nil)
	w := httptest.NewRecorder()

	PlayHandler(w, req, nil)
	fmt.Printf("%+v\n", w)
}
