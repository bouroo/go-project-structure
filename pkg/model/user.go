package model

type UserAccount struct {
	ID       string `json:"id"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password,omitempty" validate:"required,min=6,max=20"`
}

type UserProfile struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Phone     string `json:"phone"`
	Avatar    string `json:"avatar"`
}

type UserAddress struct {
	Number   string `json:"number"`
	Street   string `json:"street"`
	City     string `json:"city"`
	Province string `json:"province"`
	Country  string `json:"country"`
	PostCode string `json:"postal_code"`
}

type UserDetail struct {
	UserAccount UserAccount   `json:"user_account"`
	UserProfile UserProfile   `json:"user_profile"`
	UserAddress []UserAddress `json:"user_addresses"`
}
