package user

import (
	"context"
	"time"

	"encore.app/common/constants"
	"encore.app/database"
	"encore.app/database/models"
	"github.com/asaskevich/govalidator"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//encore:service
type Service struct {
	db *gorm.DB
}

func initService() (*Service, error) {
	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: database.GetDB().Stdlib(),
	}), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &Service{db: db}, nil
}

// Example API for user creation
// In this example, we are creating a user and returning a basic response. It does not include any password hashing.
// In real world, you should hash the password before saving it to the database.
//
// encore:api public method=POST path=/api/v1/user/create
func (s *Service) CreateUser(ctx context.Context, request *constants.CreateUser) (*constants.BasicResponse, error) {
	_, err := govalidator.ValidateStruct(request)
	if err != nil {
		return nil, err
	}

	trx := s.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			trx.Rollback()
		}
	}()

	if err = trx.Error; err != nil {
		return nil, err
	}

	// create user
	user := &models.User{
		FirstName: request.FirstName,
		Surname:   request.LastName,
		Email:     request.Email,
		Phone:     request.PhoneNumber,
		Password:  request.Password,
		CreatedAt: time.Now(),
		UpdatedAt: nil,
		DeletedAt: gorm.DeletedAt{},
	}

	result := trx.Create(user)
	if result.Error != nil {
		trx.Rollback()
		return nil, result.Error
	}

	err = trx.Commit().Error
	if err != nil {
		trx.Rollback()
		return nil, err
	}

	return &constants.BasicResponse{
		Message: "User created successfully",
	}, nil
}
