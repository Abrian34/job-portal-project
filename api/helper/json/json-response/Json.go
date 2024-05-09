package jsonresponse

import (
	"encoding/json"
	"errors"
	"job-portal-project/api/utils/constant"
	"net/http"
)

func WriteToResponseBody(writer http.ResponseWriter, response interface{}) error {
	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	err := encoder.Encode(response)
	if err != nil {
		return errors.New(constant.JsonError)
	}
	return nil
}
