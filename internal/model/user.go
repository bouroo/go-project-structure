package model

type UserAccount struct {
	ID       string `json:"id"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password,omitempty" validate:"required,min=6,max=20"`
}

type UserProfile struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Phone     string `json:"phone"`
	Avatar    string `json:"avatar"`
}

type UserAddress struct {
	Number   string `json:"number"`
	Street   string `json:"street"`
	City     string `json:"city"`
	Province string `json:"province"`
	Country  string `json:"country"`
	PostCode string `json:"postCode"`
}

type UserDetail struct {
	UserAccount UserAccount   `json:"userAccount"`
	UserProfile UserProfile   `json:"userProfile"`
	UserAddress []UserAddress `json:"userAddress"`
}
