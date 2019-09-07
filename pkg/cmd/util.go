package cmd

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func messageWithUnderline(msg string) string {
	if msg == "" {
		return ""
	}
	underline := strings.Repeat("=", utf8.RuneCountInString(msg))

	return fmt.Sprintf("%s\n%s\n", msg, underline)
}

func dryRunBanner() string {
	return "--- Dry Run ---"
}
