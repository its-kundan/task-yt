package youtube;

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
	"youtube-fetcher/config"
	"youtube-fetcher/db"

	"gorm.io/gorm/clause" // Import the clause package explicitly
)

type YouTubeResponse struct {
	Items []struct {
		ID struct {
			VideoID string `json:"videoId"`
		} `json:"id"`
		Snippet struct {
			Title       string    `json:"title"`
			Description string    `json:"description"`
			PublishedAt time.Time `json:"publishedAt"`
			Thumbnails  struct {
				Default struct {
					URL string `json:"url"`
				} `json:"default"`
			} `json:"thumbnails"`
		} `json:"snippet"`
	} `json:"items"`
}

var apiIndex = 0

func fetchVideos(query string) {
	apiKey := config.AppConfig.APIKeys[apiIndex]
	url := fmt.Sprintf("https://www.googleapis.com/youtube/v3/search?part=snippet&type=video&order=date&q=%s&key=%s", query, apiKey)

	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Error fetching YouTube API: %v", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == 403 { // API quota exhausted
		apiIndex = (apiIndex + 1) % len(config.AppConfig.APIKeys)
		log.Println("Switching API key due to quota exhaustion")
		return
	}

	var data YouTubeResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		log.Printf("Error decoding response: %v", err)
		return
	}

	for _, item := range data.Items {
		video := db.Video{
			VideoID:       item.ID.VideoID,
			Title:         item.Snippet.Title,
			Description:   item.Snippet.Description,
			PublishedAt:   item.Snippet.PublishedAt,
			ThumbnailsURL: item.Snippet.Thumbnails.Default.URL,
		}

		// Use clause.OnConflict to avoid duplicate entries
		db.DB.Clauses(clause.OnConflict{
			DoNothing: true, // Ignore if conflict occurs
		}).Create(&video)
	}
	log.Println("Fetched and stored latest videos")
}

func StartFetcher(query string) {
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		fetchVideos(query)
	}
}
