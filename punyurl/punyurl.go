package punyurl

import (
	"errors"
	"fmt"

	"github.com/olshevskiy87/gopuny/punyurl/request"
	"github.com/olshevskiy87/gopuny/punyurl/result"
)

// PunyURL is a struct with an initial URL
type PunyURL struct {
	URL string
}

// New returns new PunyURL object
func New(url string) (*PunyURL, error) {
	if url == "" {
		return nil, errors.New("URL is empty")
	}
	return &PunyURL{
		URL: url,
	}, nil
}

// Short transforms a long url and returns a pointer to Result with short url
func (p *PunyURL) Short() (*result.Result, error) {
	return p.applyStrategy(request.StrategyShort)
}

// Expand transforms a short url and returns a pointer to Result with long url
func (p *PunyURL) Expand() (*result.Result, error) {
	return p.applyStrategy(request.StrategyExpand)
}

func (p *PunyURL) applyStrategy(strategy string) (*result.Result, error) {
	res, err := request.Do(strategy, p.URL)
	if err != nil {
		return nil, err
	}
	if res.URL == "" {
		return nil, fmt.Errorf("retrieved url is empty")
	}
	return res, nil
}
