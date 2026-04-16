package builder

import (
	"strings"
	"testing"
)

func TestBuildRequiresProvider(t *testing.T) {
	t.Parallel()

	_, err := NewLoggerBuilder().Build()
	if err == nil || !strings.Contains(err.Error(), "at least one provider") {
		t.Fatalf("expected missing provider error, got %v", err)
	}
}

func TestBuildReturnsFileValidationError(t *testing.T) {
	t.Parallel()

	_, err := NewLoggerBuilder().AddFile("").Build()
	if err == nil || !strings.Contains(err.Error(), "file path is required") {
		t.Fatalf("expected file validation error, got %v", err)
	}
}
