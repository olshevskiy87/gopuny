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
	res, err := request.Do("GetCompressedURLByURL", p.URL)
	if err != nil {
		return nil, err
	}
	if res.ASCII == "" {
		return nil, fmt.Errorf("could not short url: %s", p.URL)
	}
	return res, nil
}

// Expand transforms a short url and returns a pointer to Result with long url
func (p *PunyURL) Expand() (*result.Result, error) {
	res, err := request.Do("GetURLByCompressedURL", p.URL)
	if err != nil {
		return nil, err
	}
	if res.URL == "" {
		return nil, fmt.Errorf("could not expand url: %s", p.URL)
	}
	return res, nil
}
