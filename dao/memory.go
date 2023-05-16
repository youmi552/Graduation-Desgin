package dao

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"log"
)

type MemoryDao struct {
}

func (MemoryDao) GetMeroryInfo() {
	stringMap, err := redis.String(rd.Do("info", "memory"))
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(stringMap)
}
