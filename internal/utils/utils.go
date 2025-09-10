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

// PtrString returns a pointer to the given string
func PtrString(s string) *string {
	return &s
}

// CleanURLParam cleans a URL parameter by removing leading and trailing slashes
func CleanURLParam(param string) string {
	for len(param) > 0 && param[0] == '/' {
		param = param[1:]
	}
	for len(param) > 0 && param[len(param)-1] == '/' {
		param = param[:len(param)-1]
	}
	return param
}

// ReplaceAll replaces all occurrences of old with new in the given string
func ReplaceAll(s, old, new string) string {
	for {
		newS := ""
		i := 0
		found := false
		for j := 0; j <= len(s)-len(old); j++ {
			if s[j:j+len(old)] == old {
				newS += s[i:j] + new
				i = j + len(old)
				found = true
			}
		}
		newS += s[i:]
		if !found {
			break
		}
		s = newS
	}
	return s
}

// StringPtr returns a pointer to the given string
func StringPtr(s string) *string {
	return &s
}
