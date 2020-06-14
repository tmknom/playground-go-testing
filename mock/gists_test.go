package mock

import (
	"io"
	"strings"
	"testing"
)

func TestRealListGists1(t *testing.T) {
	urls, err := ListGists1()
	if err != nil {
		t.Fatalf("list gists caused error: %s", err)
	}

	if expected := 7; len(urls) != expected {
		t.Fatalf("want %d, got %d (%s)", expected, len(urls), urls)
	}
}

func TestRealListGists2(t *testing.T) {
	c := &Client{&Gister{}}
	urls, err := c.ListGists2()
	if err != nil {
		t.Fatalf("list gists caused error: %s", err)
	}

	if expected := 7; len(urls) != expected {
		t.Fatalf("want %d, got %d (%s)", expected, len(urls), urls)
	}
}

func TestMockListGists1(t *testing.T) {
	doGistsRequest = func() (io.Reader, error) {
		return strings.NewReader(`
			[{"html_url":"https://dummy1"},{"html_url":"https://dummy2"}]
		`), nil
	}
	urls, err := ListGists1()
	if err != nil {
		t.Fatalf("list gists caused error: %s", err)
	}

	if expected := 2; len(urls) != expected {
		t.Fatalf("want %d, got %d (%s)", expected, len(urls), urls)
	}
}

type dummyDoer struct{}

func (g *dummyDoer) doGistsRequest() (io.Reader, error) {
	return strings.NewReader(`
			[{"html_url":"https://dummy1"},{"html_url":"https://dummy2"}]
		`), nil
}

func TestMockListGists2(t *testing.T) {
	c := &Client{&dummyDoer{}}
	urls, err := c.ListGists2()
	if err != nil {
		t.Fatalf("list gists caused error: %s", err)
	}

	if expected := 2; len(urls) != expected {
		t.Fatalf("want %d, got %d (%s)", expected, len(urls), urls)
	}
}
