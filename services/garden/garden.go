package services

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type Garden struct{}

func (g *Garden) Search(place string) Search {

	res, err := http.Get("https://radio.garden/api/search?q=" + url.QueryEscape(place))

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	places := Search{}
	json.NewDecoder(res.Body).Decode(&places)

	return places
}

func (g *Garden) Stream(id string) io.ReadCloser {
	target := strings.Split(id, "/")

	res, err := http.Get("https://radio.garden/api/ara/content/listen/" + url.QueryEscape(target[len(target)-1]) + "/channel.mp3")

	if err != nil {
		panic(err)
	}

	return res.Body
}

type API struct {
	Search Search
}

type Search struct {
	Took int `json:"took"`
	Hits struct {
		Hits []struct {
			ID     string  `json:"_id"`
			Score  float64 `json:"_score"`
			Source struct {
				Code     string `json:"code"`
				Subtitle string `json:"subtitle"`
				Type     string `json:"type"`
				Title    string `json:"title"`
				URL      string `json:"url"`
			} `json:"_source"`
		} `json:"hits"`
	} `json:"hits"`
	Query      string `json:"query"`
	Version    string `json:"version"`
	APIVersion int    `json:"apiVersion"`
}
