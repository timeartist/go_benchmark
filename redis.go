package main

import "github.com/garyburd/redigo/redis"

type RedisFixture struct {
    R redis.Conn
}

func (self *RedisFixture) Set(key string, value string) {
    result, err := redis.String(self.R.Do("SET", key, value))
    
    if result != "OK" || err != nil {
        panic("error with redis call")
    }
}

func (self *RedisFixture) Get(key string) string {
    result, err := redis.String(self.R.Do("GET", key))
    
    if err != nil {
        panic(err)
    }
    
    return result
    
}

func (self *RedisFixture) Close() {
    self.R.Close()
}

func createRedis(url string) RedisFixture {
    r, err := redis.DialURL(url)
    if err != nil {
        panic(err)
    }
    return RedisFixture{R:r}
}