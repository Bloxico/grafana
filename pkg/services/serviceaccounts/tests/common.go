package tests

import (
	"context"
	"testing"

	"github.com/grafana/grafana/pkg/models"
	"github.com/grafana/grafana/pkg/services/accesscontrol"
	accesscontrolmock "github.com/grafana/grafana/pkg/services/accesscontrol/mock"
	"github.com/grafana/grafana/pkg/services/serviceaccounts"
	"github.com/grafana/grafana/pkg/services/sqlstore"
	"github.com/stretchr/testify/require"
)

type TestUser struct {
	Login            string
	IsServiceAccount bool
}

func SetupUserServiceAccount(t *testing.T, sqlStore *sqlstore.SQLStore, testUser TestUser) *models.User {
	u1, err := sqlStore.CreateUser(context.Background(), models.CreateUserCommand{
		Login:            testUser.Login,
		IsServiceAccount: testUser.IsServiceAccount,
	})
	require.NoError(t, err)
	return u1
}

// create mock for serviceaccountservice
type ServiceAccountMock struct{}

func (s *ServiceAccountMock) CreateServiceAccount(ctx context.Context, saForm *serviceaccounts.CreateServiceaccountForm) (*models.User, error) {
	return nil, nil
}

func (s *ServiceAccountMock) DeleteServiceAccount(ctx context.Context, orgID, serviceAccountID int64) error {
	return nil
}

func (s *ServiceAccountMock) Migrated(ctx context.Context, orgID int64) bool {
	return false
}

func SetupMockAccesscontrol(t *testing.T, userpermissionsfunc func(c context.Context, siu *models.SignedInUser) ([]*accesscontrol.Permission, error), disableAccessControl bool) *accesscontrolmock.Mock {
	t.Helper()
	acmock := accesscontrolmock.New()
	if disableAccessControl {
		acmock = acmock.WithDisabled()
	}
	acmock.GetUserPermissionsFunc = userpermissionsfunc
	return acmock
}

// this is a way to see
// that the Mock implements the store interface
var _ serviceaccounts.Store = new(ServiceAccountsStoreMock)

type Calls struct {
	CreateServiceAccount   []interface{}
	ListServiceAccounts    []interface{}
	RetrieveServiceAccount []interface{}
	DeleteServiceAccount   []interface{}
	UpgradeServiceAccounts []interface{}
	ConvertServiceAccounts []interface{}
	ListTokens             []interface{}
}

type ServiceAccountsStoreMock struct {
	Calls Calls
}

func (s *ServiceAccountsStoreMock) CreateServiceAccount(ctx context.Context, cmd *serviceaccounts.CreateServiceaccountForm) (*models.User, error) {
	// now we can test that the mock has these calls when we call the function
	s.Calls.CreateServiceAccount = append(s.Calls.CreateServiceAccount, []interface{}{ctx, cmd})
	return nil, nil
}

func (s *ServiceAccountsStoreMock) DeleteServiceAccount(ctx context.Context, orgID, serviceAccountID int64) error {
	// now we can test that the mock has these calls when we call the function
	s.Calls.DeleteServiceAccount = append(s.Calls.DeleteServiceAccount, []interface{}{ctx, orgID, serviceAccountID})
	return nil
}

func (s *ServiceAccountsStoreMock) UpgradeServiceAccounts(ctx context.Context) error {
	s.Calls.UpgradeServiceAccounts = append(s.Calls.UpgradeServiceAccounts, []interface{}{ctx})
	return nil
}

func (s *ServiceAccountsStoreMock) ConvertToServiceAccounts(ctx context.Context, keys []int64) error {
	s.Calls.ConvertServiceAccounts = append(s.Calls.ConvertServiceAccounts, []interface{}{ctx})
	return nil
}

func (s *ServiceAccountsStoreMock) ListTokens(ctx context.Context, orgID int64, serviceAccount int64) ([]*models.ApiKey, error) {
	s.Calls.ListTokens = append(s.Calls.ListTokens, []interface{}{ctx, orgID, serviceAccount})
	return nil, nil
}
func (s *ServiceAccountsStoreMock) ListServiceAccounts(ctx context.Context, orgID int64, serviceAccountID int64) ([]*models.OrgUserDTO, error) {
	s.Calls.ListServiceAccounts = append(s.Calls.ListServiceAccounts, []interface{}{ctx, orgID})
	return nil, nil
}

func (s *ServiceAccountsStoreMock) RetrieveServiceAccount(ctx context.Context, orgID, serviceAccountID int64) (*models.OrgUserDTO, error) {
	s.Calls.RetrieveServiceAccount = append(s.Calls.RetrieveServiceAccount, []interface{}{ctx, orgID, serviceAccountID})
	return nil, nil
}
