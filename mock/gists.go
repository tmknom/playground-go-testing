package mock

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

type Gist struct {
	Rawurl string `json:"html_url"`
}

var doGistsRequest = func() (io.Reader, error) {
	resp, err := http.Get("https://api.github.com/users/tmknom/gists")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var buf bytes.Buffer
	if _, err := io.Copy(&buf, resp.Body); err != nil {
		return nil, err
	}
	return &buf, nil
}

func ListGists1() ([]string, error) {
	r, err := doGistsRequest()
	if err != nil {
		return nil, err
	}
	var gists []Gist
	if err := json.NewDecoder(r).Decode(&gists); err != nil {
		return nil, err
	}

	urls := make([]string, 0, len(gists))
	for _, u := range gists {
		urls = append(urls, u.Rawurl)
	}
	return urls, nil
}

type Doer interface {
	doGistsRequest() (io.Reader, error)
}

type Client struct {
	Gister Doer
}

type Gister struct{}

func (g *Gister) doGistsRequest() (io.Reader, error) {
	return doGistsRequest()
}

func (c *Client) ListGists2() ([]string, error) {
	r, err := c.Gister.doGistsRequest()
	if err != nil {
		return nil, err
	}
	var gists []Gist
	if err := json.NewDecoder(r).Decode(&gists); err != nil {
		return nil, err
	}

	urls := make([]string, 0, len(gists))
	for _, u := range gists {
		urls = append(urls, u.Rawurl)
	}
	return urls, nil
}
