package main

import (
	"context"
	"encoding/xml"
	"html"
	"io"
	"net/http"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return &RSSFeed{}, err
	}
	req.Header.Add("User-Agent", "gator")

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return &RSSFeed{}, err
	}

	defer response.Body.Close()

	data, err := io.ReadAll(response.Body)
	if err != nil {
		return &RSSFeed{}, err
	}

	feedResponse := RSSFeed{}
	err = xml.Unmarshal(data, &feedResponse)
	if err != nil {

		return &RSSFeed{}, err
	}
	feedResponse.Channel.Title = html.UnescapeString(feedResponse.Channel.Title)
	feedResponse.Channel.Description = html.UnescapeString(feedResponse.Channel.Description)

	return &feedResponse, nil
}
