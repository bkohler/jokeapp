package main

import (
	"fmt"
	"strings"
)

// buildPrompt constructs the prompt string based on joke categories
func buildPrompt(categories []string) string {
	if len(categories) == 0 {
		return "Tell me a random short joke."
	}
	if len(categories) == 1 {
		return fmt.Sprintf("Tell me a short joke about %s.", categories[0])
	}
	return fmt.Sprintf("Tell me a short joke combining the themes of %s.", strings.Join(categories, " and "))
}
