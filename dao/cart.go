package dao

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"time"
)

type CartDao struct {
}

// 添加商品到购物车
func (CartDao) AddGoodInCategory(uid int, gid int) error {
	time := time.Now().Unix()
	cart := fmt.Sprintf("cart:%d", uid)
	_, err := redis.Int64(rd.Do("zadd", cart, time, gid))
	if err != nil {
		return err
	}
	return nil
}

// 移除商品到购物车
func (CartDao) RemoveGoodInCategory(uid int, gid int) error {
	cart := fmt.Sprintf("cart:%d", uid)
	_, err := redis.Int64(rd.Do("zrem", cart, gid))
	if err != nil {
		return err
	}
	return nil
}

// 获得用户的购物车
func (CartDao) GetCartByUid(uid int) ([]string, error) {
	cart := fmt.Sprintf("cart:%d", uid)
	reply, err := redis.Strings(rd.Do("zrevrange", cart, 0, -1))
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func (CartDao) ClearCart(uid int) error {
	cart := fmt.Sprintf("cart:%d", uid)
	_, err := rd.Do("expire", cart, 0)
	if err != nil {
		return err
	}
	return nil
}
