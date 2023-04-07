package postgres

import (
	"app/models"
	"context"
	"fmt"
	"testing"

	"github.com/go-faker/faker/v4"
)

// go test
// go test -run TestCreateCategory

func TestCreateCategory(t *testing.T) {
	tests := []struct {
		Name    string
		Input   *models.CreateCategory
		Output  string
		WantErr bool
	}{
		{
			Name: "Case 1",
			Input: &models.CreateCategory{
				CategoryName: faker.FirstName(),
			},
			WantErr: false,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			id, err := categoryTestRepo.Create(context.Background(), test.Input)

			if test.WantErr {
				t.Errorf("%s: got: %v", test.Name, err)
				return
			}

			if id == "" {
				t.Errorf("%s: got: %v", test.Name, err)
				return
			}
		})
	}
}

func TestGetByIdCategory(t *testing.T) {
	tests := []struct {
		Name    string
		Input   *models.CategoryPrimaryKey
		Output  *models.Category
		WantErr bool
	}{
		{
			Name: "Case 1",
			Input: &models.CategoryPrimaryKey{
				CategoryId: "9420a46e-ef3c-4fe9-85d1-90ae3a4dec2a",
			},
			Output: &models.Category{
				CategoryName: "Marguerite",
			},
			WantErr: false,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			category, err := categoryTestRepo.GetById(context.Background(), test.Input)

			if test.WantErr {
				t.Errorf("%s: got: %v", test.Name, err)
				return
			}

			if category.CategoryName != test.Output.CategoryName {
				t.Errorf("%s: got: %v, expected: %v", test.Name, category, test.Output)
				return
			}
		})
	}
}

func TestUpdateCategory(t *testing.T) {
	tests := []struct {
		Name    string
		Input   *models.UpdateCategory
		Output  *models.Category
		WantErr bool
	}{
		{
			Name: "Case 1",
			Input: &models.UpdateCategory{
				CategoryId:   "9420a46e-ef3c-4fe9-85d1-90ae3a4dec2a",
				CategoryName: "Gucci",
			},
			Output: &models.Category{
				CategoryId:   "9420a46e-ef3c-4fe9-85d1-90ae3a4dec2a",
				CategoryName: "Gucci",
			},
			WantErr: false,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			res, err := categoryTestRepo.Update(context.Background(), test.Input)

			category, err := categoryTestRepo.GetById(context.Background(), &models.CategoryPrimaryKey{
				CategoryId: test.Input.CategoryId,
			})

			if category.CategoryName != test.Output.CategoryName {
				t.Errorf("%s: got: %v, expected: %v", test.Name, category, test.Output)
				return
			}

			if test.WantErr {
				t.Errorf("%s: got: %v", test.Name, err)
				return
			}

			if res <= 0 {
				t.Errorf("%s: got: %v", test.Name, err)
				return
			}
		})
	}
}

func TestDeleteCategory(t *testing.T) {
	tests := []struct {
		Name    string
		Input   *models.CategoryPrimaryKey
		WantErr bool
	}{
		{
			Name: "Case 1",
			Input: &models.CategoryPrimaryKey{
				CategoryId: "9420a46e-ef3c-4fe9-85d1-90ae3a4dec2a",
			},
			WantErr: false,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			res, err := categoryTestRepo.Delete(context.Background(), &models.CategoryPrimaryKey{
				CategoryId: test.Input.CategoryId,
			})

			if test.WantErr {
				t.Errorf("%s: got: %v", test.Name, err)
				return
			}

			if res <= 0 {
				t.Errorf("%s: got: %v", test.Name, err)
				return
			}

			fmt.Println("Deleted Successfully")
		})
	}
}
