package adaptive

import (
	"os"
	"testing"
)

func TestGetDiskInfo(t *testing.T) {
	info, err := GetDiskInfo(".")
	if err != nil {
		t.Fatalf("GetDiskInfo failed: %v", err)
	}

	if info.TotalBytes == 0 {
		t.Fatal("expected total disk space to be greater than zero")
	}

	if info.FreeBytes == 0 {
		t.Fatal("expected free disk space to be greater than zero")
	}

	if info.FreePercent < 0 || info.FreePercent > 100 {
		t.Fatalf("free percentage out of range: %.2f", info.FreePercent)
	}

	t.Logf(
		"disk: free=%d bytes, total=%d bytes, free=%.2f%%",
		info.FreeBytes,
		info.TotalBytes,
		info.FreePercent,
	)

	// Confirm the test is actually running from a usable directory.
	if _, err := os.Stat("."); err != nil {
		t.Fatalf("current directory is not accessible: %v", err)
	}
}
