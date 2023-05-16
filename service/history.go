package service

import (
	"GraduationDesign/dao"
	"GraduationDesign/models"
	"strconv"
)

type HistoryService struct {
}

var historyDao dao.HistoryDao

// 添加历史记录
func (HistoryService) AddHistory(uid int, gid int) {
	historyDao.AddHistory(uid, gid)
	count, err := historyDao.CountHistory(uid)
	if err != nil {
		return
	}
	if count > 10 {
		err := historyDao.RemHistory(uid)
		if err != nil {
			return
		}
	}
}

func (HistoryService) GetUserHistory(uid int) ([]models.GoodsData, error) {
	var gids []int
	result, err := historyDao.GetHistoryByUid(uid)
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
