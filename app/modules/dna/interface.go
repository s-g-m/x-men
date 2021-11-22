package dna

type Controller interface {
	IsMutant(httpAdapter HttpAdapter)
}

type Service interface {
	IsMutant(dna []string) (isMutant bool, err error)
}

type HttpAdapter interface {
	SendResponse(status int, response interface{})
	GetBody() (body []byte)
}

type Repository interface {
	SaveDNA(dna []string, isMutant bool) (err error)
}
