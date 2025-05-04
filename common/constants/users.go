package constants

type UserStatus string

const (
	UserActive   UserStatus = "ACTIVE"
	UserInactive UserStatus = "INACTIVE"
)

type CreateUser struct {
	FirstName   string `json:"firstName" valid:"required;First Name is required"`
	LastName    string `json:"lastName" valid:"required;Last Name is required"`
	Email       string `json:"email" valid:"required;email;Email is required"`
	Password    string `json:"password" valid:"required;Password is required"`
	PhoneNumber string `json:"phoneNumber" valid:"required;Phone Number is required"`
}
