package dao

import (
	"GraduationDesign/models"
	"github.com/gomodule/redigo/redis"
)

type NoticeDao struct {
}

func (d NoticeDao) GetNotice() ([]string, error) {
	notices, err := redis.Strings(rd.Do("hgetall", "notice"))
	if err != nil {
		return nil, err
	}
	return notices, nil
}

func (d NoticeDao) UpdateNotice(date string, notice models.Notice) error {
	_, err := rd.Do("hdel", "notice", date)
	if err != nil {
		return err
	}
	_, err = rd.Do("hset", "notice", notice.Date, notice.Content)
	if err != nil {
		return err
	}
	return nil
}
