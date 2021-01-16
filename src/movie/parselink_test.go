package movie

import (
	"strings"
	"testing"
)

func TestParseLink(t *testing.T) {
	tests := []struct{
		link	string
		err		string
		service	int
		id		string
	}{
		{ "ht tp://wrongurl", "cannot contain colon", 0, "" },
		{ "", errUnsupportedService.Error(), 0, "" },
		{
			"https://www.youtube.com/watch?v=jTSvthW34GU&t=320s",
			"",
			YoutubeId,
			"jTSvthW34GU",
		},
		{
			"https://MyUnsupportedService.com/watch?v=jTSvthW34GU&t=320s",
			errUnsupportedService.Error(),0, "",
		},
	}

	for _, test := range tests {
		got, err := ParseLink(test.link)
		if err != nil {
			if test.err == "" || !strings.Contains(err.Error(), test.err) {
				t.Errorf("failed error for %q: wanted %q, got %q", test.link, test.err, err)
			}
			continue
		}
		if got.serviceType != test.service {
			t.Errorf("a service type of ParseLink(%q) == %d, wanted %v", test.link, test.service, got.serviceType)
		}
		if got.movieId != test.id {
			t.Errorf("a movie id of ParseLink(%q) == %q, wanted %q", test.link, test.id, got.movieId)
		}
	}
}

func TestIsYoutube(t *testing.T) {
	tests := []struct{
		host string
		result bool
	}{
		{"www.youtube.com", true},
		{"youtube.com", true},
		{"google.com", false},
	}
	for _, test := range tests {
		got := isYoutube(test.host)
		if got != test.result {
			t.Errorf("isYoutube(%q) == %t, wanted %t",
				test.host, got, test.result)
		}
	}
}