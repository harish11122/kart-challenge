package tests

import (
	"backend-challenge/utils"
	"testing"
)

func TestIsValidPromoCode(t *testing.T) {
	utils.LoadValidPromoCodes()

	tests := []struct {
		name     string
		code     string
		expected bool
	}{
		{"Valid code", "HAPPYHRS", true},
		{"Invalid code", "INVALIDCODE", false},
		{"Too short", "SHORT", false},
		{"Too long", "THISISWAYTOOLONGCODE", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := utils.IsValidPromoCode(tt.code); got != tt.expected {
				t.Errorf("IsValidPromoCode(%q) = %v; want %v", tt.code, got, tt.expected)
			}
		})
	}
}
