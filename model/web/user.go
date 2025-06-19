package web

type RegisterUserRequest struct {
    Name     string `validate:"required" json:"name"`
    Email    string `validate:"required,email" json:"email"`
    Password string `validate:"required" json:"password"`
    Role string `validate:"required" json:"role"`
}

type UpdateUserRequest struct {
    Name     string `validate:"required" json:"name"`
    Email    string `validate:"required,email" json:"email"`
    Password string `validate:"required" json:"password"`
    Role string `validate:"required" json:"role"`
}

type LoginUserRequest struct {
    Email    string `validate:"required,email" json:"email"`
    Password string `validate:"required" json:"password"`
}
