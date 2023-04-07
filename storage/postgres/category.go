package postgres

import (
	"app/models"
	"app/pkg/helper"
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type categoryRepo struct {
	db *pgxpool.Pool
}

func NewCategoryRepo(db *pgxpool.Pool) *categoryRepo {
	return &categoryRepo{
		db: db,
	}
}

func (c *categoryRepo) Create(ctx context.Context, req *models.CreateCategory) (string, error) {
	id := uuid.New().String()
	query := `
		INSERT INTO category(
			id,
			name, 
			updated_at
		) VALUES($1, $2, NOW())
	`

	_, err := c.db.Exec(
		ctx,
		query,
		id,
		helper.NewNullString(req.CategoryName),
	)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (c *categoryRepo) GetById(ctx context.Context, req *models.CategoryPrimaryKey) (*models.Category, error) {
	var (
		id        sql.NullString
		name      sql.NullString
		createdAt sql.NullString
		updatedAt sql.NullString
	)

	query := `
		SELECT 
			id,
			name,
			created_at,
			updated_at
		FROM category
		WHERE id = $1
	`

	err := c.db.QueryRow(ctx, query, req.CategoryId).Scan(
		&id,
		&name,
		&createdAt,
		&updatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &models.Category{
		CategoryId:   id.String,
		CategoryName: name.String,
		CreatedAt:    createdAt.String,
		UpdatedAt:    updatedAt.String,
	}, nil
}

func (c *categoryRepo) GetList(ctx context.Context, req *models.GetListCategoryRequest) (*models.GetListCategoryResponse, error) {
	var (
		resp   = models.GetListCategoryResponse{}
		filter = " WHERE TRUE "
		offset = " OFFSET 0"
		limit  = " LIMIT 5"

		query string
	)

	if req.Limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
	}
	if req.Offset > 0 {
		offset = fmt.Sprintf(" OFFSET %d", req.Offset)
	}
	if len(req.Search) > 0 {
		filter += " AND category_name ILIKE '%' || '" + req.Search + "' || '%' "
	}

	query += filter + offset + limit

	rows, err := c.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var category models.Category

		err = rows.Scan(
			&category.CategoryId,
			&category.CategoryName,
			&category.CreatedAt,
			&category.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		resp.Categories = append(resp.Categories, &category)
	}

	resp.Count = len(resp.Categories)

	return &resp, nil
}

func (c *categoryRepo) Update(ctx context.Context, req *models.UpdateCategory) (int64, error) {
	var (
		query  string
		params map[string]interface{}
	)

	query = `
		UPDATE 
			category
		SET 
			name = :name,
			updated_at = now()
		WHERE id = :id
	`

	params = map[string]interface{}{
		"id":   req.CategoryId,
		"name": req.CategoryName,
	}

	query, args := helper.ReplaceQueryParams(query, params)

	rowsAffected, err := c.db.Exec(ctx, query, args...)
	if err != nil {
		return 0, nil
	}

	return rowsAffected.RowsAffected(), nil
}

func (c *categoryRepo) Delete(ctx context.Context, req *models.CategoryPrimaryKey) (int64, error) {
	rowsAffected, err := c.db.Exec(ctx, `DELETE FROM category WHERE id = $1`, req.CategoryId)
	if err != nil {
		return 0, err
	}

	return rowsAffected.RowsAffected(), nil
}
