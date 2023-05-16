package dao

import (
	"GraduationDesign/models"
	"fmt"
	"log"
)

type OrderDao struct {
}

func (OrderDao) CreateOrder(order models.Order) error {
	tx := db.Create(&order)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (OrderDao) GetOrdersByUid(uid int) ([]models.Order, error) {
	var orders []models.Order
	tx := db.Model(&models.Order{}).Where("uid=?", uid).Find(&orders)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return orders, nil
}

func (OrderDao) CountAllOrder() (int, error) {
	var count int64
	tx := db.Model(&models.Order{}).Count(&count)
	if tx.Error != nil {
		log.Println(tx.Error)
		return 0, tx.Error
	}
	return int(count), nil
}

func (OrderDao) GetOrdersBySid(sid int) ([]models.Order, error) {
	var orders []models.Order
	tx := db.Model(&models.Order{}).Where("sid=?", sid).Find(&orders)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return orders, nil
}

func (OrderDao) GetOrderByOid(oid int) (models.Order, error) {
	var order models.Order
	tx := db.Where("oid=?", oid).First(&order)
	if tx.Error != nil {
		return order, tx.Error
	}
	return order, nil
}

func (OrderDao) UpdateOrder(order models.Order) error {
	fmt.Println(order)
	tx := db.Model(models.Order{}).Where("oid=?", order.Oid).Updates(&order)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
