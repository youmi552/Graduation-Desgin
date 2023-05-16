package dao

import (
	"GraduationDesign/models"
	"errors"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"gorm.io/gorm"
	"log"
	"time"
)

type GoodsDao struct {
}

// 获取所有物品信息
func (GoodsDao) GetGoods() ([]models.Goods, error) {
	var goods []models.Goods
	tx := db.Find(&goods)
	if tx.Error != nil {
		log.Println(tx.Error)
		return nil, tx.Error
	}
	return goods, nil
}

// 获取所有物品信息(分页)
func (GoodsDao) GetGoodsByPage(offset int, limit int, keyword string) ([]models.Goods, error) {
	var goods []models.Goods
	var tx *gorm.DB
	if keyword == "" {
		tx = db.Where("status=1").Offset(offset).Limit(limit).Find(&goods)
	} else {
		tx = db.Where("goods_name LIKE ?", "%"+keyword+"%").Offset(offset).Limit(limit).Find(&goods)
	}
	if tx.Error != nil {
		log.Println(tx.Error)
		return nil, tx.Error
	}
	return goods, nil
}

// 管理员获取所有物品信息(分页)
func (GoodsDao) GetGoodsByPageWithAdmin(offset int, limit int, keyword string) ([]models.Goods, error) {
	var goods []models.Goods
	var tx *gorm.DB
	if keyword == "" {
		tx = db.Order("status").Offset(offset).Limit(limit).Find(&goods)
	} else {
		tx = db.Where("goods_name LIKE ?", "%"+keyword+"%").Offset(offset).Limit(limit).Find(&goods)
	}
	if tx.Error != nil {
		log.Println(tx.Error)
		return nil, tx.Error
	}
	return goods, nil
}

func (GoodsDao) GetGoodsByCategory(offset int, limit int, cid int) ([]models.Goods, error) {
	var goods []models.Goods
	tx := db.Where("cid=? and status=1", cid).Offset(offset).Limit(limit).Find(&goods)
	if tx.Error != nil {
		log.Println(tx.Error)
		return nil, tx.Error
	}
	return goods, nil
}

// 获得物品数量
func (GoodsDao) CountAllGoods() (int, error) {
	var count int64
	tx := db.Where("status=1").Model(&models.Goods{}).Count(&count)
	if tx.Error != nil {
		log.Println(tx.Error)
		return 0, tx.Error
	}
	return int(count), nil
}

// 获得物品数量
func (GoodsDao) CountAllGoodsWithAdmin() (int, error) {
	var count int64
	tx := db.Model(&models.Goods{}).Count(&count)
	if tx.Error != nil {
		log.Println(tx.Error)
		return 0, tx.Error
	}
	return int(count), nil
}

// 获得物品数量
func (GoodsDao) CountBanGoods() (int, error) {
	var count int64
	tx := db.Model(&models.Goods{}).Where("status=?", 0).Count(&count)
	if tx.Error != nil {
		log.Println(tx.Error)
		return 0, tx.Error
	}
	return int(count), nil
}
func (GoodsDao) CountGoodsByCid(cid int) (int, error) {
	var count int64
	tx := db.Model(&models.Goods{}).Where("cid=? and status=1", cid).Count(&count)
	if tx.Error != nil {
		log.Println(tx.Error)
		return 0, tx.Error
	}
	return int(count), nil
}
func (GoodsDao) CountGoodsByCidWithAdmin(cid int) (int, error) {
	var count int64
	tx := db.Model(&models.Goods{}).Where("cid=? ", cid).Count(&count)
	if tx.Error != nil {
		log.Println(tx.Error)
		return 0, tx.Error
	}
	return int(count), nil
}

// 新建物品
func (GoodsDao) CreateGoods(goods models.Goods) (int, error) {
	tx := db.Create(&goods)
	if tx.Error != nil {
		return goods.Uid, tx.Error
	}
	return goods.Gid, nil
}

// 获取商品的图片信息
func (GoodsDao) GetPicturesByGid(gid int) ([]string, error) {
	var pictures []string
	tx := db.Model(&models.Pictures{}).Select("picture").Where("gid=?", gid).Find(&pictures)
	if tx.Error != nil {
		return pictures, tx.Error
	}
	return pictures, nil
}

// 获取商品的展示图片信息
func (GoodsDao) GetPictureByGid(gid int) (string, error) {
	var picture string
	tx := db.Model(&models.Pictures{}).Select("picture").Where("gid=?", gid).First(&picture)
	if tx.Error != nil {
		return picture, tx.Error
	}
	return picture, nil
}

func (GoodsDao) CreatePicture(gid int, picture string) error {
	var p = &models.Pictures{Gid: gid, Picture: picture}
	tx := db.Create(&p)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (GoodsDao) GetGoodsByGid(gid int) (models.Goods, error) {
	var goods models.Goods
	tx := db.Where("gid=?", gid).First(&goods)
	if tx.Error != nil {
		return goods, tx.Error
	}
	return goods, nil
}

//func (GoodsDao) GetCategoryNameByCid(cid int) (string, error) {
//	reply, err := redis.String(rd.Do("hget", "category", cid))
//	if err != nil {
//		return "", err
//	}
//	return reply, nil
//}

// 创建浏览量
func (GoodsDao) CreateViews(gid int) error {
	_, err := rd.Do("hsetnx", "views", gid, 0)
	if err != nil {
		return err
	}
	return nil
}

// 增加浏览量
func (GoodsDao) AddViews(gid int) (int, error) {
	views, err := redis.Int(rd.Do("hincrby", "views", gid, "1"))
	if err != nil {
		return -1, err
	}
	return views, nil

}

// 获得用户的信息
func (GoodsDao) GetGoodsByUid(uid int, limit int, offset int) ([]models.Goods, error) {
	var goods []models.Goods
	tx := db.Offset(offset).Limit(limit).Where("uid=?", uid).Find(&goods)
	if tx.Error != nil {
		log.Println(tx.Error)
		return nil, tx.Error
	}

	return goods, nil
}

func (GoodsDao) CountGoodsByUid(uid int) (int, error) {
	var count int64
	tx := db.Model(&models.Goods{}).Where("uid=?", uid).Count(&count)
	if tx.Error != nil {
		log.Println(tx.Error)
		return 0, tx.Error
	}
	return int(count), nil
}

func (GoodsDao) CheckCollection(uid int, gid int) (bool, error) {
	collection := fmt.Sprintf("collection:%d", uid)
	collect, err := redis.Int(rd.Do("zscore", collection, gid))
	if err != nil {
		if err.Error() == "redigo: nil returned" {
			return false, nil
		}
		return false, err
	}
	if collect > 0 {
		return true, nil
	}
	return false, errors.New("redis查询收藏量发生了未知的错误")
}

// 取消收藏
func (GoodsDao) CancelCollection(uid int, gid string) error {
	collection := fmt.Sprintf("collection:%d", uid)
	time := time.Now().Unix()
	_, err := redis.Int(rd.Do("zrem", collection, time, gid))
	if err != nil {
		return err
	}
	return nil
}

// 收藏
func (GoodsDao) Collection(uid int, gid string) error {
	collection := fmt.Sprintf("collection:%d", uid)
	time := time.Now().Unix()
	_, err := redis.Int(rd.Do("zadd", collection, time, gid))
	if err != nil {
		return err
	}
	return nil
}

// 新建收藏量
func (GoodsDao) CreateCollection(gid int) error {
	_, err := rd.Do("hsetnx", "collection", gid, 0)
	if err != nil {
		return err
	}
	return nil
}

// 查询收藏量
func (GoodsDao) GetCollection(gid int) (int, error) {
	collection, err := redis.Int(rd.Do("hget", "collection", gid))
	if err != nil {
		return -1, err
	}
	return collection, nil
}

// 增加收藏量
func (GoodsDao) AddCollection(gid int) (int, error) {
	collection, err := redis.Int(rd.Do("hincrby", "collection", gid, "1"))
	if err != nil {
		return -1, err
	}
	return collection, nil
}

// 减少收藏量
func (GoodsDao) ReduceCollection(gid int) (int, error) {
	collection, err := redis.Int(rd.Do("hincrby", "collection", gid, "-1"))
	if err != nil {
		return -1, err
	}
	return collection, nil
}

// 通过一定gid查询物品
func (GoodsDao) GetGoodsByGids(gids []int) ([]models.Goods, error) {
	//fmt.Println(gids)
	//var goods []models.Goods
	//tx := db.Find(&goods, &gids)
	//if tx.Error != nil {
	//	return nil, tx.Error
	//}
	//return goods, nil
	var goods []models.Goods
	for _, gid := range gids {
		var good models.Goods
		tx := db.First(&good, gid)
		if tx.Error != nil {
			return nil, tx.Error
		}
		goods = append(goods, good)
	}
	return goods, nil
}

func (GoodsDao) GetCollectionByUid(uid int) ([]string, error) {
	collection := fmt.Sprintf("collection:%d", uid)
	reply, err := redis.Strings(rd.Do("zrevrange", collection, 0, -1))
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func (GoodsDao) GetUserUidByGid(gid int) (int, error) {
	var uid int
	tx := db.Model(&models.Goods{}).Select("uid").Where("gid=?", gid).First(&uid)
	if tx.Error != nil {
		return uid, tx.Error
	}
	return uid, tx.Error
}

func (GoodsDao) CountGoodsByKeyword(keyword string) (int, error) {
	var count int64
	tx := db.Model(&models.Goods{}).Where("goods_name LIKE ?", "%"+keyword+"%").Count(&count)
	if tx.Error != nil {
		log.Println(tx.Error)
		return 0, tx.Error
	}
	return int(count), nil
}

func (GoodsDao) UpdateGoodsStatus(gid int, status int) error {
	tx := db.Model(&models.Goods{}).Where("gid=?", gid).Update("status", status)
	if tx.Error != nil {
		log.Println(tx.Error)
		return tx.Error
	}
	return nil
}
