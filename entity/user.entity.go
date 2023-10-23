package entity

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	// "gin-starter/utils"
)

const (
	userTableName = "main.users"
)

type User struct {
	ID       uuid.UUID    `json:"id"`
	Name     string       `json:"name"`
	Email    string       `json:"email"`
	Password string       `json:"password"`
	DOB      sql.NullTime `json:"dob"`
	Roleid   *UserRole     `gorm:"foreignKey:id" json:"role"`
	Auditable
}

// TableName specifies table name
func (model *User) TableName() string {
	return userTableName
}

func NewUser(
	id uuid.UUID,
	name string,
	email string,
	password string,
	dob sql.NullTime,
	createdBy string,
) *User {
	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	return &User{
		ID:        id,
		Name:      name,
		Email:     email,
		Password:  string(passwordHash),
		DOB:       dob,
		Auditable: NewAuditable(createdBy),
	}
}

// MapUpdateFrom mapping from model
func (model *User) MapUpdateFrom(from *User) *map[string]interface{} {
	if from == nil {
		return &map[string]interface{}{
			"name":       model.Name,
			"email":      model.Email,
			"updated_at": model.UpdatedAt,
		}
	}

	mapped := make(map[string]interface{})

	if model.Name != from.Name {
		mapped["name"] = from.Name
	}

	if model.Email != from.Email {
		mapped["email"] = from.Email
	}

	if model.DOB != from.DOB {
		mapped["dob"] = from.DOB
	}

	mapped["updated_at"] = time.Now()
	return &mapped
}
