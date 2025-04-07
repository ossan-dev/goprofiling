package logic

import "strings"

var savedNames []string

// GetSavedNames returns the savedNames variable
func GetSavedNames(separator string) string {
	return strings.Join(savedNames, separator)
}

// Append saves and returns the slice
func Append(name string) string {
	data := []byte(name)
	stringData := string(data)
	savedNames = append(savedNames, stringData)
	return strings.Join(savedNames, ",")
}
