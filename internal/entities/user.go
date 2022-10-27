package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID                  uuid.UUID `gorm:"type:uuid;primary_key"`
	FirstName           string    `gorm:"not null"`
	LastName            string
	Username            string `gorm:"not null;unique"`
	Email               string `gorm:"not null"`
	EmailVerified       *bool  `gorm:"not null;default:false"`
	PhoneCountryCode    string
	PhoneNumber         string
	PhoneNumberVerified *bool
	Status              string          `gorm:"not null"`
	Attributes          []UserAttribute `gorm:"ForeignKey:UserID"`
}

// TableName overrides the table name used by User to `USERS`
func (User) TableName() string {
	return "USER"
}

// BeforeCreate will set a UUID rather than numeric ID.
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return
}

type UserAttribute struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key"`
	Name      string    `gorm:"not null;unique"`
	UserID    string    `gorm:"column:user_id;not null"`
	Value     string
	CreatedAt time.Time `gorm:"autoCreateTime:milli"`
	UpdatedAt time.Time `gorm:"autoUpdateTime:milli"`
}

// TableName overrides the table name used by User to `user_attribute`
func (UserAttribute) TableName() string {
	return "USER_ATTRIBUTE"
}

// BeforeCreate will set a UUID rather than numeric ID.
func (u *UserAttribute) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return
}
