package movie

import (
	"net/url"
	"strings"
)

func ParseLink(link string) (*UWMovie, error) {
	u, err := url.Parse(link)
	if err != nil {
		return nil, err
	}

	if isYoutube(u.Host) {
		return &UWMovie{
			serviceType: YoutubeId,
			movieId:     u.Query().Get("v"),
		}, nil
	}
	return nil, errUnsupportedService
}

func isYoutube(host string) bool {
	if strings.HasSuffix(host, "youtube.com") {
		return true
	}
	return false
}
