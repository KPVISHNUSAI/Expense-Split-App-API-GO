// services/expense_service.go
package services

import (
	"splitwise-backend/models"
	"splitwise-backend/repositories"
)

type ExpenseService interface {
	CreateExpense(expense *models.Expense) error
	GetExpenseByID(id int) (*models.Expense, error)
	UpdateExpense(expense *models.Expense) error
	DeleteExpense(id int) error
}

type expenseService struct {
	expenseRepo repositories.ExpenseRepository
}

func NewExpenseService(expenseRepo repositories.ExpenseRepository) ExpenseService {
	return &expenseService{expenseRepo: expenseRepo}
}

func (s *expenseService) CreateExpense(expense *models.Expense) error {
	return s.expenseRepo.CreateExpense(expense)
}

func (s *expenseService) GetExpenseByID(id int) (*models.Expense, error) {
	return s.expenseRepo.GetExpenseByID(id)
}

func (s *expenseService) UpdateExpense(expense *models.Expense) error {
	return s.expenseRepo.UpdateExpense(expense)
}

func (s *expenseService) DeleteExpense(id int) error {
	return s.expenseRepo.DeleteExpense(id)
}
