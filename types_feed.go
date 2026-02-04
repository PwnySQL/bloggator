package main

import (
	"fmt"
	"html"
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

func (r *RSSFeed) Unescape() {
	r.Channel.Title = html.UnescapeString(r.Channel.Title)
	r.Channel.Description = html.UnescapeString(r.Channel.Description)
	for i, item := range r.Channel.Item {
		r.Channel.Item[i].Title = html.UnescapeString(item.Title)
		r.Channel.Item[i].Description = html.UnescapeString(item.Description)
	}
}

func (r *RSSFeed) Print() {
	fmt.Printf("Title: %s\nDescription: %s\nLink: %s\n", r.Channel.Title, r.Channel.Description, r.Channel.Link)
	for _, item := range r.Channel.Item {
		fmt.Printf("Title: %s\nDescription: %s\nLink: %s\nPubDate: %s\n", item.Title, item.Description, item.Link, item.PubDate)
	}
}
