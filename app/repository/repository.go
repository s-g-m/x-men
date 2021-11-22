package repository

import (
	"x-men/app/modules/dna"
	"x-men/app/modules/statistic"
	"x-men/app/repository/db/mongo"
	"x-men/app/repository/db/redis"
	"x-men/config"
)

type Repository interface {
	statistic.Repository
	dna.Repository
}

type repository struct {
	*mongo.Mongo
	*redis.RedisPub
}

func NewRepository() Repository {
	mongoDB := mongo.NewMongo(config.DbUri(), config.DbName())
	redisDB := redis.NewRedisPub(config.RedisUri(), config.RedisKey())
	return repository{Mongo: &mongoDB, RedisPub: &redisDB}
}
