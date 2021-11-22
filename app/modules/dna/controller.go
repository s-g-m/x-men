package dna

import (
	"encoding/json"
	"net/http"
)

type controller struct {
	service Service
}

func NewController(service Service) Controller {
	return controller{service: service}
}

func (c controller) IsMutant(httpAdapter HttpAdapter) {
	request := Request{}
	body := httpAdapter.GetBody()
	err := json.Unmarshal(body, &request)

	if err != nil {
		httpAdapter.SendResponse(http.StatusBadRequest, err.Error())
		return
	}

	isMutant, err := c.service.IsMutant(request.Dna)
	if err != nil {
		httpAdapter.SendResponse(http.StatusBadRequest, err.Error())
		return
	}
	if !isMutant {
		httpAdapter.SendResponse(http.StatusForbidden, "no is mutant")
		return
	}

	httpAdapter.SendResponse(http.StatusOK, "is mutant")
}
