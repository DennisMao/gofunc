package http

import (
	"net/url"
)

import (
	"testing"
)

func TestUrlParse(t *testing.T) {
	testList := []string{
		"htt://127.0.0.1:42379",
		"htt//127.0.0.1:42379",
		"http://localhost:1001",
	}
	for _, tU := range testList {

		u, err := url.Parse(tU)
		if err != nil {
			t.Error(err)
			return
		}

		t.Logf(u.String())
	}

}
