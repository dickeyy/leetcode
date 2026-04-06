package main

import (
	"strings"
	"unicode"
)

func decodeCiphertext(encodedText string, rows int) string {
	if rows == 1 {
		return encodedText
	}

	n := len(encodedText) / rows
	var sb strings.Builder
	sb.Grow(len(encodedText))

	for startCol := range n {
		for i, j := 0, startCol; i < rows && j < n; i, j = i+1, j+1 {
			idx := (i * n) + j
			sb.WriteByte(encodedText[idx])
		}
	}

	return strings.TrimRightFunc(sb.String(), unicode.IsSpace)
}
