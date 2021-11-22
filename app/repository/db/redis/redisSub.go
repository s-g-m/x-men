package redis

import (
	"encoding/json"
	"time"
	"x-men/app/repository/db/entity"
	"x-men/util/constant"
	"x-men/util/logs"
	"x-men/util/security"

	"github.com/go-redis/redis/v8"
)

type Repository interface {
	CreateDNA(dna entity.Dna) (err error)
}

type RedisSub struct {
	repository Repository
	redis      *redis.Client
}

func NewRedisSub(repository Repository, uri, key string) RedisSub {
	redis := redis.NewClient(&redis.Options{
		Addr:     uri,
		Password: key,
	})
	logs.Info("connected RedisSub")
	return RedisSub{repository: repository, redis: redis}
}

func (r RedisSub) SaveDNA() {
	subscriber := r.redis.Subscribe(ctx, constant.ChannelNameSaveDNA)

	dna := entity.Dna{}
	for {
		message, err := subscriber.ReceiveMessage(ctx)
		if err != nil {
			logs.Error("RedisSub.SaveDNA ReceiveMessage", err.Error())
			continue
		}

		err = json.Unmarshal([]byte(message.Payload), &dna)
		if err != nil {
			logs.Error("RedisSub.SaveDNA Unmarshal", err.Error())
			continue
		}

		dna.Id = security.GenerateHash(dna.Dna...)
		dna.CreatedAt = time.Now().Format(constant.DateFormat)

		r.repository.CreateDNA(dna)
	}
}
