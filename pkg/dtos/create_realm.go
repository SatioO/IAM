package dtos

type CreateRealmDTO struct {
	Name                    string `json:"name;omitempty"`
	DisplayName             string `json:"display_name;omitempty"`
	Logo                    string `json:"logo;omitempty"`
	SupportURL              string `json:"support_url;omitempty"`
	SupportEmail            string `json:"support_email;omitempty"`
	DuplicateEmailAllowed   bool   `json:"duplicate_email_allowed;omitempty"`
	DuplicatePhoneAllowed   bool   `json:"duplicate_phone_allowed;omitempty"`
	RegisterEmailAsUsername bool   `json:"register_email_as_username;omitempty"`
	RegisterPhoneAsUsername bool   `json:"register_phone_as_username;omitempty"`
	Enabled                 bool   `json:"enabled;omitempty;"`
}
