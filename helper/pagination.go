package helper

import (
	"net/http"
	"strconv"

	"github.com/faridlan/restful-api-market/model/domain"
)

func Pagination(request *http.Request) domain.Pagination {
	queryParam := request.URL.Query()
	var pagintaion domain.Pagination
	if len(queryParam) == 0 {
		pagintaion.Page = 1
		pagintaion.Limit = 10
		return pagintaion
	} else {
		page, err := strconv.Atoi(queryParam.Get("page"))
		PanicIfError(err)
		limit, err := strconv.Atoi(queryParam.Get("limit"))
		PanicIfError(err)

		pagintaion.Page = (page - 1) * limit
		pagintaion.Limit = limit
		return pagintaion
	}
}
