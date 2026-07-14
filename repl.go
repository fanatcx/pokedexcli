package main

import "strings"

// CleanInput lowercases text and splits it into words on whitespace.
// Empty or whitespace-only input yields an empty slice — not an error;
// the caller decides what to do with nothing.
func CleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}