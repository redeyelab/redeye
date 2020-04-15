package http

import (
	"net/http/httptest"
	"testing"
)

// TestServer
func TestServer(t *testing.T) {

	r := NewRouter()
	if r == nil {
		t.Error("Failed to get a new router")
	}

	req := httptest.NewRequest("GET", "/api/camera/play", nil)
	w := httptest.NewRecorder()

	// At least run the play handle
	PlayHandler(w, req, nil)
	if w.Code != 200 {
		t.Errorf("PlayHandler returned an error expected (200) go (%d)", w.Code)
	}
}
