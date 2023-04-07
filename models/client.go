package models

type Client struct {
	ClientId    string `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	PhoneNumber string `json:"phone_number"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type ClientPrimaryKey struct {
	ClientId string `json:"id"`
}

type CreateClient struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	PhoneNumber string `json:"phone_number"`
}

type UpdateClient struct {
	ClientId    string `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	PhoneNumber string `json:"phone_number"`
	UpdatedAt   string `json:"updated_at"`
}

type GetListClientRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}

type GetListClientResponse struct {
	Count   int       `json:"count"`
	Clients []*Client `json:"clients"`
}
