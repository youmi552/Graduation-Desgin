package dao

import (
	"GraduationDesign/models"
	"log"
)

type AdviceDao struct {
}

func (AdviceDao) AddAdvice(advice models.Advice) (int, error) {
	tx := db.Create(&advice)
	if tx.Error != nil {
		return advice.Uid, tx.Error
	}
	return advice.Aid, nil
}

func (AdviceDao) CreatePhoto(aid int, photo string) error {
	var p = &models.Photo{Aid: aid, Photo: photo}
	tx := db.Create(&p)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (AdviceDao) CountAllAdvice() (int, error) {
	var count int64
	tx := db.Model(&models.Advice{}).Count(&count)
	if tx.Error != nil {
		log.Println(tx.Error)
		return 0, tx.Error
	}
	return int(count), nil
}

func (AdviceDao) GetAdvice() ([]models.Advice, error) {
	var advice []models.Advice
	tx := db.Order("status").Find(&advice)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return advice, nil
}

func (AdviceDao) GetAdviceByAid(aid int) (models.Advice, error) {
	var advice models.Advice
	tx := db.Where("aid=?", aid).First(&advice)
	if tx.Error != nil {
		return advice, tx.Error
	}
	return advice, nil
}

func (AdviceDao) GetPhotosByAid(aid int) ([]string, error) {
	var photos []string
	tx := db.Model(&models.Photo{}).Select("photo").Where("aid=?", aid).Find(&photos)
	if tx.Error != nil {
		return photos, tx.Error
	}
	return photos, nil
}

func (AdviceDao) UpdateAdvice(advice models.Advice) error {
	tx := db.Model(&models.Advice{}).Where("aid=?", advice.Aid).Updates(&advice)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (AdviceDao) DeleteAdviceByStatus() error {
	var count int
	tx := db.Model(&models.Advice{}).Where("status=?", 1).Delete(&count)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
