package dtos

import "github.com/google/uuid"

type ListClientsDTO struct {
	ID                       uuid.UUID `json:"id"`
	ClientID                 string    `json:"client_id"`
	Name                     string    `json:"name"`
	Description              string    `json:"description"`
	Protocol                 string    `json:"protocol"`
	PublicClient             *bool     `json:"public_client"`
	StandardFlowEnabled      *bool     `json:"standard_flow_enabled"`
	ImplicitFlowEnabled      *bool     `json:"implicit_flow_enabled"`
	DirectAccessGrantEnabled *bool     `json:"direct_access_grant_enabled"`
	ServiceAccountsEnabled   *bool     `json:"service_account_enabled"`
	RootURL                  string    `json:"root_url"`
	BaseURL                  string    `json:"base_url"`
	RedirectURIs             string    `json:"redirect_uris"`
	Enabled                  *bool     `json:"enabled"`
}

type GetClientDTO struct {
	ID                       uuid.UUID `json:"id"`
	ClientID                 string    `json:"client_id"`
	Description              string    `json:"description"`
	Protocol                 string    `json:"protocol"`
	PublicClient             *bool     `json:"public_client"`
	StandardFlowEnabled      *bool     `json:"standard_flow_enabled"`
	ImplicitFlowEnabled      *bool     `json:"implicit_flow_enabled"`
	DirectAccessGrantEnabled *bool     `json:"direct_access_grant_enabled"`
	ServiceAccountsEnabled   *bool     `json:"service_account_enabled"`
	RootURL                  string    `json:"root_url"`
	BaseURL                  string    `json:"base_url"`
	RedirectURIs             string    `json:"redirect_uris"`
	Enabled                  *bool     `json:"enabled"`
}

type CreateClientDTO struct {
	ClientID                 string    `json:"client_id"`
	RealmID                  uuid.UUID `json:"realm_id"`
	Description              string    `json:"description"`
	Protocol                 string    `json:"protocol"`
	PublicClient             *bool     `json:"public_client"`
	StandardFlowEnabled      *bool     `json:"standard_flow_enabled"`
	ImplicitFlowEnabled      *bool     `json:"implicit_flow_enabled"`
	DirectAccessGrantEnabled *bool     `json:"direct_access_grant_enabled"`
	ServiceAccountsEnabled   *bool     `json:"service_account_enabled"`
	RootURL                  string    `json:"root_url"`
	BaseURL                  string    `json:"base_url"`
	RedirectURIs             string    `json:"redirect_uris"`
	Enabled                  *bool     `json:"enabled"`
}

type UpdateClientDTO struct {
	ClientID                 string `json:"client_id"`
	Description              string `json:"description"`
	Protocol                 string `json:"protocol"`
	PublicClient             *bool  `json:"public_client"`
	StandardFlowEnabled      *bool  `json:"standard_flow_enabled"`
	ImplicitFlowEnabled      *bool  `json:"implicit_flow_enabled"`
	DirectAccessGrantEnabled *bool  `json:"direct_access_grant_enabled"`
	ServiceAccountsEnabled   *bool  `json:"service_account_enabled"`
	RootURL                  string `json:"root_url"`
	BaseURL                  string `json:"base_url"`
	RedirectURIs             string `json:"redirect_uris"`
}
