package dao

import (
	"GraduationDesign/models"
	"fmt"
	"github.com/gomodule/redigo/redis"
)

type CategoryDao struct {
}

// 添加种类
func (CategoryDao) CreateCategory(category models.CategoryParam) error {
	_, err := redis.Int64(rd.Do("hset", category.Type, category.Cid, category.CName))
	if err != nil {
		return err
	}
	return nil
}

// 通过Cid获得种类名称
func (CategoryDao) GetCategoryNameByCid(cid int) (string, error) {
	reply, err := redis.String(rd.Do("hget", "category", cid))
	if err != nil {
		return "", err
	}
	return reply, nil
}

// 获得所有种类
func (CategoryDao) GetAllCategorys() (map[string]string, error) {
	fmt.Println("这里出错")
	reply, err := redis.StringMap(rd.Do("hgetall", "category"))
	if err != nil {
		return nil, err
	}
	return reply, nil
}

// 添加建议种类
func (CategoryDao) CreateAdviceCategory(category models.Category) error {
	_, err := redis.Int64(rd.Do("hset", "adcategory", category.Cid, category.CName))
	if err != nil {
		return err
	}
	return nil
}

// 通过Cid获得建议种类名称
func (CategoryDao) GetAdviceCategoryNameByCid(cid int) (string, error) {
	reply, err := redis.String(rd.Do("hget", "adcategory", cid))
	if err != nil {
		return "", err
	}
	return reply, nil
}

// 获得所有建议种类
func (CategoryDao) GetAllAdviceCategorys() (map[string]string, error) {
	reply, err := redis.StringMap(rd.Do("hgetall", "adcategory"))
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func (CategoryDao) DeleteCategory(category models.CategoryParam) error {
	_, err := redis.Int64(rd.Do("hdel", category.Type, category.Cid))
	if err != nil {
		return err
	}
	return nil
}
