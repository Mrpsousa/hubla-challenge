package dto

type DtoSellers struct {
	Seller string  `json:"seller"`
	TValue float64 `json:"value"`
}

type DtoCourses struct {
	Type      int8    `json:"type"`
	CreatedAt string  `json:"created_at"`
	Product   string  `json:"product"`
	Value     float64 `json:"value"`
	Seller    string  `json:"seller"`
}

type GetJWTInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type GetJWTOutput struct {
	AccessToken string `json:"access_token"`
}

type CreateUserInput struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
