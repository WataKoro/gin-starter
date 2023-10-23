package entity

import (
	"time"

	"github.com/google/uuid"
)

const (
	userRoleTableName = "main.roles"
)

// UserRole define for table role
type UserRole struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
	Auditable
}

// TableName specifies table name
func (model *UserRole) TableName() string {
	return userRoleTableName
}

// NewUserRole create new entity UserRole
func NewUserRole(
	id uuid.UUID,
	name string,
	createdBy string,
) *UserRole {
	return &UserRole{
		ID:        id,
		Name:      name,
		Auditable: NewAuditable(createdBy),
	}
}

// MapUpdateFrom mapping from model
func (model *UserRole) MapUpdateFrom(from *UserRole) *map[string]interface{} {
	if from == nil {
		return &map[string]interface{}{
			"name":       model.Name,
			"updated_at": model.UpdatedAt,
		}
	}

	mapped := make(map[string]interface{})

	mapped["updated_at"] = time.Now()
	return &mapped
}
