package main

import (
	"strings"
	"testing"
)

func TestPromptRandom(t *testing.T) {
	categories := []string{}
	prompt := buildPrompt(categories)
	expected := "Tell me a random short joke."
	if prompt != expected {
		t.Errorf("Expected '%s', got '%s'", expected, prompt)
	}
}

func TestPromptSingleCategory(t *testing.T) {
	categories := []string{"geek"}
	prompt := buildPrompt(categories)
	expected := "Tell me a short joke about geek."
	if prompt != expected {
		t.Errorf("Expected '%s', got '%s'", expected, prompt)
	}
}

func TestPromptMultipleCategories(t *testing.T) {
	categories := []string{"geek", "running"}
	prompt := buildPrompt(categories)
	expected := "Tell me a short joke combining the themes of geek and running."
	if !strings.EqualFold(prompt, expected) {
		t.Errorf("Expected '%s', got '%s'", expected, prompt)
	}
}
