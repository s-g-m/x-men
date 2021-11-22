package subscriber

import (
	"x-men/app/repository/db/mongo"
	"x-men/app/repository/db/redis"
	"x-men/config"
	"x-men/util/logs"
)

func Start() {
	repository := mongo.NewMongo(config.DbUri(), config.DbName())
	redisSub := redis.NewRedisSub(repository, config.RedisUri(), config.RedisKey())
	go redisSub.SaveDNA()
	logs.Info("subscribers Ready")
}
