package movie

import "errors"

var (
	errUnsupportedService = errors.New("unsupported service")
)

const Youtube = "youtube.com"

const (
	YoutubeId = iota
)

type UWMovie struct {
	serviceType int
	movieId string
}
