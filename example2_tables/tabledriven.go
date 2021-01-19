package example2_tables

import "strings"

/*
Split string s into all substrings separated by sep and
returns a slice of the substrings between those separators.
*/

func Split(s, sep string) []string {
	var result []string
	i := strings.Index(s, sep)
	for i > -1 {
		result = append(result, s[:i])
		s = s[i+len(sep):]
		i = strings.Index(s, sep)
	}
	return append(result, s)
}

func Concat(strings []string) string {
	result := ""

	for _, s := range strings {
		result = result + s
	}

	return result
}
