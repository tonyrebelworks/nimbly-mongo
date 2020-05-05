package request

// RegisterRequest ...
type RegisterRequest struct {
	Fname       string `josn:"fname" validate:"required"`
	Lname       string `josn:"lname" validate:"required"`
	Email       string `josn:"email" validate:"required"`
	Phone       string `josn:"phone" validate:"required"`
	Password    string `josn:"password" validate:"required"`
	RecheckPass string `josn:"recheckPassword" validate:"required"`
}
