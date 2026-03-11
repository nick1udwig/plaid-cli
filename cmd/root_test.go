package cmd

import (
	"io"
	"testing"
)

func TestRootCommandUsesInjectedVersion(t *testing.T) {
	prev := version
	version = "9.9.9-test"
	t.Cleanup(func() {
		version = prev
	})

	cmd := NewRootCmd(&Options{
		Stdout:   io.Discard,
		Stderr:   io.Discard,
		StateDir: t.TempDir(),
	})

	if cmd.Version != "9.9.9-test" {
		t.Fatalf("expected root version %q, got %q", "9.9.9-test", cmd.Version)
	}
}
