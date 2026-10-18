package util

import "testing"

func TestIsLocalOrIPHost(t *testing.T) {
	tests := []struct {
		host string
		want bool
	}{
		{"localhost", true},
		{"localhost:8080", true},
		{"app.localhost:8080", true},
		{"127.0.0.1", true},
		{"192.0.2.1:8080", true},
		{"::1", true},
		{"[::1]", true},
		{"[::1]:8080", true},
		{"2001:db8::1", true},
		{"[2001:db8::1]:443", true},
		{"::ffff:192.0.2.1", true},
		{"[fe80::1%eth0]:8080", true},
		{"fe80::1%eth0", true},
		{" [::1]:8080 ", true},
		{"example.com", false},
		{"example.com:443", false},
		{"localhost.example.com", false},
		{"", false},
		{":8080", false},
	}
	for _, tt := range tests {
		t.Run(tt.host, func(t *testing.T) {
			if got := IsLocalOrIPHost(tt.host); got != tt.want {
				t.Errorf("IsLocalOrIPHost(%q) = %v, want %v", tt.host, got, tt.want)
			}
		})
	}
}

func TestIsLocalOrIPURL(t *testing.T) {
	tests := []struct {
		url  string
		want bool
	}{
		{"http://localhost:8080", true},
		{"http://app.localhost", true},
		{"http://192.0.2.1:8080", true},
		{"http://[::1]", true},
		{"http://[::1]:8080/path", true},
		{"https://[2001:db8::1]:443", true},
		{"http://[fe80::1%25eth0]:8080", true},
		{"https://example.com", false},
		{"https://example.com:443/path", false},
		{"/relative/path", false},
		{"", false},
		{"://", false},
	}
	for _, tt := range tests {
		t.Run(tt.url, func(t *testing.T) {
			if got := IsLocalOrIPURL(tt.url); got != tt.want {
				t.Errorf("IsLocalOrIPURL(%q) = %v, want %v", tt.url, got, tt.want)
			}
		})
	}
}
