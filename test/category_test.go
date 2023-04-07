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

var c int64

func TestActor(t *testing.T) {
	c = 0
	wg := &sync.WaitGroup{}

	for i := 0; i < 10; i++ {

		wg.Add(1)
		go func() {
			defer wg.Done()
			id := CreateCategory(t)
			deleteCategory(t, id)
		}()

	}

	wg.Wait()

	fmt.Println("c: ", c)
}

func CreateCategory(t *testing.T) string {
	response := &models.Category{}

	request := &models.CreateCategory{
		CategoryName: faker.FirstName(),
	}

	resp, err := PerformRequest(http.MethodPost, "/category", request, response)

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

	return response.CategoryId
}

func updateCategory(t *testing.T, id string) string {
	response := &models.Category{}
	request := &models.UpdateCategory{
		CategoryName: faker.FirstName(),
	}

	resp, err := PerformRequest(http.MethodPut, "/category/"+id, request, response)

	assert.NoError(t, err)

	assert.NotNil(t, resp)

	if resp != nil {
		assert.Equal(t, resp.StatusCode, 200)
	}

	fmt.Println(resp)

	return response.CategoryId
}

func deleteCategory(t *testing.T, id string) string {

	resp, _ := PerformRequest(
		http.MethodDelete,
		fmt.Sprintf("/category/%s", id),
		nil,
		nil,
	)

	assert.NotNil(t, resp)

	if resp != nil {
		assert.Equal(t, resp.StatusCode, 204)
	}

	return ""
}
