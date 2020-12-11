package cipher

import (
	"strings"
)

type Cipher interface {
	Encode(string) string
	Decode(string) string
}

func convertToPlainText(s string) string {
	var result string
	const alpha = "abcdefghijklmnopqrstuvwxyz"
  for _, char := range s {
		c := strings.ToLower(string(char))
		if strings.Contains(alpha, c) {
			result += c 
		}
	}
	return result
}
