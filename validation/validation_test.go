package validation

import (
	"testing"
)

func TestIsEmail(t *testing.T) {
	tests := []struct {
		email string
		want  bool
	}{
		{"test@example.com", true},
		{"test.name@example.co.uk", true},
		{"test+label@example.com", true},
		{"test@localhost", true},
		{"test", false},
		{"test@", false},
		{"@example.com", false},
		{"test@.com", false},
		{"", false},
	}

	for _, tt := range tests {
		t.Run(tt.email, func(t *testing.T) {
			if got := IsEmail(tt.email); got != tt.want {
				t.Errorf("IsEmail(%q) = %v, want %v", tt.email, got, tt.want)
			}
		})
	}
}

func TestIsURL(t *testing.T) {
	tests := []struct {
		url  string
		want bool
	}{
		{"http://example.com", true},
		{"https://example.com", true},
		{"https://example.com/path", true},
		{"https://example.com/path?query=1", true},
		{"https://user:pass@example.com", true},
		{"ftp://example.com", true},
		{"example.com", false},    // No scheme
		{"http://", false},        // No host
		{"http://.", false},       // Invalid host
		{"://example.com", false}, // No scheme
		{"", false},               // Empty string
	}

	for _, tt := range tests {
		t.Run(tt.url, func(t *testing.T) {
			if got := IsURL(tt.url); got != tt.want {
				t.Errorf("IsURL(%q) = %v, want %v", tt.url, got, tt.want)
			}
		})
	}
}

func TestIsNumeric(t *testing.T) {
	tests := []struct {
		str  string
		want bool
	}{
		{"123", true},
		{"0", true},
		{"123456789", true},
		{"", false},
		{"123a", false},
		{"a123", false},
		{"1.23", false},
		{"-123", false},
		{"+123", false},
	}

	for _, tt := range tests {
		t.Run(tt.str, func(t *testing.T) {
			if got := IsNumeric(tt.str); got != tt.want {
				t.Errorf("IsNumeric(%q) = %v, want %v", tt.str, got, tt.want)
			}
		})
	}
}

func TestIsAlpha(t *testing.T) {
	tests := []struct {
		str  string
		want bool
	}{
		{"abc", true},
		{"ABC", true},
		{"abcDEF", true},
		{"", false},
		{"abc123", false},
		{"123abc", false},
		{"abc-def", false},
		{"abc def", false},
	}

	for _, tt := range tests {
		t.Run(tt.str, func(t *testing.T) {
			if got := IsAlpha(tt.str); got != tt.want {
				t.Errorf("IsAlpha(%q) = %v, want %v", tt.str, got, tt.want)
			}
		})
	}
}

func TestIsAlphanumeric(t *testing.T) {
	tests := []struct {
		str  string
		want bool
	}{
		{"abc", true},
		{"123", true},
		{"abc123", true},
		{"123abc", true},
		{"", false},
		{"abc-123", false},
		{"abc 123", false},
		{"abc_123", false},
	}

	for _, tt := range tests {
		t.Run(tt.str, func(t *testing.T) {
			if got := IsAlphanumeric(tt.str); got != tt.want {
				t.Errorf("IsAlphanumeric(%q) = %v, want %v", tt.str, got, tt.want)
			}
		})
	}
}

func TestMinLength(t *testing.T) {
	tests := []struct {
		str  string
		min  int
		want bool
	}{
		{"abc", 3, true},
		{"abc", 2, true},
		{"abc", 4, false},
		{"", 1, false},
		{"", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.str, func(t *testing.T) {
			if got := MinLength(tt.str, tt.min); got != tt.want {
				t.Errorf("MinLength(%q, %d) = %v, want %v", tt.str, tt.min, got, tt.want)
			}
		})
	}
}

func TestMaxLength(t *testing.T) {
	tests := []struct {
		str  string
		max  int
		want bool
	}{
		{"abc", 3, true},
		{"abc", 4, true},
		{"abc", 2, false},
		{"", 0, true},
		{"a", 0, false},
	}

	for _, tt := range tests {
		t.Run(tt.str, func(t *testing.T) {
			if got := MaxLength(tt.str, tt.max); got != tt.want {
				t.Errorf("MaxLength(%q, %d) = %v, want %v", tt.str, tt.max, got, tt.want)
			}
		})
	}
}

func TestIsPhoneNumber(t *testing.T) {
	tests := []struct {
		phone string
		want  bool
	}{
		{"+1234567890", true},
		{"1234567890", true},
		{"+123 456 7890", true},
		{"+12345", false},            // Too short
		{"abcdefghij", false},        // Not numeric
		{"+1234567890123456", false}, // Too long
		{"", false},
	}

	for _, tt := range tests {
		t.Run(tt.phone, func(t *testing.T) {
			if got := IsPhoneNumber(tt.phone); got != tt.want {
				t.Errorf("IsPhoneNumber(%q) = %v, want %v", tt.phone, got, tt.want)
			}
		})
	}
}
