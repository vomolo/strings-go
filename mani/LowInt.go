package mani

import (
	"strconv"
	"strings"
)

func LowInt(str string) string {
	slice := strings.Fields(str)
	for i := 0; i < len(slice); i++ {
		if strings.Contains(slice[i], "(low") {
			if strings.Contains(slice[i], "(low,") {
				num, _ := strconv.Atoi(slice[i+1][:len(slice[i+1])-1])
				for k := i - num; k < i; k++ {
					slice[k] = strings.ToLower(slice[k])
				}
				slice = append(slice[:i], slice[i+2:]...)

			} else {
				slice[i-1] = strings.ToLower(slice[i-1])
				slice = append(slice[:i], slice[i+1:]...)
			}
		}
	}
	return strings.Join(slice, " ")
}
