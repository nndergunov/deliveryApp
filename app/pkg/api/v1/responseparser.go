package v1

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// DecodeResponse gets data from the response using default decoder.
func DecodeResponse(resp *http.Response, data any) error {
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("getting response body: %w", err)
	}

	err = resp.Body.Close()
	if err != nil {
		return fmt.Errorf("closing response body: %w", err)
	}

	err = Decode(respBody, data)
	if err != nil {
		return fmt.Errorf("decoding response: %w", err)
	}

	return nil
}
