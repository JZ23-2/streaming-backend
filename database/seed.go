package database

import (
	"errors"
	"main/helper"
	"main/models"
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
