package utils

import (
	"fmt"
)

func StringArrayToString(arr []string) string {
	output := ""
	for _, item := range arr {
		output += fmt.Sprintf("%s\n", item)
	}
	return output
}
