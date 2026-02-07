//go:build windows

package clipboard_test

import (
	"context"
	"os"
	"runtime"
	"testing"
	"time"

	"golang.design/x/clipboard"
)

func TestClipboardWatchImageWindowsFormats(t *testing.T) {
	if runtime.GOOS != "windows" {
		t.Skip("windows only")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	data, err := os.ReadFile("tests/testdata/clipboard.png")
	if err != nil {
		t.Fatalf("failed to read gold file: %v", err)
	}

	// clear clipboard first
	clipboard.Write(clipboard.FmtText, []byte(""))

	ch := clipboard.Watch(ctx, clipboard.FmtImage)
	clipboard.Write(clipboard.FmtImage, data)

	select {
	case <-ctx.Done():
		t.Fatalf("timed out waiting for image watch")
	case got := <-ch:
		if len(got) == 0 {
			t.Fatalf("got empty image bytes")
		}
	}
}
