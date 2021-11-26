package validation

import (
	"encoding/json"
	"fmt"
)

func SearchQuery(searchQuery string) error {
	if !json.Valid([]byte(searchQuery)) {
		return fmt.Errorf("Invalid search query, must be json")
	}
	return nil
}
