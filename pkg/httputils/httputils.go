package httputils

import (
	"fmt"
	"net/http"
	"time"

	"github.com/unrolled/render"

	localErrors "github.com/jmvdr-iscte/ImageSearch/internal/errors"
)

type SuccessfulResponse struct {
	Data      interface{} `json:"data"`
	Timestamp time.Time   `json:"timestamp"`
}

type ErrorResponse struct {
	StatusCode int             `json:"statusCode"`
	Error      localErrors.Err `json:"error"`
	Timestamp  time.Time       `json:"timestamp"`
}

type Sender struct {
	Render *render.Render
}

func (s *Sender) JSON(w http.ResponseWriter, statusCode int, v interface{}) error {
	w.Header().Set("Content-Type", "applicaton/json")

	if statusCode > 200 && statusCode < 299 {
		response := SuccessfulResponse{Data: v, Timestamp: time.Now().UTC()}
		err := s.Render.JSON(w, statusCode, response)
		return err
	}

	//Add redirect if needed

	// client and server error
	if statusCode > 399 && statusCode < 599 {

		errResponse := localErrors.Err{Message: fmt.Sprint(v)}

		if errInfo, ok := v.(localErrors.Err); ok {
			errResponse.Message = errInfo.Error()
			errResponse.Data = errInfo.Data
		}

		response := ErrorResponse{StatusCode: statusCode, Error: errResponse, Timestamp: time.Now().UTC()}
		err := s.Render.JSON(w, statusCode, response)
		return err
	}

	err := s.Render.JSON(w, statusCode, v)
	if err != nil {
		return err
	}

	return nil
}
