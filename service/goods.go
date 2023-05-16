package service

import (
	"GraduationDesign/dao"
	"GraduationDesign/models"
	"strconv"
)

type GoodsService struct {
}

var goodsDao dao.GoodsDao

// 获得物品展示信息
func (GoodsService) GetGoodsData() ([]models.GoodsData, error) {
	goods, err := goodsDao.GetGoods()
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

// 获得所有商品展示信息并分页
func (GoodsService) GetGoodsByPage(pageNumber int, pageSize int, keyword string) (models.GoodsPageData, error) {
	var err error
	offset := (pageNumber - 1) * pageSize
	count := 0
	if keyword == "" {
		count, err = goodsDao.CountAllGoods()
		if err != nil {
			return models.GoodsPageData{}, err
		}
	} else {
		count, err = goodsDao.CountGoodsByKeyword(keyword)
		if err != nil {
			return models.GoodsPageData{}, err
		}
	}
	goods, err := goodsDao.GetGoodsByPage(offset, pageSize, keyword)
	if err != nil {
		return models.GoodsPageData{}, err
	}
	var goodsData []models.GoodsData
	for _, good := range goods {
		var gd models.GoodsData
		gd.GoodsName = good.GoodsName
		gd.Gid = good.Gid
		gd.Price = good.Price
		gd.UsedTime = good.UsedTime
		gd.UserName, err = userInfoDao.GetUserNameByUid(good.Uid)
		gd.Status = good.Status
		if err != nil {
			return models.GoodsPageData{}, err
		}
		gd.Picture, err = goodsDao.GetPictureByGid(good.Gid)
		if err != nil && err.Error() != "record not found" {
			return models.GoodsPageData{}, err
		}
		goodsData = append(goodsData, gd)
	}
	var goodsPageData models.GoodsPageData
	goodsPageData.Count = count
	goodsPageData.GoodsData = goodsData
	return goodsPageData, nil
}

// 获得所有分类商品展示信息并分页
func (GoodsService) GetGoodsByCategory(pageNumber int, pageSize int, cid int) (models.GoodsPageData, error) {

	offset := (pageNumber - 1) * pageSize
	count, err := goodsDao.CountGoodsByCid(cid)
	if err != nil {
		return models.GoodsPageData{}, err
	}
	goods, err := goodsDao.GetGoodsByCategory(offset, pageSize, cid)
	if err != nil {
		return models.GoodsPageData{}, err
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
			return models.GoodsPageData{}, err
		}
		gd.Picture, err = goodsDao.GetPictureByGid(good.Gid)
		if err != nil && err.Error() != "record not found" {
			return models.GoodsPageData{}, err
		}
		goodsData = append(goodsData, gd)
	}
	var goodsPageData models.GoodsPageData
	goodsPageData.Count = count
	goodsPageData.GoodsData = goodsData
	return goodsPageData, nil
}

// 上传物品
func (GoodsService) UploadGoods(goods models.Goods) (models.GoodsInfo, error) {
	categoryName, err := categoryDao.GetCategoryNameByCid(goods.Cid)
	userName, err := userInfoDao.GetUserNameByUid(goods.Uid)
	if err != nil {
		return models.GoodsInfo{}, err
	}
	if err != nil {
		return models.GoodsInfo{}, err
	}
	gid, err := goodsDao.CreateGoods(goods)
	if err != nil {
		return models.GoodsInfo{}, err
	}
	err = goodsDao.CreateViews(gid)
	if err != nil {
		return models.GoodsInfo{}, err
	}
	err = goodsDao.CreateCollection(gid)
	if err != nil {
		return models.GoodsInfo{}, err
	}
	var goodsInfo = models.GoodsInfo{
		Gid:          gid,
		GoodsName:    goods.GoodsName,
		UserName:     userName,
		CategoryName: categoryName,
		Introduction: goods.Introduction,
		UsedTime:     goods.UsedTime,
		Price:        goods.Price,
		Status:       goods.Status,
	}
	return goodsInfo, nil
}

func (GoodsService) UploadPicture(gid int, picture string) error {
	err := goodsDao.CreatePicture(gid, picture)
	if err != nil {
		return err
	}
	return nil
}

// 获取商品信息
func (GoodsService) GetGoodsInfo(gid int, uid int) (models.GoodsInfo, error) {
	goods, err := goodsDao.GetGoodsByGid(gid)
	if err != nil {
		return models.GoodsInfo{}, err
	}
	userName, err := userInfoDao.GetUserNameByUid(goods.Uid)
	if err != nil {
		return models.GoodsInfo{}, err
	}
	categoryName, err := categoryDao.GetCategoryNameByCid(goods.Cid)
	if err != nil {
		return models.GoodsInfo{}, err
	}
	pictures, err := goodsDao.GetPicturesByGid(gid)
	if err != nil {
		return models.GoodsInfo{}, err
	}
	views, err := goodsDao.AddViews(gid)
	if err != nil {
		return models.GoodsInfo{}, err
	}
	collection, err := goodsDao.GetCollection(gid)
	if err != nil {
		return models.GoodsInfo{}, err
	}
	isCollection, err := goodsDao.CheckCollection(uid, gid)
	if err != nil {
		return models.GoodsInfo{}, err
	}
	var goodsInfo = models.GoodsInfo{
		Gid:          gid,
		GoodsName:    goods.GoodsName,
		UserName:     userName,
		CategoryName: categoryName,
		Pictures:     pictures,
		Introduction: goods.Introduction,
		UsedTime:     goods.UsedTime,
		Price:        goods.Price,
		Status:       goods.Status,
		Views:        views,
		Collection:   collection,
		IsCollection: isCollection,
	}
	return goodsInfo, nil
}

func (GoodsService) GetGoodsDataByUid(uid int, pageNumber int, pageSize int) (models.GoodsPageData, error) {
	offset := (pageNumber - 1) * pageSize
	count, err := goodsDao.CountGoodsByUid(uid)
	if err != nil {
		return models.GoodsPageData{}, err
	}
	goods, err := goodsDao.GetGoodsByUid(uid, pageSize, offset)
	if err != nil {
		return models.GoodsPageData{}, err
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
			return models.GoodsPageData{}, err
		}
		gd.Picture, err = goodsDao.GetPictureByGid(good.Gid)
		if err != nil && err.Error() != "record not found" {
			return models.GoodsPageData{}, err
		}
		goodsData = append(goodsData, gd)
	}
	var goodsPageData models.GoodsPageData
	goodsPageData.Count = count
	goodsPageData.GoodsData = goodsData
	return goodsPageData, nil
}

func (GoodsService) CancelCollection(uid int, gid string) (int, error) {
	err := goodsDao.CancelCollection(uid, gid)
	if err != nil {
		return 0, err
	}
	gid2, err := strconv.Atoi(gid)
	if err != nil {
		return 0, err
	}
	collection, err := goodsDao.ReduceCollection(gid2)
	if err != nil {
		return 0, err
	}
	return collection, nil
}

func (GoodsService) Collection(uid int, gid string) (int, error) {
	err := goodsDao.Collection(uid, gid)
	if err != nil {
		return 0, err
	}

	gid2, err := strconv.Atoi(gid)
	if err != nil {
		return 0, err
	}
	collection, err := goodsDao.AddCollection(gid2)
	if err != nil {
		return 0, err
	}
	return collection, nil
}

func (GoodsService) GetUserCollection(uid int) ([]models.GoodsData, error) {
	var gids []int
	result, err := goodsDao.GetCollectionByUid(uid)
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

func (GoodsService) UpdateGoodsStatus(gid int, status int) error {
	err := goodsDao.UpdateGoodsStatus(gid, status)
	if err != nil {
		return err
	}
	return nil
}
