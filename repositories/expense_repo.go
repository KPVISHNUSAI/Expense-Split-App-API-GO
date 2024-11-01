// repositories/expense_repo.go
package repositories

import (
	"context"
	"splitwise-backend/db"
	"splitwise-backend/models"
)

type ExpenseRepository interface {
	CreateExpense(expense *models.Expense) error
	GetExpenseByID(id int) (*models.Expense, error)
	UpdateExpense(expense *models.Expense) error
	DeleteExpense(id int) error
}

type expenseRepository struct{}

func NewExpenseRepository() ExpenseRepository {
	return &expenseRepository{}
}

func (r *expenseRepository) CreateExpense(expense *models.Expense) error {
	query := `INSERT INTO expenses (amount, description, paid_by, group_id, created_at) VALUES ($1, $2, $3, $4, $5) RETURNING id`
	err := db.DB.QueryRow(context.Background(), query, expense.Amount, expense.Description, expense.PaidBy, expense.GroupID, expense.CreatedAt).Scan(&expense.ID)
	return err
}

func (r *expenseRepository) GetExpenseByID(id int) (*models.Expense, error) {
	expense := &models.Expense{}
	query := `SELECT id, amount, description, paid_by, group_id, created_at FROM expenses WHERE id = $1`
	err := db.DB.QueryRow(context.Background(), query, id).Scan(&expense.ID, &expense.Amount, &expense.Description, &expense.PaidBy, &expense.GroupID, &expense.CreatedAt)
	if err != nil {
		return nil, err
	}
	return expense, nil
}

func (r *expenseRepository) UpdateExpense(expense *models.Expense) error {
	query := `UPDATE expenses SET amount = $1, description = $2, paid_by = $3, group_id = $4 WHERE id = $5`
	_, err := db.DB.Exec(context.Background(), query, expense.Amount, expense.Description, expense.PaidBy, expense.GroupID, expense.ID)
	return err
}

func (r *expenseRepository) DeleteExpense(id int) error {
	query := `DELETE FROM expenses WHERE id = $1`
	_, err := db.DB.Exec(context.Background(), query, id)
	return err
}
