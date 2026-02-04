package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
)

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	request, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		fmt.Errorf("Error creating request: %v", err)
		return nil, err
	}
	request.Header.Set("User-Agent", "bloggator")

	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Errorf("Error making request: %v", err)
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Errorf("Error reading reponse: %v", err)
		return nil, err
	}

	var rssFeed RSSFeed
	if err := xml.Unmarshal(body, &rssFeed); err != nil {
		fmt.Errorf("Error unmarshalling XML: %v", err)
		return nil, err
	}

	return &rssFeed, nil
}
