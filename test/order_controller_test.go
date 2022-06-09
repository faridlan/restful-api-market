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

func TestCreateOrderSuccess(t *testing.T) {
	db := setupDBTest()
	router := setupRouter(db)

	requestBody := strings.NewReader(`
	{"id_address": "37cace3de3cb11ec989c0242ac110002",
	"products": [
		{
			"id_product":"7329649fe44611ecabc0a2623c2c5e50",
			"quantity": 2
		},
		{
			"id_product":"732cef69e44611ecabc0a2623c2c5e50",
			"quantity": 2
		}
	]
	}
	`)

	var bearer = "Bearer " + "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MiwiaWRfdXNlciI6IjE0YTE4NGE0ZTNjYjExZWM5ODljMDI0MmFjMTEwMDAyIiwidXNlcm5hbWUiOiJ1ZGluIiwiZW1haWwiOiJ1ZGluQGdtYWlsLmNvbSIsInJvbGVfaWQiOjIsInRva2VuIjoiWFZsQnpnYmFpQ01SQWpXd2hUSGMiLCJleHAiOjE2NTQ4NzQxMjh9.-JMM6Cbh_iuL5q4x7dUjUJp9uiBgavvS7b3Ajzc_8uo"
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/api/orders", requestBody)
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

func TestCreateOrderFailed(t *testing.T) {
	db := setupDBTest()
	router := setupRouter(db)

	requestBody := strings.NewReader(`
	{"products": [
		{
			"id_product":"7329649fe44611ecabc0a2623c2c5e50",
			"quantity": 2
		},
		{
			"id_product":"732cef69e44611ecabc0a2623c2c5e50",
			"quantity": 2
		}
	]
	}
	`)

	var bearer = "Bearer " + "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MiwiaWRfdXNlciI6IjE0YTE4NGE0ZTNjYjExZWM5ODljMDI0MmFjMTEwMDAyIiwidXNlcm5hbWUiOiJ1ZGluIiwiZW1haWwiOiJ1ZGluQGdtYWlsLmNvbSIsInJvbGVfaWQiOjIsInRva2VuIjoiWFZsQnpnYmFpQ01SQWpXd2hUSGMiLCJleHAiOjE2NTQ4NzQxMjh9.-JMM6Cbh_iuL5q4x7dUjUJp9uiBgavvS7b3Ajzc_8uo"
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/api/orders", requestBody)
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
