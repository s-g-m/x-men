package statistic

type Response struct {
	CountMutantDNA int     `json:"count_mutant_dna,omitempty"`
	CountHumanDNA  int     `json:"count_human_dna,omitempty"`
	Ratio          float64 `json:"ratio,omitempty"`
}
