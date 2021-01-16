package movie

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func downloadCaptionsFromYoutuber(id string) (string, error) {
	resp, err := http.Get("https://www.youtube.com/get_video_info?video_id="+id)
	if err != nil { return "", err}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil { return "", err }
	u, err := url.ParseRequestURI("http://youtube.com/?" +string(b))
	if err != nil { return "", err }

	type CaptionTracks struct { BaseUrl string `json:"baseUrl"` }
	type PlayerCaptionsTracklistRenderer struct {
		CaptionTracks []CaptionTracks `json:"captionTracks"`
	}

	type Captions struct { PlayerCaptionsTracklistRenderer `json:"playerCaptionsTracklistRenderer"`}
	type PR struct { Captions `json:"captions"`	}

	playerResponse :=u.Query().Get("player_response")
	if playerResponse == "" { return "", errWrongJsonInput	}

	pr := &PR{}
	if err := json.Unmarshal([]byte(playerResponse), pr); err != nil { return "", err }
	if len(pr.CaptionTracks) == 0 {return "", errWrongJsonInput}
	baseUrl := pr.CaptionTracks[0].BaseUrl
	if baseUrl == "" {
		return "", errWrongJsonInput
	}

	fmt.Println("captionUrl:", baseUrl)
	respXml, err := http.Get(baseUrl)
	if err != nil { return "", err }

	defer respXml.Body.Close()
	bXml, _ := ioutil.ReadAll(respXml.Body)
	fmt.Println(string(bXml))
	return "none", nil
}

func (mv *UWMovie) DownloadCaptions() (string, error) {
	switch mv.serviceType {
	case YoutubeId:
		return downloadCaptionsFromYoutuber(mv.movieId)
	}
	return "", errUnsupportedService
}