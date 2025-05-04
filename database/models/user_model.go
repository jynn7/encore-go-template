package models


import (

	"time"
	"encore.dev/types/uuid"
	"gorm.io/gorm"
	"encore.app/common/constants"

)

type User struct {
	ID               uuid.UUID            `json:"id" gorm:"type:uuid; default:gen_random_uuid()"`
	FirstName        string               `json:"first_name" gorm:"type:varchar(255)" valid:"required~FirstName is required"`
	Surname          string               `json:"surname" gorm:"type:varchar(255)" valid:"required~Surname is required"`
	Email            string               `json:"email" gorm:"type:varchar(255)" valid:"required~Email is required"`
	Password         string               `json:"password" gorm:"type:varchar(255)"`
	Phone            string               `json:"phone" gorm:"type:varchar(255)"`
	//Address          common.Address       `json:"address" gorm:"embedded"`
	Status           constants.UserStatus `json:"status" gorm:"type:varchar(255);default:ACTIVE"`
	CreatedAt        time.Time            `json:"created_at"`
	UpdatedAt        *time.Time           `json:"updated_at"`
	DeletedAt        gorm.DeletedAt       `json:"deleted_at,omitempty"`
}
