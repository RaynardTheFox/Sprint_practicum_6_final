package service

import (
	"testing"
)

func TestConvert_TextToMorse(t *testing.T) {
	input := "Привет"
	result, err := Convert(input)
	if err != nil {
		t.Fatalf("Convert returned error: %v", err)
	}
	
	hasMorseChars := false
	for _, r := range result {
		if r == '.' || r == '-' {
			hasMorseChars = true
			break
		}
	}
	
	if !hasMorseChars {
		t.Errorf("Expected Morse code, got: %s", result)
	}
}

func TestConvert_MorseToText(t *testing.T) {
	input := ".--. .-. .. .-- . -"
	result, err := Convert(input)
	if err != nil {
		t.Fatalf("Convert returned error: %v", err)
	}
	
	if isMorseCode(result) {
		t.Errorf("Expected text, got Morse code: %s", result)
	}
}

func TestConvert_EmptyString(t *testing.T) {
	result, err := Convert("")
	if err != nil {
		t.Fatalf("Convert returned error: %v", err)
	}
	
	if result != "" {
		t.Errorf("Expected empty string, got: %s", result)
	}
}

func TestIsMorseCode(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"Valid Morse", ".--. .-. .. .-- . -", true},
		{"Valid Morse with spaces", "- . ... -", true},
		{"Text", "Привет", false},
		{"Empty string", "", false},
		{"Only spaces", "   ", false},
		{"Mixed", "Привет .--", false},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isMorseCode(tt.input)
			if result != tt.expected {
				t.Errorf("isMorseCode(%q) = %v, expected %v", tt.input, result, tt.expected)
			}
		})
	}
}

