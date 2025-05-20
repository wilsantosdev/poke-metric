// go
package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestApplicationJsonMiddleware_SetsContentType(t *testing.T) {
	// Dummy handler to check if it is called
	handlerCalled := false
	dummyHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handlerCalled = true
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"ok":true}`))
	})

	// Wrap with middleware
	wrapped := ApplicationJsonMiddleware(dummyHandler)

	// Create request and recorder
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	// Serve
	wrapped.ServeHTTP(rec, req)

	// Check Content-Type header
	if got := rec.Header().Get("Content-Type"); got != "application/json" {
		t.Errorf("expected Content-Type application/json, got %q", got)
	}

	// Check handler was called
	if !handlerCalled {
		t.Error("expected inner handler to be called")
	}

	// Check response body
	expectedBody := `{"ok":true}`
	if rec.Body.String() != expectedBody {
		t.Errorf("expected body %q, got %q", expectedBody, rec.Body.String())
	}
}