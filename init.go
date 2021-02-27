package redigopack

import "github.com/gomodule/redigo/redis"

const (
	defaultKeyPrefix = ""
)

type RedisCacheClient struct {
	client        *redis.Pool
	keyPrefix     string
	defaultMaxAge int
}

// InitRedis init
func InitRedis(host, password string, db int) *RedisCacheClient {
	redisPool := &redis.Pool{
		MaxActive: 10,
		MaxIdle:   10,
		Wait:      true,
		Dial: func() (redis.Conn, error) {
			return redis.Dial(
				"tcp",
				host,
				redis.DialPassword(password),
				redis.DialDatabase(db),
			)
		},
	}

	return &RedisCacheClient{redisPool, defaultKeyPrefix, 300}
}

func (c *RedisCacheClient) SetKeyPrefix(prefix string) *RedisCacheClient {
	c.keyPrefix = prefix
	return c
}
