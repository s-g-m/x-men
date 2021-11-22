package testdata

import (
	"errors"
	"x-men/util/constant"
)

/*
	sequences (H = horizontal, V = Vertical, D = Diagonal, C = Cross Diagonal, N = No Sequences )
	characteristic (I = With interception, S = Separated, F = Followed)
*/
var (
	BodyIsMutant = `{"dna":["ATGCGA","CAGTGC","TTATGT","AGAAGG","CCCCTA","TCACTG"]}`

	DnaNNN = []string{"ACACACAC", "TGTGTGTG", "ACACACAC", "TGTGTGTG", "ACACACAC", "TGTGTGTG", "ACACACAC", "TGTGTGTG"}
	dnaHNN = []string{"AAACAAAA", "TGTGTGTG", "ACACACAC", "TGTGTGTG", "ACACACAC", "TGTGTGTG", "ACACACAC", "TGTGTGTG"}
	dnaVNN = []string{"ACACACTC", "TGTGTGTG", "ACACACTC", "TGTGTGTG", "ACACACAC", "TGTGTGTG", "ACACACTC", "TGTGTGTG"}
	dnaDNN = []string{"GCACACAC", "TGTGTGTG", "ACGCACAC", "TGTGTGTG", "ACACACAC", "TGTGTGTG", "ACACACGC", "TGTGTGTG"}
	dnaCNN = []string{"ACACACAC", "TGTGTGCG", "ACACACAC", "TGTGTGTG", "ACACACAC", "TGCGTGTG", "ACACACAC", "CGTGTGTG"}

	dnaHHI = []string{"ACACACAC", "TGTGTGTG", "ACACACAC", "TGTGTGTG", "ACACACAC", "TGTGTGTG", "AAAAAAAC", "TGTGTGTG"}
	dnaVVI = []string{"ACACACAC", "TGTGTGTG", "ACACACTC", "TGTGTGTG", "ACACACTC", "TGTGTGTG", "ACACACTC", "TGTGTGTG"}
	dnaDDI = []string{"ACACACAC", "TGTGTGTG", "ACGCACAC", "TGTGTGTG", "ACACGCAC", "TGTGTGTG", "ACACACGC", "TGTGTGTG"}
	dnaCCI = []string{"ACACACAC", "TGTGTGCG", "ACACACAC", "TGTGCGTG", "ACACACAC", "TGCGTGTG", "ACACACAC", "TGTGTGTG"}
	dnaHVI = []string{"ACACACAC", "TGTGTGAG", "ACACAAAA", "TGTGTGAG", "ACACACAC", "TGTGTGTG", "ACACACAC", "TGTGTGTG"}
	dnaHDI = []string{"ATACACAC", "TTTTTGTG", "ACATACAC", "TGTGTGTG", "ACACACAC", "TGTGTGTG", "ACACACAC", "TGTGTGTG"}
	dnaHCI = []string{"ACACACAC", "TGTGTGTG", "ACGCACAC", "GGGGTGTG", "GCACACAC", "TGTGTGTG", "ACACACAC", "TGTGTGTG"}
	dnaDCI = []string{"ACACACAC", "TGTGTGTG", "ACACACAC", "TGTGTGTG", "ACACACAC", "TGCGCGTG", "ACACACAC", "TGCGCGTG"}

	dnaHHS = []string{"AAAAACAC", "TGTGTGTG", "ACACACAC", "TGTGTGTG", "ACACACAC", "TGTGTGTG", "ACACAAAA", "TGTGTGTG"}
	dnaVVS = []string{"ACTCACAC", "TGTGTGTG", "ACTCACAC", "TGTGTGTG", "ACACACTC", "TGTGTGTG", "ACACACTC", "TGTGTGTG"}
	dnaDDS = []string{"ACACGCAC", "TGTGTGTG", "ACACACGC", "TGTGTGTG", "GCACACAC", "TGTGTGTG", "ACGCACAC", "TGTGTGTG"}
	dnaCCS = []string{"ACACACAC", "TGCGTGTG", "ACACACAC", "CGTGTGTG", "ACACACAC", "TGTGTGCG", "ACACACAC", "TGTGCGTG"}
	dnaHVS = []string{"ACACACAC", "TGTGTGTG", "ACACAAAA", "TGTGTGTG", "ACACACAC", "TGAGTGTG", "ACACACAC", "TGAGTGTG"}
	dnaHDS = []string{"ACACACAC", "TGTGTTTT", "ACACACAC", "TGTGTGTG", "ACACATAC", "TGTGTGTG", "ACACACAT", "TGTGTGTG"}
	dnaHCS = []string{"ACACGCAC", "TGTGTGTG", "ACGCACAC", "TGTGTGTG", "ACACACAC", "GGGGTGTG", "ACACACAC", "TGTGTGTG"}
	dnaDCS = []string{"ACACACAC", "TGCGTGTG", "ACACACAC", "CGTGTGTG", "ACACACAC", "TGTGCGTG", "ACACACAC", "TGTGTGCG"}

	dnaHHF = []string{"AAAAAAAA", "TGTGTGTG", "ACACACAC", "TGTGTGTG", "ACACACAC", "TGTGTGTG", "ACACACAC", "TGTGTGTG"}
	dnaVVF = []string{"ACACACTC", "TGTGTGTG", "ACACACTC", "TGTGTGTG", "ACACACTC", "TGTGTGTG", "ACACACTC", "TGTGTGTG"}
	dnaDDF = []string{"GCACACAC", "TGTGTGTG", "ACGCACAC", "TGTGTGTG", "ACACGCAC", "TGTGTGTG", "ACACACGC", "TGTGTGTG"}
	dnaCCF = []string{"ACACACAC", "TGTGTGCG", "ACACACAC", "TGTGCGTG", "ACACACAC", "TGCGTGTG", "ACACACAC", "CGTGTGTG"}
	dnaHVF = []string{"AAAAACAC", "TGTGAGTG", "ACACACAC", "TGTGAGTG", "ACACCCAC", "TGTGTGTG", "ACACACAC", "TGTGTGTG"}
	dnaHDF = []string{"ACACACAC", "TTTTTGTG", "ACACATAC", "TGTGTGTG", "ACACACAT", "TGTGTGTG", "ACACACAC", "TGTGTGTG"}
	dnaHCF = []string{"ACACACAC", "TGTGGGGG", "ACGCACAC", "TGTGTGTG", "GCACACAC", "TGTGTGTG", "ACACACAC", "TGTGTGTG"}
	dnaDCF = []string{"ACACACAC", "TGTGTGTG", "ACACACAC", "TGCGTGTG", "ACACACAC", "TGTGCGCG", "ACACACAC", "TGTGCGTG"}

	dnaUnreadable = []string{"LTGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"}
	dnaIncomplete = []string{"TGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"}
	dnaRectangle  = []string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA"}
	dnaShort      = []string{"ATC", "ATC", "ATC"}
)

type controllerIsMutant struct {
	BodyIsOk    bool
	IsMutant    bool
	Status      int
	Error       error
	Description string
}

func ControllerIsMutant() (data []controllerIsMutant) {
	data = append(data, controllerIsMutant{false, false, 400, err, "request format is invalid"})
	data = append(data, controllerIsMutant{true, false, 400, err, "service response is error"})
	data = append(data, controllerIsMutant{true, false, 403, nil, "service responds is not mutant"})
	data = append(data, controllerIsMutant{true, true, 200, nil, "service responds is mutant"})
	return
}

type serviceIsMutant struct {
	Dna         []string
	IsMutant    bool
	Error       error
	Description string
}

func ServiceIsMutant() (data []serviceIsMutant) {

	data = append(data, serviceIsMutant{dnaHNN, false, nil, "H sequence without any other 4-letter sequence"})
	data = append(data, serviceIsMutant{dnaVNN, false, nil, "V sequence without any other 4-letter sequence"})
	data = append(data, serviceIsMutant{dnaDNN, false, nil, "D sequence without any other 4-letter sequence"})
	data = append(data, serviceIsMutant{dnaCNN, false, nil, "C sequence without any other 4-letter sequence"})

	data = append(data, serviceIsMutant{dnaHHI, false, nil, "H sequence intercepted by another H sequence"})
	data = append(data, serviceIsMutant{dnaVVI, false, nil, "V sequence intercepted by another V sequence"})
	data = append(data, serviceIsMutant{dnaDDI, false, nil, "D sequence intercepted by another D sequence"})
	data = append(data, serviceIsMutant{dnaCCI, false, nil, "C sequence intercepted by another C sequence"})
	data = append(data, serviceIsMutant{dnaHVI, true, nil, "H sequence intercepted by another V sequence"})
	data = append(data, serviceIsMutant{dnaHDI, true, nil, "H sequence intercepted by another D sequence"})
	data = append(data, serviceIsMutant{dnaHCI, true, nil, "H sequence intercepted by another C sequence"})
	data = append(data, serviceIsMutant{dnaDCI, true, nil, "D sequence intercepted by another C sequence"})

	data = append(data, serviceIsMutant{dnaHHS, true, nil, "H sequence separated from another H sequence"})
	data = append(data, serviceIsMutant{dnaVVS, true, nil, "V sequence separated from another V sequence"})
	data = append(data, serviceIsMutant{dnaDDS, true, nil, "D sequence separated from another D sequence"})
	data = append(data, serviceIsMutant{dnaCCS, true, nil, "C sequence separated from another C sequence"})
	data = append(data, serviceIsMutant{dnaHVS, true, nil, "H sequence separated from another V sequence"})
	data = append(data, serviceIsMutant{dnaHDS, true, nil, "H sequence separated from another D sequence"})
	data = append(data, serviceIsMutant{dnaHCS, true, nil, "H sequence separated from another C sequence"})
	data = append(data, serviceIsMutant{dnaDCS, true, nil, "D sequence separated from another C sequence"})

	data = append(data, serviceIsMutant{dnaHHF, true, nil, "H sequence followed by another H sequence"})
	data = append(data, serviceIsMutant{dnaVVF, true, nil, "V sequence followed by another V sequence"})
	data = append(data, serviceIsMutant{dnaDDF, true, nil, "D sequence followed by another D sequence"})
	data = append(data, serviceIsMutant{dnaCCF, true, nil, "C sequence followed by another C sequence"})
	data = append(data, serviceIsMutant{dnaHVF, true, nil, "H sequence followed by another V sequence"})
	data = append(data, serviceIsMutant{dnaHDF, true, nil, "H sequence followed by another D sequence"})
	data = append(data, serviceIsMutant{dnaHCF, true, nil, "H sequence followed by another C sequence"})
	data = append(data, serviceIsMutant{dnaDCF, true, nil, "D sequence followed by another C sequence"})

	data = append(data, serviceIsMutant{dnaUnreadable, false, errors.New(constant.ErrorRegexpDNA), "DNA is unreadable"})
	data = append(data, serviceIsMutant{dnaIncomplete, false, errors.New(constant.ErrorSquareDNA), "DNA is incomplete"})
	data = append(data, serviceIsMutant{dnaRectangle, false, errors.New(constant.ErrorSquareDNA), "DNA is not square"})
	data = append(data, serviceIsMutant{dnaShort, false, errors.New(constant.ErrorSizeDNA), "DNA is short"})
	return
}

type modelReadDna struct {
	Position            int
	CrossDiagonalInvert string
	CrossDiagonal       string
	DiagonalInvert      string
	Diagonal            string
	Column              string
	Row                 string
}

func ModelReadDna() (data []modelReadDna) {
	data = append(data, modelReadDna{0, "TCTCTCTC", "A", "T", "AGAGAGAG", "ATATATAT", "ACACACAC"})
	data = append(data, modelReadDna{1, "GAGAGAG", "CT", "GA", "CTCTCTC", "CGCGCGCG", "TGTGTGTG"})
	data = append(data, modelReadDna{2, "TCTCTC", "AGA", "TCT", "AGAGAG", "ATATATAT", "ACACACAC"})
	data = append(data, modelReadDna{3, "GAGAG", "CTCT", "GAGA", "CTCTC", "CGCGCGCG", "TGTGTGTG"})
	data = append(data, modelReadDna{4, "TCTC", "AGAGA", "TCTCT", "AGAG", "ATATATAT", "ACACACAC"})
	data = append(data, modelReadDna{5, "GAG", "CTCTCT", "GAGAGA", "CTC", "CGCGCGCG", "TGTGTGTG"})
	data = append(data, modelReadDna{6, "TC", "AGAGAGA", "TCTCTCT", "AG", "ATATATAT", "ACACACAC"})
	data = append(data, modelReadDna{7, "G", "CTCTCTCT", "GAGAGAGA", "C", "CGCGCGCG", "TGTGTGTG"})
	return
}
