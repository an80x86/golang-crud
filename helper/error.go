package helper

import (
	"golang-crud/data/response"
	"net/http"
)

func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}

func PanicWebIfError(writer http.ResponseWriter, err error) error {
	if err != nil {

		webResponse := response.WebResponse{
			Code:   400,
			Status: "Error",
			Data:   err.Error(),
		}

		WriteResponseBody(writer, webResponse)
	}

	return err
}
