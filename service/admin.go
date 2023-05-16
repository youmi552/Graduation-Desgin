package service

import (
	"GraduationDesign/models"
	"GraduationDesign/util"
	"strconv"
)

type AdminService struct {
}

func (AdminService) GetData() (models.Data, error) {
	var data models.Data
	users, err := userDao.CountAllUsers()
	if err != nil {
		return data, err
	}
	data.UserNumber = users
	goods, err := goodsDao.CountAllGoodsWithAdmin()
	if err != nil {
		return data, err
	}
	data.GoodsNumber = goods
	advice, err := adviceDao.CountAllAdvice()
	if err != nil {
		return data, err
	}
	data.AdviceNumber = advice
	order, err := orderDao.CountAllOrder()
	if err != nil {
		return data, err
	}
	data.OrderNumber = order
	var goodsCategorys []models.Category
	result, err := categoryDao.GetAllCategorys()
	if err != nil {
		return data, err
	}
	for key, value := range result {
		goodsCategorys = append(goodsCategorys, models.Category{key, value})
	}
	data.GoodsCategory = goodsCategorys
	var adviceCategorys []models.Category
	result2, err := categoryDao.GetAllAdviceCategorys()
	if err != nil {
		return data, err
	}
	for key, value := range result2 {
		adviceCategorys = append(adviceCategorys, models.Category{key, value})
	}
	data.AdviceCategory = adviceCategorys

	var categoryDatas []models.CategoryData
	for _, category := range data.GoodsCategory {
		var categoryData models.CategoryData
		categoryData.Cid = category.Cid
		categoryData.CName = category.CName
		cid, _ := strconv.Atoi(category.Cid)
		count, err := goodsDao.CountGoodsByCidWithAdmin(cid)
		if err != nil {
			return data, err
		}
		categoryData.Percent = util.Decimal(float64(float64(count) / float64(data.GoodsNumber) * 100))
		categoryDatas = append(categoryDatas, categoryData)
	}
	data.CategoryData = categoryDatas
	return data, nil
}

func (AdminService) GetUserData() ([]models.UserData, error) {
	var usersData []models.UserData
	users, err := userDao.GetAllUser()
	if err != nil {
		return nil, err
	}
	for _, user := range users {
		var userData models.UserData
		userData.UserName = user.UserName
		userData.Uid = user.Uid
		userData.PhoneNumber = user.PhoneNumber
		userData.Level = user.Level
		usersData = append(usersData, userData)
	}
	return usersData, err
}

func (AdminService) UpdateUserData(user models.User) error {
	err := userDao.UpdateUserData(user)
	if err != nil {
		return err
	}
	return nil
}

func (AdminService) DeleteUser(uid int) error {
	err := userDao.DeleteUser(uid)
	if err != nil {
		return err
	}
	return nil
}

func (AdminService) GetAdvice() ([]models.AdviceData, error) {
	advice, err := adviceDao.GetAdvice()
	if err != nil {
		return nil, err
	}
	var adviceData []models.AdviceData
	for _, item := range advice {
		var aData models.AdviceData
		aData.Aid = item.Aid
		user, err := userDao.GetUserByUid(item.Uid)
		if err != nil {
			return nil, err
		}
		aData.UserName = user.UserName
		categoryName, err := categoryDao.GetAdviceCategoryNameByCid(item.Cid)
		if err != nil {
			return nil, err
		}
		aData.Category = categoryName
		aData.Title = item.Title
		aData.Status = item.Status
		aData.CreatedAt = item.CreatedAt.Format("2006-01-02 15:04:05")
		adviceData = append(adviceData, aData)
	}
	return adviceData, nil
}

func (AdminService) ConfirmAdvice(aid int) error {
	var advice models.Advice
	advice.Aid = aid
	advice.Status = 1
	err := adviceDao.UpdateAdvice(advice)
	if err != nil {
		return err
	}
	return nil
}

func (AdminService) DeleteAdvice() error {
	err := adviceDao.DeleteAdviceByStatus()
	if err != nil {
		return err
	}
	return nil
}

func (AdminService) GetGoodsByPage(pageNumber int, pageSize int, keyword string) (models.GoodsPageDataWithAdmin, error) {
	var err error
	offset := (pageNumber - 1) * pageSize
	count := 0
	ban := 0
	if keyword == "" {
		ban, err = goodsDao.CountBanGoods()
		count, err = goodsDao.CountAllGoodsWithAdmin()
		if err != nil {
			return models.GoodsPageDataWithAdmin{}, err
		}
	} else {
		count, err = goodsDao.CountGoodsByKeyword(keyword)
		if err != nil {
			return models.GoodsPageDataWithAdmin{}, err
		}
	}
	goods, err := goodsDao.GetGoodsByPageWithAdmin(offset, pageSize, keyword)
	if err != nil {
		return models.GoodsPageDataWithAdmin{}, err
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
			return models.GoodsPageDataWithAdmin{}, err
		}
		goodsData = append(goodsData, gd)
	}
	var goodsPageData models.GoodsPageDataWithAdmin
	goodsPageData.Ban = ban
	goodsPageData.Count = count
	goodsPageData.GoodsData = goodsData
	return goodsPageData, nil
}
