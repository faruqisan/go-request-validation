package requests

// CreateUserRequest define request for create an user object
type CreateUserRequest struct {
	Name     string `validate:"required"`
	Email    string `validate:"required,email"`
	Password string `validate:"required"`
}

// UpdateUserRequest define request for update an user object
type UpdateUserRequest struct {
	ID   int64  `validate:"required"`
	Name string `validate:"required"`
}
