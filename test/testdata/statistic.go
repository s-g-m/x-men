package testdata

import (
	"x-men/app/modules/statistic"
)

/*
	Count (M = Mutant, H = Human, Number = Count )
*/
var (
	responseM0H0 = statistic.Response{CountMutantDNA: 0, CountHumanDNA: 0, Ratio: 0}
	responseM8H8 = statistic.Response{CountMutantDNA: 8, CountHumanDNA: 8, Ratio: 1}
	responseM8H0 = statistic.Response{CountMutantDNA: 8, CountHumanDNA: 0, Ratio: 0}
	responseM0H8 = statistic.Response{CountMutantDNA: 0, CountHumanDNA: 8, Ratio: 0}
	responseM8H2 = statistic.Response{CountMutantDNA: 8, CountHumanDNA: 2, Ratio: 4}
	responseM8H5 = statistic.Response{CountMutantDNA: 8, CountHumanDNA: 5, Ratio: 1.6}
	responseM2H8 = statistic.Response{CountMutantDNA: 2, CountHumanDNA: 8, Ratio: 0.25}
)

type controllerGetStatistic struct {
	Response    statistic.Response
	Status      int
	Error       error
	Description string
}

func ControllerGetStatistic() (data []controllerGetStatistic) {
	data = append(data, controllerGetStatistic{responseM0H0, 200, nil, "service response succeed"})
	data = append(data, controllerGetStatistic{responseM0H0, 500, err, "service response failed"})
	return
}

type serviceGetStatistic struct {
	Response    statistic.Response
	Error       error
	Description string
}

func ServiceGetStatistic() (data []serviceGetStatistic) {
	data = append(data, serviceGetStatistic{statistic.Response{}, err, "repository response failed"})
	data = append(data, serviceGetStatistic{responseM0H0, nil, "ratio must be 0"})
	data = append(data, serviceGetStatistic{responseM8H8, nil, "ratio must be 1"})
	data = append(data, serviceGetStatistic{responseM8H0, nil, "ratio must be 0"})
	data = append(data, serviceGetStatistic{responseM0H8, nil, "ratio must be 0"})
	data = append(data, serviceGetStatistic{responseM8H2, nil, "ratio must be 4"})
	data = append(data, serviceGetStatistic{responseM8H5, nil, "ratio must be 1.6"})
	data = append(data, serviceGetStatistic{responseM2H8, nil, "ratio must be 0.25"})
	return
}
