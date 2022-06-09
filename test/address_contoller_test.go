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
	truncateUser(db)
	router := setupRouter(db)

	requestBody := strings.NewReader(`
	{"username": "userAB",
	"email": "userAB@mail.com",
	"password": "userA1234"}
	`)

	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/api/register", requestBody)
	request.Header.Add("Content-Type", "application/json")

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
