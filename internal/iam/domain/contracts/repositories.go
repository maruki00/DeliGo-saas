package contracts

import (
	"context"
	"deligo/internal/iam/infra/models"
)

type Pagination struct{}

type IUserRepository interface {
	FindByID(ctx context.Context, id string) (*models.User, error)
	FindByEmail(ctx context.Context, email string) (*models.User, error)
	FindByUsername(ctx context.Context, username string) (*models.User, error)
	Save(ctx context.Context, user *models.User) error
	Update(ctx context.Context, user *models.User) error
	Delete(ctx context.Context, id string) error
	ListByTenant(ctx context.Context, tenantID string, pagination Pagination) ([]*models.User, error)
}

type IPolicyRepository interface {
	FindByID(ctx context.Context, id string) (*models.Policy, error)
	FindByName(ctx context.Context, name string) (*models.Policy, error)
	Save(ctx context.Context, policy *models.Policy) error
	Delete(ctx context.Context, id string) error
	ListForTenant(ctx context.Context, tenantID string) ([]*models.Policy, error)
}

type IPermissionRepository interface {
	FindByID(ctx context.Context, id string) (*models.Permission, error)
	FindByPolicyID(ctx context.Context, policyID string) ([]*models.Permission, error)
	Save(ctx context.Context, permission *models.Permission) error
	Delete(ctx context.Context, id string) error
}

type IGroupRepository interface {
	AssignUserToGroup(ctx context.Context, userID, groupID string) error
	RemoveUserFromGroup(ctx context.Context, userID, groupID string) error
	GetUserGroups(ctx context.Context, userID string) ([]*models.Group, error)
}
