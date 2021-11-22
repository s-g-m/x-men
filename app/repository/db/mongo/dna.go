package mongo

import (
	"context"
	"x-men/app/repository/db/entity"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const dnaCollection = "dna"

func (m Mongo) CreateDNA(dna entity.Dna) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeOut)
	defer cancel()

	_, err = m.db.Collection(dnaCollection).InsertOne(ctx, dna)
	return
}

func (m Mongo) GetStatistic() (mutant int, human int, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeOut)
	defer cancel()

	group := bson.D{{"$group", bson.D{{"_id", "$is_mutant"}, {"count", bson.D{{"$sum", 1}}}}}}

	cursor, err := m.db.Collection(dnaCollection).Aggregate(ctx, mongo.Pipeline{group})
	if err != nil {
		return
	}
	defer cursor.Close(ctx)

	results := []Statistic{}
	err = cursor.All(ctx, &results)
	if err != nil {
		return
	}

	for _, result := range results {
		if result.Id {
			mutant = result.Count
		} else {
			human = result.Count
		}
	}

	return
}
