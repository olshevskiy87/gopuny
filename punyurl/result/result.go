package result

import "fmt"

// Result keeps shorten URL parts and original URL
type Result struct {
	ASCII   string `xml:"ascii"`
	Preview string `xml:"preview"`
	Puny    string `xml:"puny"`
	URL     string `xml:"url"`
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
