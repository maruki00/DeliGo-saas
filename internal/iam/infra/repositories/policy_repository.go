package repositories

import (
	"context"
	"deligo/internal/iam/infra/models"
	shared_models "deligo/internal/shared/infra/models"
	pkgPostgres "deligo/pkg/postgres"

	"gorm.io/gorm"
)

type PolicyRepository struct {
	db *pkgPostgres.PGHandler
}

func (_this *PolicyRepository) Save(ctx context.Context, entity *models.Policy) error {
	return _this.db.DB.Transaction(func(tx *gorm.DB) error {
		sql := `INSERT INTO "policies" ("id", "name", "group_id") VALUES (?, ?, ?)`
		if err := tx.Exec(sql,
			entity.GetID(),
			entity.GetName(),
			entity.GetGroupID(),
		).Error; err != nil {
			return err
		}

		return nil
	})
}
func (_this *PolicyRepository) Delete(ctx context.Context, id string) error {
	err := _this.db.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(&models.Policy{}, map[string]interface{}{
			"id": id,
		}).Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return err
	}

	return nil
}
func (_this *PolicyRepository) FindByID(ctx context.Context, id string) (*models.Policy, error) {
	policy := models.Policy{}
	err := _this.db.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&policy).Where("id = ?", id).Find(&policy).Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return &policy, nil
}
func (_this *PolicyRepository) FindByName(ctx context.Context, name string) (*models.Policy, error) {
	policy := models.Policy{}
	err := _this.db.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&policy).Where("name = ?", name).Find(&policy).Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return &policy, nil
}
func (_this *PolicyRepository) AffectPermission(ctx context.Context, id, policy_id string, permission_id string) error {

	return _this.db.DB.Transaction(func(tx *gorm.DB) error {
		sql := `INSERT INTO "policies_permissions" ("id", "policy_id", "permission_id") VALUES (?, ?, ?)`
		if err := tx.Exec(sql,
			id,
			policy_id,
			permission_id,
		).Error; err != nil {
			return err
		}

		return nil
	})
}
func (_this *PolicyRepository) ListForTenant(ctx context.Context, id string, pagination shared_models.Pagination) ([]*models.Policy, error) {
	policies := make([]*models.Policy, pagination.Limit)
	err := _this.db.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&models.Policy{}).Find(policies).Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return policies, nil
}
