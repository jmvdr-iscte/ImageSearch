package handlers

import (
	"github.com/go-playground/validator/v10"
	"github.com/jmvdr-iscte/ImageSearch/pkg/httputils"
)

type Handlers struct {
	Sender *httputils.Sender
	//Add storage
}

// validation services for the handlers
var Validate *validator.Validate = validator.New(validator.WithRequiredStructEnabled())
