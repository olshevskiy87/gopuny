package result

import (
	"encoding/json"
	"fmt"
)

// Result keeps shorten URL parts and original URL
type Result struct {
	ASCII   string `xml:"ascii" json:"ascii"`
	Preview string `xml:"preview" json:"preview"`
	Puny    string `xml:"puny" json:"puny"`
	URL     string `xml:"url" json:"url"`
}

func (r *Result) String() string {
	return fmt.Sprintf(
		"%-10s%s\n%-10s%s\n%-10s%s\n%-10s%s",
		"Unicode", r.Puny,
		"ASCII", r.ASCII,
		"Preview", r.Preview,
		"Original", r.URL,
	)
}

// AsJSON returns Result struct encoded as json
func (r *Result) AsJSON() (string, error) {
	jsonText, err := json.Marshal(r)
	if err != nil {
		return "", err
	}
	return string(jsonText), nil
}
