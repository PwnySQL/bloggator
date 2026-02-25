package main

import (
	"fmt"
	"html"
	"strings"
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

func (r *RSSFeed) String() string {
	var sb strings.Builder
	sb.WriteString("##############################################################################################\n")
	sb.WriteString(fmt.Sprintf("Title: %s\nDescription: %s\nLink: %s\n", r.Channel.Title, r.Channel.Description, r.Channel.Link))
	for _, item := range r.Channel.Item {
		sb.WriteString(item.String())
	}
	sb.WriteString("##############################################################################################\n")
	return sb.String()
}

func (i *RSSItem) String() string {
	return fmt.Sprintf("\nTitle: %s - %s\nDescription: %s\nLink: %s\n", i.Title, i.PubDate, i.Description, i.Link)
}
