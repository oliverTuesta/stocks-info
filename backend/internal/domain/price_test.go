package domain

import (
	"encoding/json"
	"testing"
)

func TestPrice_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected Price
		wantErr  bool
	}{
		{
			name:     "valid price with dollar sign",
			input:    `"$123.45"`,
			expected: Price(12345), // 123.45 * 100
			wantErr:  false,
		},
		{
			name:     "valid price with comma",
			input:    `"1,234.56"`,
			expected: Price(123456), // 1234.56 * 100
			wantErr:  false,
		},
		{
			name:     "valid price with dollar sign and comma",
			input:    `"$1,234.56"`,
			expected: Price(123456), // 1234.56 * 100
			wantErr:  false,
		},
		{
			name:     "valid price without formatting",
			input:    `"99.99"`,
			expected: Price(9999), // 99.99 * 100
			wantErr:  false,
		},
		{
			name:     "valid price with spaces",
			input:    `" $ 123.45 "`,
			expected: Price(12345), // 123.45 * 100
			wantErr:  false,
		},
		{
			name:     "valid integer price",
			input:    `"100"`,
			expected: Price(10000), // 100 * 100
			wantErr:  false,
		},
		{
			name:     "valid zero price",
			input:    `"0"`,
			expected: Price(0),
			wantErr:  false,
		},
		{
			name:     "invalid non-numeric string",
			input:    `"invalid"`,
			expected: Price(0),
			wantErr:  true,
		},
		{
			name:     "invalid empty string",
			input:    `""`,
			expected: Price(0),
			wantErr:  true,
		},
		{
			name:     "invalid JSON format",
			input:    `123.45`,
			expected: Price(0),
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var p Price
			err := json.Unmarshal([]byte(tt.input), &p)

			if (err != nil) != tt.wantErr {
				t.Errorf("Price.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if p != tt.expected {
				t.Errorf("Price.UnmarshalJSON() = %v, want %v", p, tt.expected)
			}
		})
	}
}

func TestPrice_UnmarshalJSON_EdgeCases(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected Price
		wantErr  bool
	}{
		{
			name:     "very large number",
			input:    `"999999999.99"`,
			expected: Price(99999999999), // 999999999.99 * 100
			wantErr:  false,
		},
		{
			name:     "negative price",
			input:    `"-123.45"`,
			expected: Price(-12345), // -123.45 * 100
			wantErr:  false,
		},
		{
			name:     "price with multiple commas",
			input:    `"1,234,567.80"`,
			expected: Price(123456780), // 1234567.89 * 100
			wantErr:  false,
		},
		{
			name:     "price with multiple dollar signs",
			input:    `"$$123.45"`,
			expected: Price(12345), // Should still work, removes all $ signs
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var p Price
			err := json.Unmarshal([]byte(tt.input), &p)

			if (err != nil) != tt.wantErr {
				t.Errorf("Price.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if p != tt.expected {
				t.Errorf("Price.UnmarshalJSON() = %v, want %v", p, tt.expected)
			}
		})
	}
}
