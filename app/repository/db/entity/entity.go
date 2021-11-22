package entity

type Dna struct {
	Id        string   `bson:"_id"`
	Dna       []string `bson:"dna"`
	IsMutant  bool     `bson:"is_mutant"`
	CreatedAt string   `bson:"createdAt"`
}
