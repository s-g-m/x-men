package statistic

type Controller interface {
	GetStatistic(httpAdapter HttpAdapter)
}

type Service interface {
	GetStatistic() (response Response, err error)
}

type HttpAdapter interface {
	SendResponse(status int, response interface{})
}

type Repository interface {
	GetStatistic() (mutant int, human int, err error)
}
