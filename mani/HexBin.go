package mani

import (
	"fmt"
	"regexp"
	"strconv"
)

func ReplaceHexAndBinNumbers(line string) string {
	reHex := regexp.MustCompile(`(?i)(\b\w+\b)\s?\(hex\)`)
	matchesHex := reHex.FindStringSubmatch(line)
	hexNumber := "0"
	if len(matchesHex) > 1 {
		hexNumber = matchesHex[1]
	}

	decimalHex, err := strconv.ParseInt(hexNumber, 16, 64)
	if err != nil {
		// If the hex number cannot be parsed, it's already in decimal format, so leave it as is
		return line
	}

	reBin := regexp.MustCompile(`(?i)(\b[0-1]+\b)\s?\(bin\)`)
	matchesBin := reBin.FindStringSubmatch(line)
	binNumber := "0"
	if len(matchesBin) > 1 {
		binNumber = matchesBin[1]
	}

	decimalBin, err := strconv.ParseInt(binNumber, 2, 64)
	if err != nil {
		// If the bin number cannot be parsed, it's already in decimal format, so leave it as is
		return line
	}

	// Replace the hexadecimal and binary numbers with their decimal equivalents
	replacedLine := reHex.ReplaceAllString(line, fmt.Sprintf("%d", decimalHex))
	replacedLine = reBin.ReplaceAllString(replacedLine, fmt.Sprintf("%d", decimalBin))

	return replacedLine
}
