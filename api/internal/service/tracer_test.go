package service

import (
	"context"
	"os"
	"testing"
)

// TestNewTracerService_Success tests that NewTracerService returns a shutdown function and does not panic.
func TestNewTracerService_Success(t *testing.T) {
	// Set a dummy endpoint to avoid nil endpoint error
	os.Setenv("JAEGER_ENDPOINT", "localhost:4317")
	defer os.Unsetenv("JAEGER_ENDPOINT")

	ctx := context.Background()
	shutdown := NewTracerService(ctx)
	if shutdown == nil {
		t.Fatal("expected shutdown function, got nil")
	}
	// Call shutdown to ensure it does not panic
	shutdown()
}
