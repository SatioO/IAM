package dtos

import "github.com/google/uuid"

type ListUserDTO struct {
	ID                  uuid.UUID         `json:"id"`
	FirstName           string            `json:"first_name"`
	LastName            string            `json:"last_name"`
	Username            string            `json:"username"`
	Email               string            `json:"email"`
	EmailVerified       bool              `json:"email_verified"`
	PhoneCountryCode    string            `json:"phone_country_code"`
	PhoneNumber         string            `json:"phone_number"`
	PhoneNumberVerified bool              `json:"phone_number_verified"`
	Status              string            `json:"status"`
	Attributes          map[string]string `json:"attributes"`
}

type GetUserDetailsDTO struct {
	ID                  uuid.UUID         `json:"id"`
	FirstName           string            `json:"first_name"`
	LastName            string            `json:"last_name"`
	Username            string            `json:"username"`
	Email               string            `json:"email"`
	EmailVerified       bool              `json:"email_verified"`
	PhoneCountryCode    string            `json:"phone_country_code"`
	PhoneNumber         string            `json:"phone_number"`
	PhoneNumberVerified bool              `json:"phone_number_verified"`
	Status              string            `json:"status"`
	Attributes          map[string]string `json:"attributes"`
}

type CreateUserDTO struct {
	FirstName           string            `json:"first_name"`
	LastName            string            `json:"last_name"`
	Username            string            `json:"username"`
	Email               string            `json:"email"`
	EmailVerified       bool              `json:"email_verified"`
	PhoneCountryCode    string            `json:"phone_country_code"`
	PhoneNumber         string            `json:"phone_number"`
	PhoneNumberVerified bool              `json:"phone_number_verified"`
	Attributes          map[string]string `json:"attributes"`
}

type UpdateUserDTO struct {
	FirstName           string            `json:"first_name"`
	LastName            string            `json:"last_name"`
	Username            string            `json:"username"`
	Email               string            `json:"email"`
	EmailVerified       bool              `json:"email_verified"`
	PhoneCountryCode    string            `json:"phone_country_code"`
	PhoneNumber         string            `json:"phone_number"`
	PhoneNumberVerified bool              `json:"phone_number_verified"`
	Attributes          map[string]string `json:"attributes"`
}
