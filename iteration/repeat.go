package iteration

import "strings"

func Repeat(c string, i int) string {
	var repeated strings.Builder
	for range i {
		repeated.WriteString(c)
	}
	return repeated.String()
}
