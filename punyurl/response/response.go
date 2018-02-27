package response

import (
	"encoding/xml"

	"github.com/olshevskiy87/gopuny/punyurl/result"
)

// Parse returns Result structure from xml content
func Parse(content []byte) (*result.Result, error) {
	result := result.Result{}
	err := xml.Unmarshal(content, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
