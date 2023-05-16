package service

import (
	"GraduationDesign/dao"
	"GraduationDesign/models"
	"strconv"
)

type CartService struct {
}

var cartDao dao.CartDao

// 获得物品展示信息
func (CartService) GetUserCart(uid int) ([]models.GoodsData, error) {
	var gids []int
	result, err := cartDao.GetCartByUid(uid)
	if err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, err
	}
	for _, value := range result {
		gid, err := strconv.Atoi(value)
		if err != nil {
			return nil, err
		}
		gids = append(gids, gid)
	}
	goods, err := goodsDao.GetGoodsByGids(gids)
	if err != nil {
		return []models.GoodsData{}, err
	}
	var goodsData []models.GoodsData
	for _, good := range goods {
		var gd models.GoodsData
		gd.GoodsName = good.GoodsName
		gd.Gid = good.Gid
		gd.Price = good.Price
		gd.UsedTime = good.UsedTime
		gd.UserName, err = userInfoDao.GetUserNameByUid(good.Uid)
		if err != nil {
			return nil, err
		}
		gd.Picture, err = goodsDao.GetPictureByGid(good.Gid)
		if err != nil && err.Error() != "record not found" {
			return nil, err
		}
		goodsData = append(goodsData, gd)
	}
	return goodsData, nil
}

func (CartService) AddGoodInCart(uid int, gid int) error {
	err := cartDao.AddGoodInCategory(uid, gid)
	if err != nil {
		return err
	}
	return nil
}

func (CartService) RemoveGoodInCart(uid int, gid int) error {
	err := cartDao.RemoveGoodInCategory(uid, gid)
	if err != nil {
		return err
	}
	return nil
}

func (CartService) ClearCart(uid int) error {
	err := cartDao.ClearCart(uid)
	if err != nil {
		return err
	}
	return nil
}
