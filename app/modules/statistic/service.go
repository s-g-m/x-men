package statistic

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return service{repository}
}

func (s service) GetStatistic() (response Response, err error) {
	nMutant, nHuman, err := s.repository.GetStatistic()
	ratio := float64(0)

	if err != nil {
		return
	}

	if nMutant != 0 && nHuman != 0 {
		ratio = float64(nMutant) / float64(nHuman)
	}

	response = Response{
		CountMutantDNA: nMutant,
		CountHumanDNA:  nHuman,
		Ratio:          ratio,
	}

	return
}
