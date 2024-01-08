import (
	"encoding/json"
	"fmt"
)
// there's empty fillings
func recursiveFilterCases(cases json.RawMessage, result *[]testCase, excludedTests map[string]struct{}) error {
	var js []map[string]json.RawMessage

	// 'cases' is always an array, where every item is either a single test or an object containing nested cases
	err := json.Unmarshal(cases, &js)
	if err != nil {
		return err
	}
