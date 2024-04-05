package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: <program> <inputFileName>")
		return
	}

	inputFileName := os.Args[1]

	file, err := os.Open(inputFileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	processedLines := processLines(lines)

	file, err = os.Create(inputFileName) // Overwrite the original file.
	if err != nil {
		fmt.Println("Error opening file for overwrite:", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, line := range processedLines {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
	}
	writer.Flush()
}

func processLines(lines []string) []string {
	var processed []string
	inEnum := false
	var enumValues []string
	var enumPrefixes []string // Keep track of the original prefixes (indentation + "-") for each enum value

	for _, line := range lines {
		if strings.Contains(line, "enum:") {
			inEnum = true
			enumValues = nil
			enumPrefixes = nil
			processed = append(processed, line)
			continue
		}

		if inEnum {
			if line == "" || !strings.Contains(line, "- ") {
				inEnum = false
				prefixToRemove := findLongestCommonPrefix(enumValues)
				for i, val := range enumValues {
					updatedContent := strings.Replace(val, prefixToRemove, "", 1)
					lowCase := strings.ToLower(updatedContent)
					processed = append(processed, enumPrefixes[i]+lowCase)
				}
				if line != "" {
					processed = append(processed, line)
				}
				continue
			}

			prefix, value := extractPrefixAndValue(line)
			enumPrefixes = append(enumPrefixes, prefix)
			enumValues = append(enumValues, value)
		} else {
			processed = append(processed, line)
		}
	}
	return processed
}

func findLongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	prefix := strs[0]
	for _, s := range strs {
		for strings.Index(s, prefix) != 0 {
			if len(prefix) == 0 {
				return ""
			}
			prefix = prefix[:len(prefix)-1]
		}
	}
	if !strings.Contains(prefix, "_") {
		return ""
	}
	return prefix
}

// Extracts the prefix (indentation + "-") and the actual value part of an enum line.
func extractPrefixAndValue(line string) (string, string) {
	trimmedLine := strings.TrimLeft(line, " ")
	indentation := line[:len(line)-len(trimmedLine)]
	// Extract the "- " part along with the indentation
	if strings.HasPrefix(trimmedLine, "-") {
		indentation += "-"
		trimmedLine = strings.TrimPrefix(trimmedLine, "-")
	}
	value := strings.TrimPrefix(trimmedLine, " ")
	return indentation + " ", value // Ensure to add a space after "-" to maintain the original format
}
