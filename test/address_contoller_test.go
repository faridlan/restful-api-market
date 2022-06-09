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

func TestCreateAddressSuccess(t *testing.T) {
	db := setupDBTest()
	router := setupRouter(db)

	requestBody := strings.NewReader(`
	{"name": "UdinHAHA",
	"handphone_number": "0897654332",
	"street": "Jl Udin",
	"district": "Udinhiang",
	"postal_code": 46151,
	"comment": "LALALALA"}
	`)

	var bearer = "Bearer " + "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MiwiaWRfdXNlciI6IjE0YTE4NGE0ZTNjYjExZWM5ODljMDI0MmFjMTEwMDAyIiwidXNlcm5hbWUiOiJ1ZGluIiwiZW1haWwiOiJ1ZGluQGdtYWlsLmNvbSIsInJvbGVfaWQiOjIsInRva2VuIjoiWFZsQnpnYmFpQ01SQWpXd2hUSGMiLCJleHAiOjE2NTQ4NzQxMjh9.-JMM6Cbh_iuL5q4x7dUjUJp9uiBgavvS7b3Ajzc_8uo"
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/api/addresses", requestBody)
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
	assert.Equal(t, "UdinHAHA", responseBody["data"].(map[string]interface{})["name"])
	assert.Equal(t, "0897654332", responseBody["data"].(map[string]interface{})["handphone_number"])
	assert.Equal(t, "Jl Udin", responseBody["data"].(map[string]interface{})["street"])
	assert.Equal(t, "Udinhiang", responseBody["data"].(map[string]interface{})["district"])
	assert.Equal(t, 46151, int(responseBody["data"].(map[string]interface{})["postal_code"].(float64)))
	assert.Equal(t, "LALALALA", responseBody["data"].(map[string]interface{})["comment"])
}

func TestCreateAddressFailed(t *testing.T) {
	db := setupDBTest()
	router := setupRouter(db)

	requestBody := strings.NewReader(`
	{"name": "UdinHAHA",
	"handphone_number": "0897654332",
	"street": "Jl Udin",
	"district": "Udinhiang",
	"postal_code": 46151
	}
	`)

	var bearer = "Bearer " + "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MiwiaWRfdXNlciI6IjE0YTE4NGE0ZTNjYjExZWM5ODljMDI0MmFjMTEwMDAyIiwidXNlcm5hbWUiOiJ1ZGluIiwiZW1haWwiOiJ1ZGluQGdtYWlsLmNvbSIsInJvbGVfaWQiOjIsInRva2VuIjoiWFZsQnpnYmFpQ01SQWpXd2hUSGMiLCJleHAiOjE2NTQ4NzQxMjh9.-JMM6Cbh_iuL5q4x7dUjUJp9uiBgavvS7b3Ajzc_8uo"
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/api/addresses", requestBody)
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
