package statistic

import "net/http"

type controller struct {
	service Service
}

func NewController(service Service) Controller {
	return controller{service: service}
}

func (c controller) GetStatistic(httpAdapter HttpAdapter) {
	response, err := c.service.GetStatistic()

	if err != nil {
		httpAdapter.SendResponse(http.StatusInternalServerError, err.Error())
		return
	}

	httpAdapter.SendResponse(http.StatusOK, response)
}
