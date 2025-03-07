package model

type DeleteUserResponse struct {
	DeletedUserID string `json:"deletedUserId"`
}

type Mutation struct {
}

type Query struct {
}

type User struct {
	ID       string `json:"_id"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
	Email    string `json:"email"`
	Relation string `json:"relation"`
}

type CreateUserInput struct {
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
	Email    string `json:"email"`
	Relation string `json:"relation"`
}

type FilterInput struct {
	Field string `json:"field"`
	Value string `json:"value"`
}

type UpdateUserInput struct {
	ID       string  `json:"_id"`
	Name     *string `json:"name,omitempty"`
	Phone    *string `json:"phone,omitempty"`
	Address  *string `json:"address,omitempty"`
	Email    *string `json:"email,omitempty"`
	Relation *string `json:"relation,omitempty"`
}
