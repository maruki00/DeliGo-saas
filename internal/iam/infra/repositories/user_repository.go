package repositories

import (
	"context"
	"deligo/internal/iam/domain/entities"
	valueobjects "deligo/internal/iam/domain/valueobject"
	"deligo/internal/iam/infra/models"
	shared_models "deligo/internal/shared/infra/models"
	pkgPostgres "deligo/pkg/postgres"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *pkgPostgres.PGHandler
}

func NewUserRepository(db *pkgPostgres.PGHandler) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (ur *UserRepository) Save(ctx context.Context, entity entities.UserEntity) error {
	err := ur.db.DB.Transaction(func(tx *gorm.DB) error {
		sql := `INSERT INTO "users" 
			("id", "username", "email", "tenant_id", "password", "password_changed_at", "is_active", "last_login", "mfa_enabled", "mfa_secret", "deleted_at")
			VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

		if err := tx.Exec(sql,
			entity.GetID(),
			entity.GetUsername(),
			entity.GetEmail(),
			entity.GetTenantID(),
			entity.GetPassword(),
			entity.GetPasswordChangedAt(),
			entity.GetIsActive(),
			entity.GetLastLogin(),
			entity.GetMFAEnabled(),
			entity.GetMFASecret(),
			entity.GetDeletedAt(),
		).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func (ur *UserRepository) Delete(ctx context.Context, id valueobjects.ID) error {
	err := ur.db.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("id = ?", id).Delete(&models.User{}).Error; err != nil {
			tx.Rollback()
			return err
		}
		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func (ur *UserRepository) Update(ctx context.Context, id valueobjects.ID, fields map[string]interface{}) error {
	err := ur.db.DB.Transaction(func(tx *gorm.DB) error {
		err := tx.Model(&models.User{}).Where("id = ?", string(id)).Updates(fields).Error
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func (ur *UserRepository) FindByID(ctx context.Context, id valueobjects.ID) (*models.User, error) {
	var user models.User
	err := ur.db.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("id = ?", id).First(&user).Error; err != nil {
			tx.Rollback()
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (ur *UserRepository) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	err := ur.db.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("email = ?", email).First(&user).Error; err != nil {
			tx.Rollback()
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (ur *UserRepository) FindByUsername(ctx context.Context, username string) (*models.User, error) {
	var user models.User
	err := ur.db.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("username = ?", username).First(&user).Error; err != nil {
			tx.Rollback()
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (ur *UserRepository) ListByTenant(ctx context.Context, tenantID valueobjects.ID, pagination shared_models.Pagination) ([]*models.User, error) {
	users := make([]*models.User, pagination.Limit)
	err := ur.db.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&models.User{}).Where("tenant_id = ?", tenantID).Limit(pagination.GetLimit()).Offset(pagination.GetOffset()).Find(&users).Error; err != nil {
			tx.Rollback()
			return err
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (_this *PolicyRepository) AffectRole(ctx context.Context, id, role_id string) error {

	return _this.db.DB.Transaction(func(tx *gorm.DB) error {
		sql := `INSERT INTO "roles_policies" ("id", "role_id", "policy_id") VALUES (?, ?, ?)`
		if err := tx.Exec(sql,
			id,
			role_id,
			policy_id,
		).Error; err != nil {
			return err
		}

		return nil
	})
}
