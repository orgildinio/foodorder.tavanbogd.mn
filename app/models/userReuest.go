package models

type UserRequest struct {
	Model struct {
		Role            int         `json:"role"`
		LastName        interface{} `json:"last_name"`
		FirstName       interface{} `json:"first_name"`
		RegisterNumber  string      `json:"register_number"`
		Email           interface{} `json:"email"`
		Phone           interface{} `json:"phone"`
		BirthDate       interface{} `json:"birth_date"`
		Gender          string      `json:"gender"`
		Login           interface{} `json:"login"`
		Password        interface{} `json:"password"`
		PasswordConfirm string      `json:"password_confirm"`
	} `json:"model"`
	EditMode bool `json:"editMode"`
}

type Field struct {
	Field string                 `json:"field"`
	Value interface{}            `json:"value"`
	Props map[string]interface{} `json:"props"`
}
