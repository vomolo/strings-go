package mani

import (
	"strconv"
	"strings"
)

func CapInt(str string) string {
	slice := strings.Fields(str)
	for i := 0; i < len(slice); i++ {
		if strings.Contains(slice[i], "(cap") {
			if strings.Contains(slice[i], "(cap,") {
				num, _ := strconv.Atoi(slice[i+1][:len(slice[i+1])-1])
				for j := i - num; j < i; j++ {
					slice[j] = strings.Title(slice[j])
				}
				slice = append(slice[:i], slice[i+2:]...)
			} else {
				slice[i-1] = strings.Title(slice[i-1])

				slice = append(slice[:i], slice[i+1:]...)
			}
		}
	}
	return strings.Join(slice, " ")
}
