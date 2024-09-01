package scraper

import (
	"encoding/xml"
	"errors"
	"net/http"
)

type RssFeed struct {
	Channel struct {
		Title string `xml:"title"`
		Link  string `xml:"link"`
		Items []struct {
			Title       string `xml:"title"`
			Link        string `xml:"link"`
			Description string `xml:"description"`
		} `xml:"item"`
	} `xml:"channel"`
}

func ScrapeFeed(rawUrl string) (RssFeed, error) {
	resp, err := http.Get(rawUrl)
	if err != nil {
		return RssFeed{}, errors.New("Unable to reach URL")
	}
	feed := RssFeed{}
	err = xml.NewDecoder(resp.Body).Decode(&feed)
	if err != nil {
		return RssFeed{}, errors.New("Unable to parse XML response")
	}
	return feed, nil
}
