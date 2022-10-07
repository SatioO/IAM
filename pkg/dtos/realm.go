package dtos

type ListRealmDTO struct {
	Name                    string            `json:"name"`
	DisplayName             string            `json:"display_name"`
	Logo                    string            `json:"logo"`
	SupportURL              string            `json:"support_url"`
	SupportEmail            string            `json:"support_email"`
	DuplicateEmailAllowed   bool              `json:"duplicate_email_allowed"`
	DuplicatePhoneAllowed   bool              `json:"duplicate_phone_allowed"`
	RegisterEmailAsUsername bool              `json:"register_email_as_username"`
	RegisterPhoneAsUsername bool              `json:"register_phone_as_username"`
	Enabled                 bool              `json:"enabled"`
	Attributes              map[string]string `json:"attributes"`
}

type CreateRealmDTO struct {
	Name                    string            `json:"name"`
	DisplayName             string            `json:"display_name"`
	Logo                    string            `json:"logo"`
	SupportURL              string            `json:"support_url"`
	SupportEmail            string            `json:"support_email"`
	DuplicateEmailAllowed   bool              `json:"duplicate_email_allowed"`
	DuplicatePhoneAllowed   bool              `json:"duplicate_phone_allowed"`
	RegisterEmailAsUsername bool              `json:"register_email_as_username"`
	RegisterPhoneAsUsername bool              `json:"register_phone_as_username"`
	Enabled                 bool              `json:"enabled"`
	Attributes              map[string]string `json:"attributes"`
}
