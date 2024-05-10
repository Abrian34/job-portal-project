package payloads

import (
	jsonresponse "job-portal-project/api/helper/json/json-response"
	"net/http"
)

func ResponseToken(writer http.ResponseWriter, data interface{}, message string, statusCode int) error {
	res := Response{
		StatusCode: statusCode,
		Message:    message,
		Data:       data,
	}

	err := jsonresponse.WriteToResponseBody(writer, res)
	if err != nil {
		return err
	}
	return nil
}
