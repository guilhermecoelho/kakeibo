package data

import (
	"fmt"

	"github.com/guilhermecoelho/kakeibo/configurations"
	"github.com/guilhermecoelho/kakeibo/models"
)

func GetUsers() (models.Users, error) {

	user := models.Users{}
	result := configurations.DBgorm.Find(&user)

	if result.Error != nil {
		return user, fmt.Errorf("error on: %v", result.Error)
	}

	return user, nil
}

func GetUserById(id int) (models.User, error) {

	user := models.User{}
	result := configurations.DBgorm.First(&user, id)

	if result.Error != nil {
		return user, fmt.Errorf("error on: %v", result.Error)
	}

	return user, nil
}

func GetUserByName(name string) (models.User, error) {

	user := models.User{}
	result := configurations.DBgorm.Where("name = ?", name).First(&user)

	if result.Error != nil {
		return user, fmt.Errorf("error on: %v", result.Error)
	}

	return user, nil
}

func PutUser(g *models.User) (models.User, error) {

	user := *g
	result := configurations.DBgorm.Save(&user)

	if result.Error != nil {
		return user, fmt.Errorf("error on: %v", result.Error)
	}

	user, err := GetUserById(user.Id)
	if err != nil {
		return user, fmt.Errorf("error on: %v", result.Error)
	}

	return user, nil

}

func PostUser(g *models.User) (models.User, error) {

	user := *g
	user.Id = 0
	result := configurations.DBgorm.Create(&user)

	if result.Error != nil {
		return user, fmt.Errorf("error on: %v", result.Error)
	}

	user, err := GetUserById(user.Id)
	if err != nil {
		return user, fmt.Errorf("error on: %v", result.Error)
	}

	return user, nil
}

func DeleteUser(g *models.User) error {

	user := *g
	result := configurations.DBgorm.Delete(&user)

	if result.Error != nil {
		return fmt.Errorf("error on: %v", result.Error)
	}

	return nil
}
