package service

import (
	"GraduationDesign/dao"
	"GraduationDesign/models"
)

type IdentificationService struct {
}

var identificationDao dao.IdentificationDao

func (IdentificationService) GetIdentificationByUid(uid int) (int, error) {
	status, err := identificationDao.GetIdentificationByUid(uid)
	if err != nil {
		return -1, err
	}
	return status, nil
}

func (IdentificationService) AddIdentification(identification models.Identification) error {
	err := identificationDao.CreateIdentification(identification)
	if err != nil {
		return err
	}
	return nil
}

func (IdentificationService) UpdateIdentification(identification models.Identification) error {
	err := identificationDao.UpdateIdentification(identification)
	if err != nil {
		return err
	}
	return nil
}

func (IdentificationService) GetAllIdentification(params models.PageParams) (models.IdentificationResp, error) {
	count, err := identificationDao.CountIdentification(params.Keyword)
	if err != nil {
		return models.IdentificationResp{}, err
	}
	offset := (params.PageNumber - 1) * params.PageSize
	identifications, err := identificationDao.GetAllIdentification(offset, params.PageSize, params.Keyword)
	var identificationsData []models.IdentificationData
	for _, identification := range identifications {
		var identificationData models.IdentificationData
		user, err := userDao.GetUserByUid(identification.Uid)
		if err != nil {
			return models.IdentificationResp{}, err
		}
		identificationData.UserName = user.UserName
		identificationData.PhoneNumber = user.PhoneNumber
		identificationData.Status = identification.Status
		identificationData.Id = identification.Id
		identificationData.StudentId = identification.StudentId
		identificationData.Name = identification.Name
		identificationsData = append(identificationsData, identificationData)
	}
	return models.IdentificationResp{
		Count:              count,
		IdentificationData: identificationsData,
	}, nil

}

func (IdentificationService) AcceptIdentification(id int) error {
	err := identificationDao.UpdateStatusIdentification(id, 1)
	if err != nil {
		return err
	}
	return nil
}

func (IdentificationService) RefuseIdentification(id int) error {
	err := identificationDao.UpdateStatusIdentification(id, 2)
	if err != nil {
		return err
	}
	return nil
}
