package redis

import (
	"context"
	"encoding/json"
	"x-men/app/repository/db/entity"
	"x-men/util/constant"
	"x-men/util/logs"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

type RedisPub struct {
	redis *redis.Client
}

func NewRedisPub(uri, key string) RedisPub {
	redis := redis.NewClient(&redis.Options{
		Addr: uri,
	})
	logs.Info("connected RedisPub")
	return RedisPub{redis: redis}
}

func (r RedisPub) SaveDNA(dna []string, isMutant bool) (err error) {
	newDna := entity.Dna{
		Dna:      dna,
		IsMutant: isMutant,
	}

	payload, err := json.Marshal(newDna)
	if err != nil {
		return
	}

	err = r.redis.Publish(ctx, constant.ChannelNameSaveDNA, payload).Err()
	if err != nil {
		logs.Error("RedisPub.SaveDNA", err.Error())
	}
	return
}
