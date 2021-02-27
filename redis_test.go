package redigopack

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sirupsen/logrus"
)

var redisClient *RedisCacheClient

func TestMain(m *testing.M) {
	logrus.Info("redigo pack test begin")
	redisClient = InitRedis("127.0.0.1:6379", "", 0)
	m.Run()
	logrus.Info("redigo pack test end")
}

func TestSetGetDel(t *testing.T) {
	logrus.Info("redigo pack set test begin")
	type setTable struct {
		key    string
		value  []byte
		expire int
	}

	// set command test table

	st := []setTable{
		{key: "redigopack_1", value: []byte("redigopack_value_1"), expire: 10},
		{key: "redigopack_2", value: []byte("redigopack_value_2"), expire: 100},
		{key: "redigopack_3", value: []byte("redigopack_value_3"), expire: 1000},
		{key: "redigopack_4", value: []byte("redigopack_value_4")},
	}

	for _, s := range st {
		err := redisClient.Set(s.key, s.value, s.expire)
		if err != nil {
			t.Errorf("set value error, key: %v, err: %v", s.key, err)
		}
	}

	value1, err := redisClient.Get("redigopack_1")
	if err != nil {
		t.Errorf("get value error, key: %v, err: %v", "redigopack_1", err)
	}

	assert.Equal(t, "redigopack_value_1", string(value1), "The two words should be the same")

	// get empty value
	_, err = redisClient.Get("redigopack_5")
	if err != nil {
		logrus.Info("get redigopack_5 value, empty error")
	}

	// del empty value
	err = redisClient.Del("redigopack_5")
	if err != nil {
		t.Errorf("del value error, key: %v, err: %v", "redigopack_5", err)
	}

	for _, s := range st {
		err := redisClient.Del(s.key)
		if err != nil {
			t.Errorf("del value error, key: %v, err: %v", s.key, err)
		}
	}
}
