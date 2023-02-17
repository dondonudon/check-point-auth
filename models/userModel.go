package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID                 uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primary_key"`
	Name               string    `gorm:"type:varchar(255);not null"`
	Email              string    `gorm:"uniqueIndex;not null"`
	Password           string    `gorm:"not null"`
	Role               string    `gorm:"type:varchar(255);not null"`
	DOB                time.Time `gorm:"type:date"`
	Age                int       `gorm:"type:int"`
	Gender             string    `gorm:"type:varchar(255)"`
	Occupation         string    `gorm:"type:varchar(255)"`
	VerificationCode   string
	Verified           bool `gorm:"not null"`
	PasswordResetToken string
	PasswordResetAt    time.Time
	LoggedIn           bool
	CreatedAt          time.Time `gorm:"not null"`
	UpdatedAt          time.Time `gorm:"not null"`
}

type SignUpInput struct {
	Name            string    `json:"name" binding:"required"`
	Email           string    `json:"email" binding:"required"`
	Password        string    `json:"password" binding:"required,min=8"`
	PasswordConfirm string    `json:"passwordConfirm" binding:"required"`
	DOB             time.Time `json:"dob" binding:"required"`
	Age             int       `json:"age" binding:"required"`
	Gender          string    `json:"gender" binding:"required"`
	Occupation      string    `json:"occupation" binding:"required"`
}

type SignInInput struct {
	Email    string `json:"email"  binding:"required"`
	Password string `json:"password"  binding:"required"`
}

type UserResponse struct {
	ID         uuid.UUID `json:"id,omitempty"`
	Name       string    `json:"name,omitempty"`
	Email      string    `json:"email,omitempty"`
	Role       string    `json:"role,omitempty"`
	DOB        time.Time `json:"dob,omitempty"`
	Age        int       `json:"age,omitempty"`
	Gender     string    `json:"gender,omitempty"`
	Occupation string    `json:"occupation,omitempty"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// ForgotPasswordInput struct
type ForgotPasswordInput struct {
	Email string `json:"email" binding:"required"`
}

// ResetPasswordInput struct
type ResetPasswordInput struct {
	Password        string `json:"password" binding:"required"`
	PasswordConfirm string `json:"passwordConfirm" binding:"required"`
}
