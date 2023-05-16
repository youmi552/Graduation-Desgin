package dao

import "GraduationDesign/models"

type IdentificationDao struct {
}

func (IdentificationDao) GetIdentificationByUid(uid int) (int, error) {
	var status int
	tx := db.Model(&models.Identification{}).Select("status").Where("uid=?", uid).First(&status)

	if tx.Error != nil {
		if tx.Error.Error() == "record not found" {
			return -1, nil
		}
		return status, tx.Error
	}
	return status, nil
}

func (IdentificationDao) CreateIdentification(identification models.Identification) error {
	tx := db.Create(&identification)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (IdentificationDao) UpdateIdentification(identification models.Identification) error {
	db.Model(&models.Identification{}).Where("uid=?", identification.Uid).Update("status", 0)
	tx := db.Where("uid=?", identification.Uid).Updates(&identification)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (IdentificationDao) CountIdentification(keyword string) (int, error) {
	var count int64
	tx := db.Model(&models.Identification{}).Where("status=?", keyword).Count(&count)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return int(count), nil
}

func (IdentificationDao) GetAllIdentification(offset int, pageSize int, keyword string) ([]models.Identification, error) {
	var identifications []models.Identification
	tx := db.Where("status=?", keyword).Offset(offset).Limit(pageSize).Find(&identifications)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return identifications, nil
}

func (IdentificationDao) UpdateStatusIdentification(id int, status int) error {
	tx := db.Model(&models.Identification{}).Where("id=?", id).Update("status", status)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
