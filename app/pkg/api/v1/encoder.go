package v1

import (
	"encoding/json"
	"fmt"
)

// Encode parses any data to JSON.
func Encode(data any) ([]byte, error) {
	encData, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("json encoding fail: %w", err)
	}

	return encData, nil
}

// EncodeIndent parses any data to JSON with specified indentation.
func EncodeIndent(data any, prefix, indent string) ([]byte, error) {
	encData, err := json.MarshalIndent(data, prefix, indent)
	if err != nil {
		return nil, fmt.Errorf("json encoding fail: %w", err)
	}

	return encData, nil
}
