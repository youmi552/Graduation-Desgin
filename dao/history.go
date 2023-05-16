package dao

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"time"
)

type HistoryDao struct {
}

// 查询历史记录数量
func (HistoryDao) CountHistory(uid int) (int, error) {
	history := fmt.Sprintf("history:%d", uid)
	count, err := redis.Int64(rd.Do("zcount", history, 0, 1000000000))
	if err != nil {
		return 0, err
	}
	return int(count), nil
}

// 添加历史记录
func (HistoryDao) AddHistory(uid int, gid int) {
	time := time.Now().Unix()
	history := fmt.Sprintf("history:%d", uid)
	rd.Do("zadd", history, time, gid)
}

func (HistoryDao) RemHistory(uid int) error {
	history := fmt.Sprintf("history:%d", uid)
	_, err := rd.Do("zremrangebyrank", history, 0, 0)
	if err != nil {
		return err
	}
	return nil
}

func (HistoryDao) GetHistoryByUid(uid int) ([]string, error) {
	history := fmt.Sprintf("history:%d", uid)
	reply, err := redis.Strings(rd.Do("zrevrange", history, 0, -1))
	if err != nil {
		return nil, err
	}
	return reply, nil
}
