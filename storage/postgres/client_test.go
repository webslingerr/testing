package postgres

import (
	"app/models"
	"context"
	"fmt"
	"testing"
)

// go test
// go test -run TestCreateClient

func TestCreateClient(t *testing.T) {
	tests := []struct {
		Name    string
		Input   *models.CreateClient
		Output  string
		WantErr bool
	}{
		{
			Name: "Case 1",
			Input: &models.CreateClient{
				FirstName:   "",
				LastName:    "",
				PhoneNumber: "",
			},
			WantErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			id, err := clientTestRepo.Create(context.Background(), test.Input)

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

func TestGetByIdClient(t *testing.T) {
	tests := []struct {
		Name    string
		Input   *models.ClientPrimaryKey
		Output  *models.Client
		WantErr bool
	}{
		{
			Name: "Case 1",
			Input: &models.ClientPrimaryKey{
				ClientId: "9420a46e-ef3c-4fe9-85d1-90ae3a4dec2a",
			},
			Output: &models.Client{
				ClientId: "",
				FirstName: "",
				LastName: "",
				PhoneNumber: "",
				CreatedAt: "",
				UpdatedAt: "",
			},
			WantErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			client, err := clientTestRepo.GetById(context.Background(), test.Input)

			if test.WantErr {
				t.Errorf("%s: got: %v", test.Name, err)
				return
			}

			if client.FirstName != test.Output.FirstName || client.LastName != test.Output.LastName {
				t.Errorf("%s: got: %v, expected: %v", test.Name, client, test.Output)
				return
			}
		})
	}
}

func TestUpdateClient(t *testing.T) {
	tests := []struct {
		Name    string
		Input   *models.UpdateClient
		Output  *models.Client
		WantErr bool
	}{
		{
			Name: "Case 1",
			Input: &models.UpdateClient{
				ClientId:   "9420a46e-ef3c-4fe9-85d1-90ae3a4dec2a",
				FirstName: "Gucci",
				LastName: "Safarov",
			},
			Output: &models.Client{
				ClientId:   "9420a46e-ef3c-4fe9-85d1-90ae3a4dec2a",
				FirstName: "Gucci",
			},
			WantErr: false,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			res, err := clientTestRepo.Update(context.Background(), test.Input)

			client, err := clientTestRepo.GetById(context.Background(), &models.ClientPrimaryKey{
				ClientId: test.Input.ClientId,
			})

			if client.FirstName != test.Output.FirstName {
				t.Errorf("%s: got: %v, expected: %v", test.Name, client, test.Output)
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

func TestDeleteClient(t *testing.T) {
	tests := []struct {
		Name    string
		Input   *models.ClientPrimaryKey
		WantErr bool
	}{
		{
			Name: "Case 1",
			Input: &models.ClientPrimaryKey{
				ClientId: "9420a46e-ef3c-4fe9-85d1-90ae3a4dec2a",
			},
			WantErr: false,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			res, err := clientTestRepo.Delete(context.Background(), &models.ClientPrimaryKey{
				ClientId: test.Input.ClientId,
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
