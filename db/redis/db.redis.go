package redis

import (
	"github.com/garyburd/redigo/redis"
	"time"
)

// Redis command constant
const (
	GET  = "GET"
	SET  = "SET"
	AUTH = "AUTH"
)

type Redis struct {
	pool *redis.Pool
}

var redisPool *redis.Pool

func NewRedis(server string, password string) *Redis {
	redis := new(Redis)
	redis.pool = newRedisPool(server, password)

	return redis
}

func newRedisPool(server string, password string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     5,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", server)

			if err != nil {
				return nil, err
			}

			if password != "" {
				if _, err := conn.Do("AUTH", password); err != nil {
					conn.Close()
					return nil, err
				}
			}

			return conn, nil
		},
		TestOnBorrow: func(conn redis.Conn, time time.Time) error {
			_, err := conn.Do("PING")
			return err
		},
	}
}

func (redis *Redis) Connect() redis.Conn {
	return redis.pool.Get()
}
