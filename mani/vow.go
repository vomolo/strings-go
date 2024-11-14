package mani

import "strings"

func Vow(vow string) string {
	split := strings.Fields(vow)
	for i := 0; i < len(split); i++ {
		if (split[i] == "a" || split[i] == "A") && (strings.Contains("aeiouhAEIOUH", string(split[i+1][0]))) {
			if split[i] == "a" {
				split[i] = "an"
			} else {
				split[i] = "An"
			}
		}
	}
	return strings.Join(split, " ")
}
