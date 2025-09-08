package utils

import (
	"encoding/json"
	"io"
)

// DecodeJSON decodes JSON from the request body
func DecodeJSON(body io.ReadCloser, out interface{}) error {
	defer body.Close()
	decoder := json.NewDecoder(body)
	return decoder.Decode(out)
}

// JoinStrings	 joins a slice of strings with a given separator
func JoinStrings(elements []string, sep string) string {
	result := ""
	for i, elem := range elements {
		if i > 0 {
			result += sep
		}
		result += elem
	}
	return result
}
