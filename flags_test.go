package main

import (
	"flag"
	"strings"
	"testing"
)

// Helper to reset and parse flags with given args
func parseFlags(args []string) (help bool, categories []string) {
	fs := flag.NewFlagSet("test", flag.ContinueOnError)
	helpFlag := fs.Bool("h", false, "help")
	geekFlag := fs.Bool("g", false, "geek")
	nerdFlag := fs.Bool("n", false, "nerd")
	sportFlag := fs.Bool("s", false, "sport")
	runFlag := fs.Bool("r", false, "run")

	_ = fs.Parse(args)

	if *helpFlag {
		help = true
	}
	if *geekFlag {
		categories = append(categories, "geek")
	}
	if *nerdFlag {
		categories = append(categories, "nerd")
	}
	if *sportFlag {
		categories = append(categories, "sport")
	}
	if *runFlag {
		categories = append(categories, "running")
	}
	return
}

func TestNoFlags(t *testing.T) {
	help, cats := parseFlags([]string{})
	if help {
		t.Errorf("Help should be false by default")
	}
	if len(cats) != 0 {
		t.Errorf("Expected no categories, got %v", cats)
	}
}

func TestSingleCategoryFlag(t *testing.T) {
	_, cats := parseFlags([]string{"-g"})
	if len(cats) != 1 || cats[0] != "geek" {
		t.Errorf("Expected ['geek'], got %v", cats)
	}
}

func TestMultipleCategoryFlags(t *testing.T) {
	_, cats := parseFlags([]string{"-g", "-n", "-s"})
	expected := []string{"geek", "nerd", "sport"}
	if strings.Join(cats, ",") != strings.Join(expected, ",") {
		t.Errorf("Expected %v, got %v", expected, cats)
	}
}

func TestHelpFlag(t *testing.T) {
	help, _ := parseFlags([]string{"-h"})
	if !help {
		t.Errorf("Expected help to be true")
	}
}
