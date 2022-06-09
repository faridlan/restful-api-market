package test

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateProductSuccess(t *testing.T) {
	db := setupDBTest()
	router := setupRouter(db)

	requestBody := strings.NewReader(`
	{"product_name": "ProductA",
	"id_category": "56e194e2e26711ec93820242ac110002",
	"price": 9999,
	"quantity": 99,
	"image_url": "https://olshop.sgp1.digitaloceanspaces.com/products/LprucjiOgj.png"
	}
	`)

	var bearer = "Bearer " + "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6NywiaWRfdXNlciI6IjE5Mzk0ZjU3ZTNjYjExZWM5ODljMDI0MmFjMTEwMDAyIiwidXNlcm5hbWUiOiJhZG1pbiIsImVtYWlsIjoiYWRtaW5AYWRtaW4uY29tIiwicm9sZV9pZCI6MSwidG9rZW4iOiJYVmxCemdiYWlDTVJBald3aFRIYyIsImV4cCI6MTY1NDg3OTk1MH0.P5HVlwSnmyk_4J-fwqrOWybcoQXZa6S1QCa97ZxfNkE"
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/api/products", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", bearer)

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
}

func TestCreateProductFailed(t *testing.T) {
	db := setupDBTest()
	router := setupRouter(db)

	requestBody := strings.NewReader(`
	{"id_category": "56e194e2e26711ec93820242ac110002",
	"image_url": "https://olshop.sgp1.digitaloceanspaces.com/products/LprucjiOgj.png"
	}
	`)

	var bearer = "Bearer " + "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6NywiaWRfdXNlciI6IjE5Mzk0ZjU3ZTNjYjExZWM5ODljMDI0MmFjMTEwMDAyIiwidXNlcm5hbWUiOiJhZG1pbiIsImVtYWlsIjoiYWRtaW5AYWRtaW4uY29tIiwicm9sZV9pZCI6MSwidG9rZW4iOiJYVmxCemdiYWlDTVJBald3aFRIYyIsImV4cCI6MTY1NDg3OTk1MH0.P5HVlwSnmyk_4J-fwqrOWybcoQXZa6S1QCa97ZxfNkE"
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/api/products", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", bearer)

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	response := recorder.Result()
	assert.Equal(t, 400, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 400, int(responseBody["code"].(float64)))
	assert.Equal(t, "BAD REQUEST", responseBody["status"])
}
