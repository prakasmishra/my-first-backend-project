package user

type createUserDto struct {
	FisrtName string `json:"fisrtName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
	Email     string `json:"email" validate:"required"`
}

type updateUserDto struct {
	FisrtName string `json:"fisrtName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}
