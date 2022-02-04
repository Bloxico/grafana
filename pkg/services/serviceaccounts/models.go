package serviceaccounts

import (
	"time"

	"github.com/grafana/grafana/pkg/services/accesscontrol"
)

var (
	ScopeAll = "serviceaccounts:*"
	ScopeID  = accesscontrol.Scope("serviceaccounts", "id", accesscontrol.Parameter(":serviceaccountId"))
)

const (
	ActionRead   = "serviceaccounts:read"
	ActionWrite  = "serviceaccounts:write"
	ActionCreate = "serviceaccounts:create"
	ActionDelete = "serviceaccounts:delete"
)

type ServiceAccount struct {
	Id int64
}

type CreateServiceaccountForm struct {
	OrgID       int64  `json:"-"`
	Name        string `json:"name" binding:"Required"`
	DisplayName string `json:"displayName"`
	Description string `json:"description"`
}

type ServiceAccountIdDTO struct {
	Id      int64  `json:"id"`
	Message string `json:"message"`
}

type ServiceAccountDTO struct {
	Id     int64  `json:"id"`
	Email  string `json:"email"`
	Name   string `json:"name"`
	Login  string `json:"login"`
	OrgId  int64  `json:"orgId"`
	Tokens int64  `json:"tokens"`
}

type ServiceAccountProfileDTO struct {
	Id            int64           `json:"id"`
	Email         string          `json:"email"`
	Name          string          `json:"name"`
	Login         string          `json:"login"`
	OrgId         int64           `json:"orgId"`
	IsDisabled    bool            `json:"isDisabled"`
	AuthLabels    []string        `json:"authLabels"`
	UpdatedAt     time.Time       `json:"updatedAt"`
	CreatedAt     time.Time       `json:"createdAt"`
	AvatarUrl     string          `json:"avatarUrl"`
	AccessControl map[string]bool `json:"accessControl,omitempty"`
}
