package redis

import (
	"github.com/garyburd/redigo/redis"
	"testing"
)

func TestRedisConnection(t *testing.T) {
	pool := NewRedis("localhost:6379", "")
	conn := pool.Connect()
	defer conn.Close()

	reply, err := conn.Do("PING")

	if err != nil {
		t.Log("Cannot reach redis server")
	}

	result, err := redis.String(reply, err)

	if result != "PONG" {
		t.Error("Cannot connect to redis server")
	}
}
