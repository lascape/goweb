package global

import (
	"github.com/bsm/redislock"
	"github.com/go-redis/redis/v8"
)

var _rd *redis.Client
var _rdLock *redislock.Client

func SetRB(rd *redis.Client) {
	_rd = rd
	_rdLock = redislock.New(rd)
}

func RD() *redis.Client {
	return _rd
}

func RDLocker() *redislock.Client {
	return _rdLock
}
