package movie

import "errors"

var (
	errUnsupportedService = errors.New("unsupported service")
	errFailedHttpStatus   = errors.New("received a failed status for http request")
	errWrongJsonInput     = errors.New("wrong JSON input")
)

const Youtube = "youtube.com"

const (
	YoutubeId = iota
)

type UWMovie struct {
	serviceType int
	movieId     string
}
