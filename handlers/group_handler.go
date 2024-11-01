// handlers/group_handler.go
package handlers

import (
	"encoding/json"
	"net/http"
	"splitwise-backend/models"
	"splitwise-backend/services"
	"strconv"

	"github.com/gorilla/mux"
)

type GroupHandler struct {
	groupService services.GroupService
}

func NewGroupHandler(groupService services.GroupService) *GroupHandler {
	return &GroupHandler{groupService: groupService}
}

func (h *GroupHandler) CreateGroup(w http.ResponseWriter, r *http.Request) {
	var group models.Group
	if err := json.NewDecoder(r.Body).Decode(&group); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := h.groupService.CreateGroup(&group); err != nil {
		http.Error(w, "Failed to create group", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(group)
}

func (h *GroupHandler) GetGroup(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	group, err := h.groupService.GetGroupByID(id)
	if err != nil {
		http.Error(w, "Group not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(group)
}

func (h *GroupHandler) UpdateGroup(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var group models.Group
	if err := json.NewDecoder(r.Body).Decode(&group); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	group.ID = id

	if err := h.groupService.UpdateGroup(&group); err != nil {
		http.Error(w, "Failed to update group", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(group)
}

func (h *GroupHandler) DeleteGroup(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	if err := h.groupService.DeleteGroup(id); err != nil {
		http.Error(w, "Failed to delete group", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func RegisterGroupRoutes(router *mux.Router, groupService services.GroupService) {
	handler := NewGroupHandler(groupService)
	router.HandleFunc("/groups", handler.CreateGroup).Methods("POST")
	router.HandleFunc("/groups/{id}", handler.GetGroup).Methods("GET")
	router.HandleFunc("/groups/{id}", handler.UpdateGroup).Methods("PUT")
	router.HandleFunc("/groups/{id}", handler.DeleteGroup).Methods("DELETE")
}
