package test

import (
	"database/sql"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/faridlan/restful-api-market/app"
	"github.com/faridlan/restful-api-market/controller"
	"github.com/faridlan/restful-api-market/helper"
	"github.com/faridlan/restful-api-market/middleware"
	"github.com/faridlan/restful-api-market/repository"
	"github.com/faridlan/restful-api-market/service"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

func setupDBTest() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/olshop_test")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(50)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}

func setupRouter(db *sql.DB) http.Handler {

	validate := validator.New()

	uuidRepository := repository.NewUuidRepository()

	//Status Order
	statusOrderRepository := repository.NewStatusOrderRepository()
	statusOrderService := service.NewStatusOrderService(statusOrderRepository, uuidRepository, db, validate)
	statusOrderContoller := controller.NewStatusOrderController(statusOrderService)

	//Roles
	roleRepository := repository.NewRoleRepository()
	roleService := service.NewRoleService(roleRepository, uuidRepository, db, validate)
	roleController := controller.NewRoleController(roleService)

	//Category
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, uuidRepository, db, validate)
	CategoryController := controller.NewCategoryController(categoryService)

	//Admin
	//Product Feature
	productRepository := repository.NewProdcutRepository()
	productService := service.NewProductServie(productRepository, categoryRepository, uuidRepository, db, validate)
	productController := controller.NewProductController(productService)

	//auth
	userRepository := repository.NewUserRepository()
	blacklistRepository := repository.NewBlacklistRepository()
	userService := service.NewAuthService(userRepository, blacklistRepository, roleRepository, uuidRepository, db, validate)
	authController := controller.NewAuthController(userService)

	//product detail
	cartRepository := repository.NewCartRepository()

	//shopping cart
	shoppingCartService := service.NewShoppingCartService(cartRepository, db, validate)
	shoppingCartController := controller.NewShoppingCartController(shoppingCartService)

	//address
	addressRepository := repository.NewAddressRepository()
	addressSerice := service.NewAddressService(addressRepository, uuidRepository, db, validate)
	addressController := controller.NewAddressController(addressSerice)

	//shipping address
	orderRepository := repository.NewOrderRepository()
	orderDetailRepository := repository.NewOrderDetailRepository()
	shippingAddressService := service.NewShippingAddressService(productRepository, orderRepository, orderDetailRepository, cartRepository, addressRepository, statusOrderRepository, uuidRepository, db, validate)
	shippingAddressController := controller.NewShippingAddressController(shippingAddressService)

	//seeder
	seederService := service.NewSeedService(addressRepository, orderDetailRepository, orderRepository, productRepository, userRepository, db)
	seederController := controller.NewSeederController(seederService, userService, addressSerice, productService, shippingAddressService)

	controller := app.ControllerRouter{
		AddressController:         addressController,
		AuthController:            authController,
		CategoryController:        CategoryController,
		ProductController:         productController,
		ShippingAddressController: shippingAddressController,
		ShoppingCartController:    shoppingCartController,
		StatusOrderController:     statusOrderContoller,
		RoleController:            roleController,
		SeederController:          seederController,
	}

	router := app.NewRouter(controller)

	return middleware.NewAuthMiddleware(router, blacklistRepository, db)

}

func truncateUser(db *sql.DB) {
	db.Exec("TRUNCATE users")
}

func TestCreateUserSuccess(t *testing.T) {
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

func TestCreateUserFailed(t *testing.T) {
	db := setupDBTest()
	truncateUser(db)
	router := setupRouter(db)

	requestBody := strings.NewReader(`
	{"username": "userAB",
	"password": "userA1234"}
	`)

	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/api/register", requestBody)
	request.Header.Add("Content-Type", "application/json")

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

func TestLoginSuccess(t *testing.T) {
	db := setupDBTest()
	truncateUser(db)
	router := setupRouter(db)

	requestBody := strings.NewReader(`
	{"username": "userAB",
	"password": "userA1234"}
	`)

	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/api/login", requestBody)
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

func TestLoginFailed(t *testing.T) {
	db := setupDBTest()
	truncateUser(db)
	router := setupRouter(db)

	requestBody := strings.NewReader(`
	{"username": "userABZZ",
	"password": "userA1234"}
	`)

	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/api/login", requestBody)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	response := recorder.Result()
	assert.Equal(t, 401, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 401, int(responseBody["code"].(float64)))
	assert.Equal(t, "UNAUTHORIZED", responseBody["status"])
}
