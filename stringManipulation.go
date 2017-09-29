package main

import (
	"strings"
)

func getLines(text string) []string {
	lines := strings.Split(text, "\n")
	return lines
}

func locateStringInArray(lines []string, s string) []int {
	var positions []int

	for k, l := range lines {
		if strings.Contains(l, s) {
			positions = append(positions, k)
		}
	}

	return positions
}

func deleteArrayElementsWithString(lines []string, s string) []string {
	var result []string
	for _, l := range lines {
		if !strings.Contains(l, s) {
			result = append(result, l)
		}
	}
	return result
}

func deleteLinesBetween(lines []string, from int, to int) []string {
	var result []string
	result = append(lines[:from], lines[to+1:]...)
	return result
}

func addElementsToArrayPosition(lines []string, newLines []string, pos int) []string {
	var result []string
	result = append(result, lines[:pos]...)
	result = append(result, newLines...)
	result = append(result, lines[pos:]...)
	/*
		result = append(lines[:pos], newLines...)
		result = append(result, lines[pos:]...)
	*/
	return result
}

func concatStringsWithJumps(lines []string) string {
	var r string
	for _, l := range lines {
		r = r + l + "\n"
	}
	return r
}
