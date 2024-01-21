package handlers

import (
	"net/http"

	"github.com/jmvdr-iscte/ImageSearch/pkg/logger"
	"github.com/sirupsen/logrus"
)

func (h *Handlers) GetHealth(w http.ResponseWriter, r *http.Request) {

	healthResponse := map[string]string{
		"service": "image_search",
		"status":  "Ok",
	}

	err := h.Sender.JSON(w, http.StatusOK, healthResponse)
	if err != nil {
		logger.OutputLog.WithFields(logrus.Fields{
			"err": err.Error(),
		}).Fatal("Error when getting service health")

		panic(err)
	}
}
