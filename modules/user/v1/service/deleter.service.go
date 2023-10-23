package service

import (
	"context"
	"gin-starter/common/errors"
	"gin-starter/config"
	"gin-starter/modules/user/v1/repository"

	"github.com/google/uuid"
)

// UserDeleter is a service for user
type UserDeleter struct {
	cfg      config.Config
	userRepo repository.UserRepositoryUseCase
	roleRepo repository.RoleRepositoryUseCase
	userRoleRepo repository.UserRoleRepositoryUseCase

}

// DeleteUser implements UserDeleterUseCase.
func (*UserDeleter) DeleteUser(ctx context.Context, id uuid.UUID) error {
	panic("unimplemented")
}

// UserDeleterUseCase is a use case for user
type UserDeleterUseCase interface {
	// DeleteUser deletes user
	DeleteUsers(ctx context.Context, id uuid.UUID) error
	// DeleteAdmin deletes admin
	DeleteAdmin(ctx context.Context, id uuid.UUID) error
	// DeleteRole deletes role
	DeleteRole(ctx context.Context, id uuid.UUID, deletedBy string) error
	// DeleteUser deletes user role
	DeleteUserRole(ctx context.Context, id uuid.UUID) error
}

// NewUserDeleter creates a new UserDeleter
func NewUserDeleter(
	cfg config.Config,
	userRepo repository.UserRepositoryUseCase,
	roleRepo repository.RoleRepositoryUseCase,
	userRoleRepo repository.UserRoleRepositoryUseCase,

) *UserDeleter {
	return &UserDeleter{
		cfg:      cfg,
		userRepo: userRepo,
		roleRepo: roleRepo,
		userRoleRepo: userRoleRepo,
	}
}

// DeleteAdmin deletes admin
func (ud *UserDeleter) DeleteAdmin(ctx context.Context, id uuid.UUID) error {
	if err := ud.userRepo.DeleteAdmin(ctx, id); err != nil {
		return errors.ErrInternalServerError.Error()
	}

	return nil
}

func (ud *UserDeleter) DeleteUsers(ctx context.Context, id uuid.UUID) error {
	if err := ud.userRepo.DeleteUsers(ctx, id); err != nil {
		return errors.ErrInternalServerError.Error()
	}

	return nil
}

// DeleteRole deletes role
func (ud *UserDeleter) DeleteRole(ctx context.Context, id uuid.UUID, deletedBy string) error {
	if err := ud.roleRepo.Delete(ctx, id, deletedBy); err != nil {
		return errors.ErrInternalServerError.Error()
	}

	return nil
}

// DeleteUserRole deletes user role
func (ud *UserDeleter) DeleteUserRole(ctx context.Context, id uuid.UUID) error {
	if err := ud.userRoleRepo.DeleteUserRole(ctx, id); err != nil {
		return errors.ErrInternalServerError.Error()
	}

	return nil
}
