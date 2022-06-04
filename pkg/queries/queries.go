package queries

import (
	"encoding/json"
	"net/http"
)

// Get performs a GET request to the specified URL and unmarshals the response into the specified interface.
func Get(url string, target any) error {
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	json.NewDecoder(res.Body).Decode(target)
	return nil
}
