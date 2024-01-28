package handlers

import (
	"net/http"

	"github.com/jmvdr-iscte/ImageSearch/pkg/logger"
	"github.com/sirupsen/logrus"
)

func (h *Handlers) GetTest(w http.ResponseWriter, r *http.Request) {

	response := map[string]string{
		"status": "Ok",
	}

	err := h.Sender.JSON(w, http.StatusOK, response)
	if err != nil {
		logger.OutputLog.WithFields(logrus.Fields{
			"err": err.Error(),
		}).Fatal("Error when getting test")

		panic(err)
	}
}
