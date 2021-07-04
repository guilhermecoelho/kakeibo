package data

import (
	"fmt"

	"github.com/guilhermecoelho/kakeibo/configurations"
	"github.com/guilhermecoelho/kakeibo/models"
)

func GetCategories() (models.Categories, error) {

	group := models.Categories{}
	result := configurations.DBgorm.Find(&group)

	if result.Error != nil {
		return group, fmt.Errorf("error on: %v", result.Error)
	}

	return group, nil
}

func GetCategoryById(id int) (models.Category, error) {

	group := models.Category{}
	result := configurations.DBgorm.First(&group, id)

	if result.Error != nil {
		return group, fmt.Errorf("error on: %v", result.Error)
	}

	return group, nil
}

func PutCategory(g *models.Category) (models.Category, error) {

	group := *g
	result := configurations.DBgorm.Save(&group)

	if result.Error != nil {
		return group, fmt.Errorf("error on: %v", result.Error)
	}

	group, err := GetCategoryById(group.Id)
	if err != nil {
		return group, fmt.Errorf("error on: %v", result.Error)
	}

	return group, nil

}

func PostCategory(g *models.Category) (models.Category, error) {

	group := *g
	result := configurations.DBgorm.Create(&group)

	if result.Error != nil {
		return group, fmt.Errorf("error on: %v", result.Error)
	}

	group, err := GetCategoryById(group.Id)
	if err != nil {
		return group, fmt.Errorf("error on: %v", result.Error)
	}

	return group, nil
}

func DeleteCategory(g *models.Category) error {

	group := *g
	result := configurations.DBgorm.Delete(&group)

	if result.Error != nil {
		return fmt.Errorf("error on: %v", result.Error)
	}

	return nil
}
