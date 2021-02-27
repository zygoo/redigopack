package redigopack

import (
	"errors"

	"github.com/gomodule/redigo/redis"
)

// Get returns value in cache
func (c *RedisCacheClient) Get(key string) ([]byte, error) {
	conn := c.client.Get()
	defer conn.Close()
	if err := conn.Err(); err != nil {
		return nil, err
	}

	data, err := conn.Do("GET", c.keyPrefix+key)
	if err != nil {
		return nil, err
	}

	// data == nil && err == nil
	if data == nil {
		return nil, errors.New("empty") // no data was associated with this key
	}

	b, err := redis.Bytes(data, err)
	if err != nil {
		return nil, err
	}

	return b, err
}

// Set set value in cache
func (c *RedisCacheClient) Set(key string, value []byte, expire int) error {
	conn := c.client.Get()
	defer conn.Close()
	if err := conn.Err(); err != nil {
		return err
	}

	if expire == 0 {
		expire = c.defaultMaxAge
	}

	_, err := conn.Do("SETEX", c.keyPrefix+key, expire, value)
	return err
}

// Del set value in cache
func (c *RedisCacheClient) Del(key string) error {
	conn := c.client.Get()
	defer conn.Close()
	if err := conn.Err(); err != nil {
		//log.Error("get redis conn err, %v", err)
		return err
	}

	_, err := conn.Do("DEL", c.keyPrefix+key)
	return err
}

// sort set in cache zadd
func (c *RedisCacheClient) Zadd(key string, score int64, member string) error {
	conn := c.client.Get()
	defer conn.Close()
	if err := conn.Err(); err != nil {
		return err
	}
	_, err := conn.Do("ZADD", c.keyPrefix+key, score, member)
	return err
}

// sort set in cache zscore
func (c *RedisCacheClient) Zscore(key string, member string) ([]byte, error) {
	conn := c.client.Get()
	defer conn.Close()
	if err := conn.Err(); err != nil {
		return nil, err
	}
	data, err := conn.Do("ZSCORE", c.keyPrefix+key, member)

	if err != nil {
		return nil, err
	}

	if data == nil {
		return nil, errors.New("empty") // no data was associated with this key
	}

	b, err := redis.Bytes(data, err)
	if err != nil {
		return nil, err
	}
	return b, err
}

// sort set in cache, Zrangebyscore
func (c *RedisCacheClient) Zrangebyscore(key string, min, max int64) ([][]byte, error) {
	var res [][]byte
	conn := c.client.Get()
	defer conn.Close()
	if err := conn.Err(); err != nil {
		return nil, err
	}
	data, err := conn.Do("ZRANGEBYSCORE", c.keyPrefix+key, min, max)
	if err != nil {
		return nil, err
	}

	if data == nil {
		return nil, errors.New("empty") // no data was associated with this key
	}

	b, err := redis.Values(data, err)
	if err != nil {
		return nil, err
	}

	for _, v := range b {
		res = append(res, v.([]byte))
	}

	return res, err
}

// sort set in cache, Zremrangebyscore
func (c *RedisCacheClient) Zremrangebyscore(key string, min, max int64) error {
	conn := c.client.Get()
	defer conn.Close()
	if err := conn.Err(); err != nil {
		return err
	}
	_, err := conn.Do("ZREMRANGEBYSCORE", c.keyPrefix+key, min, max)
	return err
}
