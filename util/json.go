package util

import (
	"encoding/json"
	"net/http"
)

func DecodeJSON(resp *http.Response, target interface{}) error {
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	err := decoder.Decode(target)

	return err
}
