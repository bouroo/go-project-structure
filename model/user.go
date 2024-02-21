package model

type UserAccount struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password,omitempty"`
}

type UserProfile struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
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
