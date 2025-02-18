package serviceaccounts

import (
	"context"

	"github.com/grafana/grafana/pkg/models"
)

type Service interface {
	CreateServiceAccount(ctx context.Context, saForm *CreateServiceaccountForm) (*models.User, error)
	DeleteServiceAccount(ctx context.Context, orgID, serviceAccountID int64) error
}

type Store interface {
	CreateServiceAccount(ctx context.Context, saForm *CreateServiceaccountForm) (*models.User, error)
	ListServiceAccounts(ctx context.Context, orgID, serviceAccountID int64) ([]*models.OrgUserDTO, error)
	RetrieveServiceAccount(ctx context.Context, orgID, serviceAccountID int64) (*models.OrgUserDTO, error)
	DeleteServiceAccount(ctx context.Context, orgID, serviceAccountID int64) error
	UpgradeServiceAccounts(ctx context.Context) error
	ConvertToServiceAccounts(ctx context.Context, keys []int64) error
	ListTokens(ctx context.Context, orgID int64, serviceAccount int64) ([]*models.ApiKey, error)
}
