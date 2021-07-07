package data

import (
	"fmt"

	"github.com/guilhermecoelho/kakeibo/configurations"
	"github.com/guilhermecoelho/kakeibo/models"
)

func GetGroups() (models.Groups, error) {

	group := models.Groups{}
	result := configurations.DBgorm.Find(&group)

	if result.Error != nil {
		return group, fmt.Errorf("error on: %v", result.Error)
	}

	return group, nil
}

func GetGroupById(id int) (models.Group, error) {

	group := models.Group{}
	result := configurations.DBgorm.First(&group, id)

	if result.Error != nil {
		return group, fmt.Errorf("error on: %v", result.Error)
	}

	return group, nil
}

func PutGroup(g *models.Group) (models.Group, error) {

	group := *g
	result := configurations.DBgorm.Save(&group)

	if result.Error != nil {
		return group, fmt.Errorf("error on: %v", result.Error)
	}

	group, err := GetGroupById(group.Id)
	if err != nil {
		return group, fmt.Errorf("error on: %v", result.Error)
	}

	return group, nil

}

func PostGroup(g *models.Group) (models.Group, error) {

	group := *g
	group.Id = 0
	result := configurations.DBgorm.Create(&group)

	if result.Error != nil {
		return group, fmt.Errorf("error on: %v", result.Error)
	}

	group, err := GetGroupById(group.Id)
	if err != nil {
		return group, fmt.Errorf("error on: %v", result.Error)
	}

	return group, nil
}

func DeleteGroup(g *models.Group) error {

	group := *g
	result := configurations.DBgorm.Delete(&group)

	if result.Error != nil {
		return fmt.Errorf("error on: %v", result.Error)
	}

	return nil
}
