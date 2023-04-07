package models

type Category struct {
	CategoryId   string `json:"id"`
	CategoryName string `json:"name"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}

type CategoryPrimaryKey struct {
	CategoryId string `json:"id"`
}

type CreateCategory struct {
	CategoryName string `json:"category_name"`
}

type UpdateCategory struct {
	CategoryId   string `json:"id"`
	CategoryName string `json:"name"`
	UpdatedAt    string `json:"updated_at"`
}

type GetListCategoryRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}

type GetListCategoryResponse struct {
	Count      int         `json:"count"`
	Categories []*Category `json:"categories"`
}
