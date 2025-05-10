package newsreader

import (
	"Gideon/internal/utils/httpproxy"
	"encoding/json"
	"fmt"
	"io"
)

type Article struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
	URLToImage  string `json:"urlToImage"`
	PublishedAt string `json:"publishedAt"`
}

type NewsResponse struct {
	Status       string    `json:"status"`
	TotalResults int       `json:"totalResults"`
	Articles     []Article `json:"articles"`
}

// Note: you might want to reset the following struct instances to use memory efficiently...
type NewsReader struct {
	ApiKey   string
	Response NewsResponse
	Cpage    int
}

func (N *NewsReader) FetchNews(subject string) error {
	url := fmt.Sprintf(newsapi_endpoint, // TODO: Add user preferences , that they can add their own params , add baseurl
		subject,
		N.ApiKey,
	)

	resp, err := httpproxy.GlobalHTTPClient.Get(url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &N.Response)
	if err != nil {
		return err
	}

	return nil
}
