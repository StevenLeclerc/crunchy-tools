package crunchyTools

import (
	"crypto/md5"
	"math"
	"strings"
	"testing"
)

// Returns a [16]byte checksum for a given error message.
func TestCheckSumErrorMsgReturnsChecksum(t *testing.T) {
	errorMsg := "This is an error message"
	expectedChecksum := md5.Sum([]byte(errorMsg))

	actualChecksum := GenMd5FromString(errorMsg)

	if actualChecksum != expectedChecksum {
		t.Errorf("Expected checksum %v, but got %v", expectedChecksum, actualChecksum)
	}
}

// Handles error messages with maximum length (2^31 - 1 bytes).
func TestCheckSumErrorMsgHandlesMaxLength(t *testing.T) {
	errorMsg := strings.Repeat("a", math.MaxInt32-1)
	expectedChecksum := md5.Sum([]byte(errorMsg))

	actualChecksum := GenMd5FromString(errorMsg)

	if actualChecksum != expectedChecksum {
		t.Errorf("Expected checksum %v, but got %v", expectedChecksum, actualChecksum)
	}
}
