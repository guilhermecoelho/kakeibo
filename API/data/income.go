package data

import (
	"fmt"

	"github.com/guilhermecoelho/kakeibo/configurations"
	"github.com/guilhermecoelho/kakeibo/models"
)

type IncomeInterface interface {
	Find() (models.Incomes, error)
}

type IncomeRequest struct {
}

func Find() (models.Incomes, error) {

	incomes := models.Incomes{}
	result := configurations.DBgorm.Find(&incomes)

	if result.Error != nil {
		return incomes, fmt.Errorf("error on: %v", result.Error)
	}

	return incomes, nil
}

func GetIncomesTest(req IncomeInterface) (models.Incomes, error) {
	return req.Find()
}

func (IncomeRequest) GetIncomesTest2(req IncomeInterface) (models.Incomes, error) {

	income := models.Incomes{}
	result := configurations.DBgorm.Find(&income)

	if result.Error != nil {
		return income, fmt.Errorf("error on: %v", result.Error)
	}

	return income, nil
}

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
