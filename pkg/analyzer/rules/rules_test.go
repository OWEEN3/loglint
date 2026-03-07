package rules

import (
	"testing"
)

func TestIsEmpty(t *testing.T) {
	tests := []struct {
		name string
		msg  string
		want bool
	}{
		{"empty string", "", true},
		{"only spaces", "   ", true},
		{"spaces and tabs", " \t\n", true},
		{"non-empty", "hello", false},
		{"non-empty with spaces", " hello ", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsEmpty(tt.msg); got != tt.want {
				t.Errorf("IsEmpty(%q) = %v, want %v", tt.msg, got, tt.want)
			}
		})
	}
}

func TestIsFirstLower(t *testing.T) {
	tests := []struct {
		name string
		msg  string
		want bool
	}{
		{"lowercase a", "apple", true},
		{"lowercase z", "zebra", true},
		{"uppercase", "Apple", false},
		{"digit first", "1apple", false},
		{"cyrillic lowercase", "яблоко", false},
		{"cyrillic uppercase", "Яблоко", false},
		{"special char", "!apple", false},
		{"space first", " apple", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsFirstLower(tt.msg); got != tt.want {
				t.Errorf("IsFirstLower(%q) = %v, want %v", tt.msg, got, tt.want)
			}
		})
	}
}

func TestIsValidChars(t *testing.T) {
	tests := []struct {
		name string
		msg  string
		want bool
	}{
		{"only letters", "helloWorld", true},
		{"letters and digits", "abc123", true},
		{"with spaces", "hello world 123", true},
		{"with punctuation", "helloworld!", false},
		{"with emoji", "hello 🚀", false},
		{"cyrillic", "привет", false},
		{"special symbols", "test@domain", false},
		{"newline", "hello\nworld", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidChars(tt.msg); got != tt.want {
				t.Errorf("IsValidChars(%q) = %v, want %v", tt.msg, got, tt.want)
			}
		})
	}
}

func TestContainsSensitive(t *testing.T) {
	tests := []struct {
		name string
		msg  string
		want bool
	}{
		{"simple password", "user password is 123", true},
		{"password alone", "password", true},
		{"token in text", "your token is abc", true},
		{"api key", "api_key=123", true},
		{"key alone", "this is a key", true},
		{"secret", "my secret data", true},
		{"passphrase", "passphrase is long", true},
		{"ssh-rsa", "ssh-rsa AAAAB3...", true},
		{"multiple sensitive", "user password and token", true},

		// Думаю в такие случаи лучше выводить предупреждение, чем нет.
		{"word boundary pass", "compass", true},               
		{"word boundary password", "password123", true},       
		{"token as part", "tokenized", true},                  
		{"key in domain", "monkey", true},                     

		{"uppercase password", "MY PASSWORD IS SECRET", true},
		{"mixed case token", "This is a ToKeN", true},
		{"lowercase key", "my Key123", true},

		{"no sensitive", "hello world", false},
		{"ssn", "ssn: 123-45-6789", false},
		{"credit card", "credit card number", false},
		{"cvv", "cvv 123", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ContainsSensitive(tt.msg)
			if tt.want && got == "" {
				t.Errorf("ContainsSensitive(%q) = \"\", want match", tt.msg)
			}
			if !tt.want && got != "" {
				t.Errorf("ContainsSensitive(%q) = %q, want no match", tt.msg, got)
			}
		})
	}
}