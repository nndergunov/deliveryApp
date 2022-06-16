package accountingapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func BindJson(req *http.Request, obj interface{}) error {
	if req == nil || req.Body == nil {
		return fmt.Errorf("invalid request")
	}
	return DecodeJSON(req.Body, obj)
}

func DecodeJSON(r io.Reader, obj interface{}) error {
	decoder := json.NewDecoder(r)

	if err := decoder.Decode(obj); err != nil {
		return err
	}
	return nil
}
