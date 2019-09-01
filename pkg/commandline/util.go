package commandline

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func printWithUnderline(msg string) {
	fmt.Println(msg)
	fmt.Println(strings.Repeat("=", utf8.RuneCountInString(msg)))
}

func printDryRunBanner() {
	fmt.Println("--- Dry Run ---")
}
