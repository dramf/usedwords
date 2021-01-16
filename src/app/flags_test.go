package main

import (
	"strings"
	"testing"
)

const (
	correctMovieLink = "https://www.youtube.com/watch?v=jTSvthW34GU"
)

func TestParseFlags(t *testing.T) {
	tests := []struct{
		args []string
		movie string
	}{
		{
			[]string{"-movie", correctMovieLink},
			correctMovieLink,
		},
	}

	for _, test := range tests {
		t.Run(strings.Join(test.args, " "), func(t *testing.T) {
			conf, output, err := parseFlags("prog", test.args)
			if err != nil {
				t.Errorf("error got %v, want nil", err)
			}
			if output != "" {
				t.Errorf("output got %q, want empty", output)
			}
			if conf.movie != test.movie {
				t.Errorf("movie got %q, want %q", conf.movie, test.movie)
			}
		})
	}
}