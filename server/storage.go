package main

import (
	"github.com/garyburd/redigo/redis"
)

const SS_PREFIX = "se:"

type User struct {
	Name     string `redis:"-"`
	Password string `redis:"password"`
	Method   string `redis:"method"`
}

type Storage struct {
	pool *redis.Pool
}

func NewStorage(server string) *Storage {
	pool := redis.NewPool(func() (conn redis.Conn, err error) {
		conn, err = redis.Dial("tcp", server)
		return
	}, 3)
	return &Storage{pool}
}

func (s *Storage) GetUser(name string) (user User, err error) {
	var conn = s.pool.Get()
	defer conn.Close()
	user.Name = name
	data, _ := redis.Values(conn.Do("HGETALL", SS_PREFIX + name))
	err = redis.ScanStruct(data, &user)
	return
}




