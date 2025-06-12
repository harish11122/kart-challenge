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

// LoadValidPromoCodes reads the valid coupons file once and loads into memory
func LoadValidPromoCodes() {
	once.Do(func() {
		file := "utils/valid_coupons.txt"

		f, err := os.Open(file)
		if err != nil {
			panic(fmt.Sprintf("failed to open %s: %v", file, err))
		}
		defer f.Close()

		scanner := bufio.NewScanner(f)

		for scanner.Scan() {
			code := strings.TrimSpace(scanner.Text())
			validPromoCodes[code] = true
		}

		if err := scanner.Err(); err != nil {
			panic(fmt.Sprintf("error reading %s: %v", file, err))
		}
	})
}

// IsValidPromoCode returns true if the code is valid and matches length rules
func IsValidPromoCode(code string) bool {
	LoadValidPromoCodes()
	return validPromoCodes[code]
}
