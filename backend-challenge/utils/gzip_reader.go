package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

var (
	validPromoCodes = make(map[string]bool)
	once            sync.Once
)

// LoadValidPromoCodes reads and filters all promo codes from plain text files
func LoadValidPromoCodes() {
	once.Do(func() {
		files := []string{"couponbase1", "couponbase2", "couponbase3"}

		for _, file := range files {
			readAndFilterFile(file)
		}
	})
}

func readAndFilterFile(path string) {
	f, err := os.Open(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening %s: %v\n", path, err)
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		code := strings.TrimSpace(scanner.Text())
		if isValidLength(code) {
			validPromoCodes[code] = true
		}
	}
}

func isValidLength(code string) bool {
	return len(code) >= 8 && len(code) <= 10
}

// IsValidPromoCode returns true if the promo code is valid
func IsValidPromoCode(code string) bool {
	LoadValidPromoCodes()
	return validPromoCodes[code]
}
