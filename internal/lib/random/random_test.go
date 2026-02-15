package random

import (
	"testing"
)

func isAllowed(r rune) bool {
	allowed := []rune("QWERTYUIOPASDFGHJKLZXCVBNMqwertyuiopasdfghjklzxcvbnm1234567890")
	for _, c := range allowed {
		if r == c {
			return true
		}
	}
	return false
}

func TestNewRandomString(t *testing.T) {
	tests := []struct {
		name string
		size int
	}{
		{
			"nothing",
			0,
		},
		{
			"normal",
			5,
		},
		{
			"big",
			100,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewRandomString(tt.size)
			if len(got) != tt.size {
				t.Errorf("len(NewRandomString()) = %v, want %v", got, tt.size)
			}
			for _, r := range got {
				if !isAllowed(r) {
					t.Errorf("unexpected rune %q in %q", r, got)
				}
			}
		})
	}
}
