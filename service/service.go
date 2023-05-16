package service

import (
	"GraduationDesign/dao"
	"GraduationDesign/models"
)

type OrderService struct {
}

var orderDao dao.OrderDao

func (OrderService) CreateOrder(order models.Order) error {
	uid, err := goodsDao.GetUserUidByGid(order.Gid)
	order.Sid = uid
	order.Status = 0
	err = orderDao.CreateOrder(order)
	if err != nil {
		return err
	}
	return nil
}

func (OrderService) GetOrdersByUid(uid int) ([]models.OrderInfo, error) {
	orders, err := orderDao.GetOrdersBySid(uid)
	if err != nil {
		return nil, err
	}
	var orderInfos []models.OrderInfo
	for _, order := range orders {
		goods, err := goodsDao.GetGoodsByGid(order.Gid)
		if err != nil {
			return nil, err
		}
		var orderInfo models.OrderInfo
		orderInfo.Goods = goods
		orderInfo.Status = order.Status
		orderInfo.CreatedAt = order.CreatedAt
		orderInfo.Notes = order.Notes
		orderInfo.Place = order.Place
		orderInfo.Oid = order.Oid
		orderInfos = append(orderInfos, orderInfo)
	}
	return orderInfos, nil
}

var goodsService GoodsService

func (OrderService) GetOrderDetail(oid int) (models.OrderInfo, error) {
	order, err := orderDao.GetOrderByOid(oid)
	if err != nil {
		return models.OrderInfo{}, err
	}
	user, err := userDao.GetUserByUid(order.Uid)
	if err != nil {
		return models.OrderInfo{}, err
	}
	sUser, err := userDao.GetUserByUid(order.Sid)
	if err != nil {
		return models.OrderInfo{}, err
	}
	goods, err := goodsService.GetGoodsInfo(order.Gid, order.Sid)
	if err != nil {
		return models.OrderInfo{}, err
	}
	var orderDetail models.OrderInfo
	orderDetail.Oid = order.Oid
	orderDetail.Status = order.Status
	orderDetail.Place = order.Place
	orderDetail.Notes = order.Notes
	orderDetail.CreatedAt = order.CreatedAt
	orderDetail.UserName = user.UserName
	orderDetail.PhoneNumber = user.PhoneNumber
	orderDetail.SUserName = sUser.UserName
	orderDetail.SPhoneNumber = sUser.PhoneNumber
	orderDetail.GoodsInfo = goods
	return orderDetail, nil
}

func (OrderService) UpdateOrder(order models.Order) error {
	order.Status = 1
	err := orderDao.UpdateOrder(order)
	if err != nil {
		return err
	}
	return nil
}
