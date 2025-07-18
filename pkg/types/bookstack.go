package types

// BookstackUser represents a user in the BookStack system
type BookstackUser struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Roles      []int  `json:"roles"`
	Language   string `json:"language"`
	SendInvite bool   `json:"sendInvite"`
}

// BookstackUserResponse represents the API response for BookStack users
type BookstackUserResponse struct {
	Data  []BookstackUser `json:"data"`
	Total int             `json:"total"`
}

// BookstackUserCreateRequest represents the request payload for creating a user
type BookstackUserCreateRequest struct {
	Name       string `json:"name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Roles      []int  `json:"roles"`
	Language   string `json:"language"`
	SendInvite bool   `json:"sendInvite"`
}

// BookstackUserUpdateRequest represents the request payload for updating a user
type BookstackUserUpdateRequest struct {
	Name     string `json:"name,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
	Roles    []int  `json:"roles,omitempty"`
	Language string `json:"language,omitempty"`
}
