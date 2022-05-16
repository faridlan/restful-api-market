package exception

import (
	"fmt"
	"net/http"

	"github.com/faridlan/restful-api-market/helper"
	"github.com/faridlan/restful-api-market/model/web"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
)

func ExceptionError(writer http.ResponseWriter, request *http.Request, err interface{}) {

	if unauthError(writer, request, err) {
		return
	}

	if notFoundError(writer, request, err) {
		return
	}

	if validationError(writer, request, err) {
		return
	}

	internalServerError(writer, request, err)

}

func internalServerError(writer http.ResponseWriter, request *http.Request, err interface{}) {
	writer.Header().Add("content-type", "application/json")
	writer.WriteHeader(http.StatusInternalServerError)

	webResponse := web.WebResponse{
		Code:   http.StatusInternalServerError,
		Status: "INTERNAL SERVER ERROR",
		Data:   err,
	}

	helper.WriteToResponseBody(writer, webResponse)
	fmt.Println(err)
}

func unauthError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {

	exception, ok := err.(UnauthError)

	if ok {
		writer.Header().Add("content-type", "application/json")
		writer.WriteHeader(http.StatusUnauthorized)

		webResponse := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
			Data:   exception.Error,
		}

		helper.WriteToResponseBody(writer, webResponse)
		return true
	} else {
		return false
	}

}

func notFoundError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {

	exception, ok := err.(NotFoundError)

	if ok {
		writer.Header().Add("content-type", "application/json")
		writer.WriteHeader(http.StatusNotFound)

		webResponse := web.WebResponse{
			Code:   http.StatusNotFound,
			Status: "NOT FOUND",
			Data:   exception.Error,
		}

		helper.WriteToResponseBody(writer, webResponse)
		return true
	} else {
		return false
	}

}

// func translateError(err error, trans ut.Translator) (errs []error) {
// 	if err == nil {
// 		return nil
// 	}
// 	validatorErrs := err.(validator.ValidationErrors)
// 	for _, e := range validatorErrs {
// 		translatedErr := fmt.Errorf(e.Translate(trans))
// 		errs = append(errs, translatedErr)
// 	}
// 	return errs
// }

func validationError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(validator.ValidationErrors)

	if ok {
		writer.Header().Add("content-type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)

		validate := validator.New()
		english := en.New()
		uni := ut.New(english, english)
		trans, _ := uni.GetTranslator("en")
		_ = enTranslations.RegisterDefaultTranslations(validate, trans)

		var errs []interface{}

		for _, v := range exception {
			translateErr := v.Translate(trans)
			errs = append(errs, translateErr)
		}

		webResponse := web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   errs,
		}

		helper.WriteToResponseBody(writer, webResponse)

		return true
	} else {
		return false
	}
}
