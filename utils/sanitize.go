package utils

import (
	"regexp"
	"strings"
)

func SanitizeBarcodeValue(input string) string {
	input = strings.TrimSpace(input)
	re := regexp.MustCompile(`[^a-zA-Z0-9\-\.\s]`)
	sanitized := re.ReplaceAllString(input, "")
	if len(sanitized) > 100 {
		sanitized = sanitized[:100]
	}
	return sanitized
}

func SanitizeQRValue(input string) string {
	const maxLength = 1000

	value := []rune(strings.TrimSpace(input))
	if len(value) > maxLength {
		value = value[:maxLength]
	}

	return string(value)
}
