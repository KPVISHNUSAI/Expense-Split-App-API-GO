// services/group_service.go
package services

import (
	"splitwise-backend/models"
	"splitwise-backend/repositories"
)

type GroupService interface {
	CreateGroup(group *models.Group) error
	GetGroupByID(id int) (*models.Group, error)
	UpdateGroup(group *models.Group) error
	DeleteGroup(id int) error
}

type groupService struct {
	groupRepo repositories.GroupRepository
}

func NewGroupService(groupRepo repositories.GroupRepository) GroupService {
	return &groupService{groupRepo: groupRepo}
}

func (s *groupService) CreateGroup(group *models.Group) error {
	return s.groupRepo.CreateGroup(group)
}

func (s *groupService) GetGroupByID(id int) (*models.Group, error) {
	return s.groupRepo.GetGroupByID(id)
}

func (s *groupService) UpdateGroup(group *models.Group) error {
	return s.groupRepo.UpdateGroup(group)
}

func (s *groupService) DeleteGroup(id int) error {
	return s.groupRepo.DeleteGroup(id)
}
