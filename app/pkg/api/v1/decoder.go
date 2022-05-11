package v1

import (
	"encoding/json"
	"fmt"
)

// Decode parses any data to JSON.
func Decode(data []byte, v any) error {
	err := json.Unmarshal(data, v)
	if err != nil {
		return fmt.Errorf("DecodeOrderData: %w", err)
	}

	return nil
}
