// repositories/group_repo.go
package repositories

import (
	"context"
	"splitwise-backend/db"
	"splitwise-backend/models"
)

type GroupRepository interface {
	CreateGroup(group *models.Group) error
	GetGroupByID(id int) (*models.Group, error)
	UpdateGroup(group *models.Group) error
	DeleteGroup(id int) error
}

type groupRepository struct{}

func NewGroupRepository() GroupRepository {
	return &groupRepository{}
}

func (r *groupRepository) CreateGroup(group *models.Group) error {
	query := `INSERT INTO groups (name, created_at) VALUES ($1, $2) RETURNING id`
	err := db.DB.QueryRow(context.Background(), query, group.Name, group.CreatedAt).Scan(&group.ID)
	return err
}

func (r *groupRepository) GetGroupByID(id int) (*models.Group, error) {
	group := &models.Group{}
	query := `SELECT id, name, created_at FROM groups WHERE id = $1`
	err := db.DB.QueryRow(context.Background(), query, id).Scan(&group.ID, &group.Name, &group.CreatedAt)
	if err != nil {
		return nil, err
	}
	return group, nil
}

func (r *groupRepository) UpdateGroup(group *models.Group) error {
	query := `UPDATE groups SET name = $1 WHERE id = $2`
	_, err := db.DB.Exec(context.Background(), query, group.Name, group.ID)
	return err
}

func (r *groupRepository) DeleteGroup(id int) error {
	query := `DELETE FROM groups WHERE id = $1`
	_, err := db.DB.Exec(context.Background(), query, id)
	return err
}
