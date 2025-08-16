package services

import (
	"errors"
	"main/database"
	"main/dtos"
	"main/helper"
	"main/models"
)

func CreateCategory(req dtos.CreateCategoryRequest) (*models.Category, error) {
	category := models.Category{
		CategoryID:   helper.GenerateID(),
		CategoryName: req.CategoryName,
	}

	if err := database.DB.Create(&category).Error; err != nil {
		return nil, errors.New("Failed to create cateogry: " + err.Error())
	}

	return &category, nil
}

func GetAllCategory() ([]models.Category, error) {
	var categories []models.Category

	if err := database.DB.Find(&categories).Error; err != nil {
		return nil, errors.New("Failed to retrieve all category: " + err.Error())
	}

	return categories, nil
}
