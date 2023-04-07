package order_service

import (
	"app/models"
	"fmt"
	"net/http"
	"sync"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/test-go/testify/assert"
)

func TestClient(t *testing.T) {
	c = 0
	wg := &sync.WaitGroup{}

	for i := 0; i < 10; i++ {

		wg.Add(1)
		go func() {
			defer wg.Done()
			id := CreateClient(t)
			deleteClient(t, id)
		}()

	}

	wg.Wait()

	fmt.Println("c: ", c)
}

func CreateClient(t *testing.T) string {
	response := &models.Client{}

	request := &models.CreateClient{
		FirstName: faker.FirstName(),
		LastName: faker.LastName(),
		PhoneNumber: faker.Phonenumber(),
	}

	resp, err := PerformRequest(http.MethodPost, "/client", request, response)

	assert.NoError(t, err)

	// a := object{} check whether the object is nil or not
	// 1 way
	assert.NotNil(t, resp)
	// another
	// b := object{}
	// reflect.DeepEqual(a, b)

	if resp != nil {
		assert.Equal(t, resp.StatusCode, 201)
	}

	fmt.Println(response)

	return response.ClientId
}

func updateClient(t *testing.T, id string) string {
	response := &models.Client{}
	request := &models.UpdateClient{
		FirstName: faker.FirstName(),
		LastName: faker.LastName(),
		PhoneNumber: faker.Phonenumber(),
	}

	resp, err := PerformRequest(http.MethodPut, "/client/"+id, request, response)

	assert.NoError(t, err)

	assert.NotNil(t, resp)

	if resp != nil {
		assert.Equal(t, resp.StatusCode, 200)
	}

	fmt.Println(resp)

	return response.ClientId
}

func deleteClient(t *testing.T, id string) string {

	resp, _ := PerformRequest(
		http.MethodDelete,
		fmt.Sprintf("/client/%s", id),
		nil,
		nil,
	)

	assert.NotNil(t, resp)

	if resp != nil {
		assert.Equal(t, resp.StatusCode, 204)
	}

	return ""
}
