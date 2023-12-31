package resource

import (
	// "mime/multipart"
	//"os"

	"gin-starter/entity"
	//"gin-starter/utils"
)

type CreateUserRequest struct {
	Name        string `form:"name" json:"name" binding:"required"`
	Email       string `form:"email" json:"email" binding:"required"`
	Password    string `form:"password" json:"password" binding:"required"`
	DOB         string `form:"dob" json:"dob" binding:"required"`
}

type CreateAdminRequest struct {
	Name        string `form:"name" json:"name" binding:"required"`
	Email       string `form:"email" json:"email" binding:"required"`
	Password    string `form:"password" json:"password" binding:"required"`
	DOB         string `form:"dob" json:"dob" binding:"required"`
	RoleId      string `form:"roleid" json:"roleid" `
}

type UpdateAdminRequest struct {
	ID          string                `form:"id" json:"id"`
	Name        string                `form:"name" json:"name"`
	Email       string                `form:"email" json:"email"`
	DOB         string                `form:"dob" json:"dob"`
	RoleId      string                `form:"roleid" json:"roleid" `
}

type UserAdmin struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	OTPIsNull   bool   `json:"otp_is_null"`
	DOB         string `json:"dob"`
	RoleId        *Role  `json:"role"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type GetUsersResponse struct {
	List  []*UserProfile `json:"list"`
	Total int64          `json:"total"`
}

type GetAdminUsersResponse struct {
	List  []*UserAdmin `json:"list"`
	Total int64        `json:"total"`
}

type DeactivateUserRequest struct {
	ID string `uri:"id" binding:"required"`
}

type DeleteAdminRequest struct {
	ID string `uri:"id" binding:"required"`
}

type DeleteUserRequest struct {
	ID string `uri:"id" binding:"required"`
}

type GetUserByIDRequest struct {
	ID string `uri:"id" binding:"required"`
}

type GetAdminUserByIDRequest struct {
	ID string `uri:"id" binding:"required"`
}

type GetAdminUsersRequest struct {
	Query  string `form:"query" json:"query"`
	Sort   string `form:"sort" json:"sort"`
	Order  string `form:"order" json:"order"`
	Limit  int    `form:"limit,default=10" json:"limit"`
	Offset int    `form:"offset,default=0" json:"offset"`
}

func NewUserAdmin(user *entity.User) *UserAdmin {
	// otpIsNull := false
	// if user.OTP.String != "" {
	// 	otpIsNull = true
	// }

	dob := "1970-01-01"
	if user.DOB.Valid {
		dob = user.DOB.Time.Format(timeFormat)
	}

	return &UserAdmin{
		ID:          user.ID.String(),
		Name:        user.Name,
		Email:       user.Email,
		DOB:         dob,
		// Role:        NewRoleResponse(user.UserRole.Role),
		CreatedAt: user.CreatedAt.Format(timeFormat),
		UpdatedAt: user.UpdatedAt.Format(timeFormat),
	}
}
