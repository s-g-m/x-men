package testdata

type utilGenerateHash struct {
	Data        []string
	Hash        string
	Description string
}

func UtilGenerateHash() (data []utilGenerateHash) {
	data = append(data, utilGenerateHash{[]string{"Prueba", "Mercado Libre"}, "BF82AC906E12D19C89BA468E3E9DCC8A3BEC22A870EC866BF97F8BF29D0A1AD5", "normal hash test"})
	data = append(data, utilGenerateHash{[]string{""}, "E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855", "hash test without data"})
	return
}
