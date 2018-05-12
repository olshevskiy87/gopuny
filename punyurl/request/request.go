package request

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/olshevskiy87/gopuny/punyurl/response"
	"github.com/olshevskiy87/gopuny/punyurl/result"
)

const endpoint = "http://services.sapo.pt/PunyURL"

// prepareURL returns trimmed and query escaped url
func prepareURL(inputURL string) string {
	return url.QueryEscape(strings.TrimSpace(inputURL))
}

// strategies names that could be used for queries
const (
	StrategyShort  = "GetCompressedURLByURL"
	StrategyExpand = "GetURLByCompressedURL"
)

var strategies = map[string]bool{
	StrategyShort:  true,
	StrategyExpand: true,
}

func closeResponseBody(body io.Closer) {
	err := body.Close()
	if err != nil {
		fmt.Println(err)
	}
}

// doHTTP sends a GET-query to sapo.pt service and returns an output body
func doHTTP(strategy string, inputURL string) (output []byte, err error) {
	if ok := strategies[strategy]; !ok {
		err = fmt.Errorf("invalid strategy: %s", strategy)
		return
	}
	var preparedURL = prepareURL(inputURL)
	if preparedURL == "" {
		err = fmt.Errorf("invalid input url: %s", inputURL)
		return
	}
	resp, err := http.Get(fmt.Sprintf("%s/%s?url=%s", endpoint, strategy, preparedURL))
	if err != nil {
		return
	}
	defer closeResponseBody(resp.Body)
	output, err = ioutil.ReadAll(resp.Body)
	return
}

// Do starts http-request and parse received output
func Do(strategy string, inputURL string) (*result.Result, error) {
	httpOutput, err := doHTTP(strategy, inputURL)
	if err != nil {
		return nil, err
	}
	result, err := response.Parse(httpOutput)
	if err != nil {
		return nil, err
	}
	return result, nil
}
