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
