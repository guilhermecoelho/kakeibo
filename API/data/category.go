package data

import (
	"fmt"

	"github.com/guilhermecoelho/kakeibo/configurations"
	"github.com/guilhermecoelho/kakeibo/models"
)

func GetCategories() (models.Categories, error) {

	category := models.Categories{}

	result := configurations.DBgorm.Find(&category)

	if result.Error != nil {
		return category, fmt.Errorf("error on: %v", result.Error)
	}

	return category, nil
}

func GetCategoryById(id int) (models.Category, error) {

	category := models.Category{}
	result := configurations.DBgorm.First(&category, id)

	if result.Error != nil {
		return category, fmt.Errorf("error on: %v", result.Error)
	}

	return category, nil
}

func PutCategory(g *models.Category) (models.Category, error) {

	category := *g
	result := configurations.DBgorm.Save(&category)

	if result.Error != nil {
		return category, fmt.Errorf("error on: %v", result.Error)
	}

	category, err := GetCategoryById(category.Id)
	if err != nil {
		return category, fmt.Errorf("error on: %v", result.Error)
	}

	return category, nil

}

func PostCategory(g *models.Category) (models.Category, error) {

	category := *g
	category.Id = 0
	result := configurations.DBgorm.Create(&category)

	if result.Error != nil {
		return category, fmt.Errorf("error on: %v", result.Error)
	}

	category, err := GetCategoryById(category.Id)
	if err != nil {
		return category, fmt.Errorf("error on: %v", result.Error)
	}

	return category, nil
}

func DeleteCategory(g *models.Category) error {

	category := *g
	result := configurations.DBgorm.Delete(&category)

	if result.Error != nil {
		return fmt.Errorf("error on: %v", result.Error)
	}

	return nil
}
