package data

import (
	"fmt"

	"github.com/guilhermecoelho/kakeibo/configurations"
	"github.com/guilhermecoelho/kakeibo/models"
)

func GetExpenses() (models.Expenses, error) {

	expense := models.Expenses{}
	result := configurations.DBgorm.Find(&expense)

	if result.Error != nil {
		return expense, fmt.Errorf("error on: %v", result.Error)
	}

	return expense, nil
}

func GetExpensesById(id int) (models.Expense, error) {

	expense := models.Expense{}
	result := configurations.DBgorm.First(&expense, id)

	if result.Error != nil {
		return expense, fmt.Errorf("error on: %v", result.Error)
	}

	return expense, nil
}

func GetExpensesByPeriod(dateStart string, dateFinish string) (models.Expenses, error) {

	expenses := models.Expenses{}
	result := configurations.DBgorm.Where("date BETWEEN ? AND ?", dateStart, dateFinish).Find(&expenses)

	if result.Error != nil {
		return expenses, fmt.Errorf("error on: %v", result.Error)
	}

	return expenses, nil
}

func PutExpenses(g *models.Expense) (models.Expense, error) {

	expense := *g
	result := configurations.DBgorm.Save(&expense)

	if result.Error != nil {
		return expense, fmt.Errorf("error on: %v", result.Error)
	}

	expense, err := GetExpensesById(expense.Id)
	if err != nil {
		return expense, fmt.Errorf("error on: %v", result.Error)
	}

	return expense, nil

}

func PostExpenses(g *models.Expense) (models.Expense, error) {

	expense := *g
	expense.Id = 0
	result := configurations.DBgorm.Create(&expense)

	if result.Error != nil {
		return expense, fmt.Errorf("error on: %v", result.Error)
	}

	expense, err := GetExpensesById(expense.Id)
	if err != nil {
		return expense, fmt.Errorf("error on: %v", result.Error)
	}

	return expense, nil
}

func DeleteExpenses(g *models.Expense) error {

	expense := *g
	result := configurations.DBgorm.Delete(&expense)

	if result.Error != nil {
		return fmt.Errorf("error on: %v", result.Error)
	}

	return nil
}
