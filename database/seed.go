package database

import (
	"errors"
	"fmt"
	"log"
	"main/helper"
	"main/models"
	"math/rand"
	"time"
)

func SeedStreamingCategories() error {
	categories := []string{
		"Music",
		"Gaming",
		"Sports",
		"Movies",
		"News",
		"Education",
		"Podcasts",
		"Technology",
		"Travel",
		"Food & Cooking",
	}

	for _, name := range categories {
		category := models.Category{
			CategoryID:   helper.GenerateID(),
			CategoryName: name,
		}

		if err := DB.Where("category_name = ?", name).FirstOrCreate(&category).Error; err != nil {
			return errors.New("failed to seed category " + name + ": " + err.Error())
		}
	}

	return nil
}

func SeedStreamInfos(n int) error {
	var categories []models.Category
	if err := DB.Find(&categories).Error; err != nil {
		return err
	}
	if len(categories) == 0 {
		return fmt.Errorf("no categories found, seed categories first")
	}

	titles := []string{
		"Chill Beats Live",
		"Pro Gaming Marathon",
		"Daily News Recap",
		"Cooking with Love",
		"Travel the World",
		"Tech Talk Live",
		"Sports Highlights",
		"Movie Night Watch Party",
		"Educational Insights",
		"Podcast Special",
	}

	for i := 0; i < n; i++ {
		hostID := fmt.Sprintf("host_info_%03d", i+1)

		title := titles[rand.Intn(len(titles))]
		category := categories[rand.Intn(len(categories))]

		streamInfo := models.StreamInfo{
			HostPrincipalID:  hostID,
			Title:            title,
			StreamCategoryID: &category.CategoryID,
		}

		if err := DB.Create(&streamInfo).Error; err != nil {
			return err
		}
	}

	log.Printf("Seeded %d stream infos successfully\n", n)
	return nil
}

func SeedStreams(n int) error {
	var streamInfos []models.StreamInfo
	if err := DB.Find(&streamInfos).Error; err != nil {
		return err
	}
	if len(streamInfos) == 0 {
		return fmt.Errorf("no stream infos found, seed stream infos first")
	}

	thumbnails := []string{
		"https://placehold.co/600x400?text=Music+Stream",
		"https://placehold.co/600x400?text=Gaming+Stream",
		"https://placehold.co/600x400?text=Sports+Stream",
		"https://placehold.co/600x400?text=Movies+Stream",
		"https://placehold.co/600x400?text=News+Stream",
	}

	for i := 0; i < n; i++ {
		streamInfo := streamInfos[rand.Intn(len(streamInfos))]

		stream := models.Stream{
			StreamID:        helper.GenerateID(),
			HostPrincipalID: streamInfo.HostPrincipalID,
			StreamInfoID:    &streamInfo.HostPrincipalID,
			CreatedAt:       time.Now().Add(-time.Duration(rand.Intn(5000)) * time.Minute),
			ThumbnailURL:    thumbnails[rand.Intn(len(thumbnails))],
			IsActive:        false,
		}

		if err := DB.Create(&stream).Error; err != nil {
			return err
		}
	}

	log.Printf("Seeded %d streams successfully\n", n)
	return nil
}

func SeedStreamHistories(n int) error {
	var streams []models.Stream
	if err := DB.Preload("StreamInfo").Find(&streams).Error; err != nil {
		return err
	}
	if len(streams) == 0 {
		return fmt.Errorf("no streams found, seed streams first")
	}

	videoUrls := []string{
		"https://example.com/videos/stream1.mp4",
		"https://example.com/videos/stream2.mp4",
		"https://example.com/videos/stream3.mp4",
		"https://example.com/videos/stream4.mp4",
		"https://example.com/videos/stream5.mp4",
	}

	for i := 0; i < n; i++ {
		stream := streams[rand.Intn(len(streams))]

		duration := rand.Intn(7200) + 300

		streamHistory := models.StreamHistory{
			StreamHistoryID:       helper.GenerateID(),
			StreamHistoryStreamID: stream.StreamID,
			HostPrincipalID:       stream.HostPrincipalID,
			VideoUrl:              videoUrls[rand.Intn(len(videoUrls))],
			Duration:              duration,
			CreatedAt:             time.Now().Add(-time.Duration(rand.Intn(5000)) * time.Minute),
		}

		if stream.StreamInfoID != nil {
			streamHistory.Title = stream.StreamInfo.Title
			if stream.StreamInfo.StreamCategoryID != nil {
				streamHistory.StreamCategoryID = stream.StreamInfo.StreamCategoryID
			}
		}

		if err := DB.Create(&streamHistory).Error; err != nil {
			return err
		}
	}

	log.Printf("Seeded %d stream histories successfully\n", n)
	return nil
}

func SeedViewerHistories() error {
	var streamHistories []models.StreamHistory
	if err := DB.Find(&streamHistories).Error; err != nil {
		return err
	}
	if len(streamHistories) == 0 {
		return fmt.Errorf("no stream histories found, seed stream histories first")
	}

	for _, history := range streamHistories {
		viewerCount := rand.Intn(30) + 1

		for i := 0; i < viewerCount; i++ {
			viewerID := fmt.Sprintf("viewer_%03d", rand.Intn(1000))

			viewerHistory := models.ViewerHistory{
				ViewerHistoryID:              helper.GenerateID(),
				ViewerHistoryStreamHistoryID: history.StreamHistoryID,
				ViewerHistoryPrincipalID:     viewerID,
			}

			if err := DB.Create(&viewerHistory).Error; err != nil {
				return err
			}
		}
	}

	log.Printf("Seeded viewer histories for %d stream histories\n", len(streamHistories))
	return nil
}

func SeedMessages() error {
	var streams []models.Stream
	if err := DB.Find(&streams).Error; err != nil {
		return err
	}
	if len(streams) == 0 {
		return fmt.Errorf("no streams found, seed streams first")
	}

	messages := []string{
		"This stream is awesome",
		"Great content",
		"Really enjoying this",
		"Love the vibe here",
		"Keep it up",
		"Nice work",
		"Very informative, thanks",
		"This is relaxing",
		"Good game",
		"Awesome discussion",
		"Great music",
		"Interesting topic",
		"Really helpful, thanks",
		"Love this part",
		"Cool stream",
		"Very entertaining",
		"Good job",
		"Nice explanation",
		"Thanks for sharing",
		"Super fun",
	}

	for _, stream := range streams {
		chatCount := rand.Intn(30) + 1
		for i := 0; i < chatCount; i++ {
			viewerID := fmt.Sprintf("viewer_%03d", rand.Intn(1000))
			content := messages[rand.Intn(len(messages))]

			msg := models.Message{
				MessageID:          helper.GenerateID(),
				StreamID:           stream.StreamID,
				MessagePrincipalID: viewerID,
				Content:            content,
				CreatedAt:          time.Now().Add(-time.Duration(rand.Intn(5000)) * time.Second),
			}

			if err := DB.Create(&msg).Error; err != nil {
				return err
			}
		}
	}

	log.Printf("Seeded messages for %d streams\n", len(streams))
	return nil
}
