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

type clientRepo struct {
	db *pgxpool.Pool
}

func NewClientRepo(db *pgxpool.Pool) *clientRepo {
	return &clientRepo{
		db: db,
	}
}

func (c *clientRepo) Create(ctx context.Context, req *models.CreateClient) (string, error) {
	id := uuid.New().String()
	query := `
		INSERT INTO client(
			id,
			first_name, 
			last_name,
			phone_number,
			updated_at
		) VALUES($1, $2, $3, $4, NOW())
	`

	_, err := c.db.Exec(
		ctx,
		query,
		id,
		helper.NewNullString(req.FirstName),
		helper.NewNullString(req.LastName),
		helper.NewNullString(req.PhoneNumber),
	)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (c *clientRepo) GetById(ctx context.Context, req *models.ClientPrimaryKey) (*models.Client, error) {
	var (
		id           sql.NullString
		first_name   sql.NullString
		last_name    sql.NullString
		phone_number sql.NullString
		createdAt    sql.NullString
		updatedAt    sql.NullString
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

	err := c.db.QueryRow(ctx, query, req.ClientId).Scan(
		&id,
		&first_name,
		&last_name,
		&phone_number,
		&createdAt,
		&updatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &models.Client{
		ClientId: id.String,
		FirstName: first_name.String,
		LastName: last_name.String,
		PhoneNumber: phone_number.String,
		CreatedAt: createdAt.String,
		UpdatedAt: updatedAt.String,
	}, nil
}

func (c *clientRepo) GetList(ctx context.Context, req *models.GetListClientRequest) (*models.GetListClientResponse, error) {
	var (
		resp   = models.GetListClientResponse{}
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
		filter += " AND (first_name || ' ' || last_name) ILIKE '%' || '" + req.Search + "' || '%' "
	}

	query += filter + offset + limit

	rows, err := c.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var client models.Client

		err = rows.Scan(
			&client.ClientId,
			&client.FirstName,
			&client.LastName,
			&client.PhoneNumber,
			&client.CreatedAt,
			&client.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		resp.Clients = append(resp.Clients, &client)
	}

	resp.Count = len(resp.Clients)

	return &resp, nil
}

func (c *clientRepo) Update(ctx context.Context, req *models.UpdateClient) (int64, error) {
	var (
		query  string
		params map[string]interface{}
	)

	query = `
		UPDATE 
			client
		SET 
			first_name = :first_name,
			last_name = :last_name,
			phone_number = :phone_number,
			updated_at = NOW()
		WHERE id = :id
	`

	params = map[string]interface{}{
		"id":           req.ClientId,
		"first_name":   req.FirstName,
		"last_name":    req.LastName,
		"phone_number": req.PhoneNumber,
	}

	query, args := helper.ReplaceQueryParams(query, params)

	rowsAffected, err := c.db.Exec(ctx, query, args...)
	if err != nil {
		return 0, nil
	}

	return rowsAffected.RowsAffected(), nil
}

func (c *clientRepo) Delete(ctx context.Context, req *models.ClientPrimaryKey) (int64, error) {
	rowsAffected, err := c.db.Exec(ctx, `DELETE FROM client WHERE id = $1`, req.ClientId)
	if err != nil {
		return 0, err
	}

	return rowsAffected.RowsAffected(), nil
}
