// Package concatenator provides string concatenation features.
package concatenator

import "bytes"

// Strings - It's the unoptimized version.
func Strings(first, second string) string {
	var buffer bytes.Buffer
	buffer.WriteString(first)
	buffer.WriteString(second)
	return buffer.String()
}

// Strings - It's the optimized version.
// func Strings(first, second string) string {
// 	return strings.Join([]string{first, second}, "")
// }
