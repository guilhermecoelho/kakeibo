package data

import (
	"fmt"

	"github.com/guilhermecoelho/kakeibo/configurations"
	"github.com/guilhermecoelho/kakeibo/models"
)

func GetIncomes() (models.Incomes, error) {

	income := models.Incomes{}
	result := configurations.DBgorm.Find(&income)

	if result.Error != nil {
		return income, fmt.Errorf("error on: %v", result.Error)
	}

	return income, nil
}

func GetIncomesById(id int) (models.Income, error) {

	income := models.Income{}
	result := configurations.DBgorm.First(&income, id)

	if result.Error != nil {
		return income, fmt.Errorf("error on: %v", result.Error)
	}

	return income, nil
}

func GetIncomesByPeriod(dateStart string, dateFinish string) (models.Incomes, error) {

	incomes := models.Incomes{}
	result := configurations.DBgorm.Where("date BETWEEN ? AND ?", dateStart, dateFinish).Find(&incomes)

	if result.Error != nil {
		return incomes, fmt.Errorf("error on: %v", result.Error)
	}

	return incomes, nil
}

func PutIncomes(g *models.Income) (models.Income, error) {

	income := *g
	result := configurations.DBgorm.Save(&income)

	if result.Error != nil {
		return income, fmt.Errorf("error on: %v", result.Error)
	}

	income, err := GetIncomesById(income.Id)
	if err != nil {
		return income, fmt.Errorf("error on: %v", result.Error)
	}

	return income, nil

}

func PostIncomes(g *models.Income) (models.Income, error) {

	income := *g
	income.Id = 0
	result := configurations.DBgorm.Create(&income)

	if result.Error != nil {
		return income, fmt.Errorf("error on: %v", result.Error)
	}

	income, err := GetIncomesById(income.Id)
	if err != nil {
		return income, fmt.Errorf("error on: %v", result.Error)
	}

	return income, nil
}

func DeleteIncomes(g *models.Income) error {

	income := *g
	result := configurations.DBgorm.Delete(&income)

	if result.Error != nil {
		return fmt.Errorf("error on: %v", result.Error)
	}

	return nil
}
