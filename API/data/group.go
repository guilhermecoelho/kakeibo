package data

import (
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
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

func GetGroupById(id string) (models.Group, error) {

	group := models.Group{}
	// result := configurations.DBgorm.First(&group, id)

	// if result.Error != nil {
	// 	return group, fmt.Errorf("error on: %v", result.Error)
	// }
	result, err := configurations.DBRedis.Get("group_" + id).Result()
	if err != nil {
		return group, fmt.Errorf("error on: %v", err)
	}

	error := json.Unmarshal([]byte(result), &group)
	if err != nil {
		return group, fmt.Errorf("error on: %v", error)
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
	group.Id = uuid.NewString()

	//result := configurations.DBgorm.Create(&group)

	// if result.Error != nil {
	// 	return group, fmt.Errorf("error on: %v", result.Error)
	// }

	// group, err := GetGroupById(group.Id)
	// if err != nil {
	// 	return group, fmt.Errorf("error on: %v", result.Error)
	// }
	transform, errorTrans := json.Marshal(group)
	if errorTrans != nil {
		return group, fmt.Errorf("error on: %v", errorTrans)
	}

	err := configurations.DBRedis.Set("group_"+group.Id, string(transform), 0).Err()
	if err != nil {
		return group, fmt.Errorf("error on: %v", err)
	}

	group, error := GetGroupById(group.Id)
	if err != nil {
		return group, fmt.Errorf("error on: %v", error)
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
