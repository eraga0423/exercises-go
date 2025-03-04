package board

import (
	"github.com/my/repo/internal/config"
	pubsub "github.com/my/repo/internal/types/redis"
	"github.com/redis/go-redis/v9"
)

type LeaderBoard struct {
	pubsub.MyRedis
	conf *config.Config
	rdb  *redis.Client
}

func NewLeaderBoard(conf *config.Config, rdb *redis.Client) *LeaderBoard {
	return &LeaderBoard{
		conf: conf,
		rdb:  rdb,
	}
}
