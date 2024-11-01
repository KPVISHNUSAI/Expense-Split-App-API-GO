// handlers/expense_handler.go
package handlers

import (
	"encoding/json"
	"net/http"
	"splitwise-backend/models"
	"splitwise-backend/services"
	"strconv"

	"github.com/gorilla/mux"
)

type ExpenseHandler struct {
	expenseService services.ExpenseService
}

func NewExpenseHandler(expenseService services.ExpenseService) *ExpenseHandler {
	return &ExpenseHandler{expenseService: expenseService}
}

func (h *ExpenseHandler) CreateExpense(w http.ResponseWriter, r *http.Request) {
	var expense models.Expense
	if err := json.NewDecoder(r.Body).Decode(&expense); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := h.expenseService.CreateExpense(&expense); err != nil {
		http.Error(w, "Failed to create expense", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(expense)
}

func (h *ExpenseHandler) GetExpense(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	expense, err := h.expenseService.GetExpenseByID(id)
	if err != nil {
		http.Error(w, "Expense not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(expense)
}

func (h *ExpenseHandler) UpdateExpense(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var expense models.Expense
	if err := json.NewDecoder(r.Body).Decode(&expense); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	expense.ID = id

	if err := h.expenseService.UpdateExpense(&expense); err != nil {
		http.Error(w, "Failed to update expense", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(expense)
}

func (h *ExpenseHandler) DeleteExpense(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	if err := h.expenseService.DeleteExpense(id); err != nil {
		http.Error(w, "Failed to delete expense", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func RegisterExpenseRoutes(router *mux.Router, expenseService services.ExpenseService) {
	handler := NewExpenseHandler(expenseService)
	router.HandleFunc("/expenses", handler.CreateExpense).Methods("POST")
	router.HandleFunc("/expenses/{id}", handler.GetExpense).Methods("GET")
	router.HandleFunc("/expenses/{id}", handler.UpdateExpense).Methods("PUT")
	router.HandleFunc("/expenses/{id}", handler.DeleteExpense).Methods("DELETE")
}
