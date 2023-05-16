package service

import (
	"GraduationDesign/dao"
	"GraduationDesign/models"
)

type AdviceService struct {
}

var adviceDao dao.AdviceDao

func (AdviceService) AddAdvice(advice models.Advice) (int, error) {
	aid, err := adviceDao.AddAdvice(advice)
	if err != nil {
		return 0, err
	}
	return aid, nil
}

func (AdviceService) UploadPhoto(aid int, photo string) error {
	err := adviceDao.CreatePhoto(aid, photo)
	if err != nil {
		return err
	}
	return nil
}

func (AdviceService) GetAdviceByAid(aid int) (models.AdviceDetail, error) {
	var adviceDetail models.AdviceDetail
	advice, err := adviceDao.GetAdviceByAid(aid)
	if err != nil {
		return adviceDetail, nil
	}
	adviceDetail.Aid = advice.Aid
	user, err := userDao.GetUserByUid(advice.Uid)
	if err != nil {
		return adviceDetail, err
	}
	adviceDetail.UserName = user.UserName
	categoryName, err := categoryDao.GetAdviceCategoryNameByCid(advice.Cid)
	if err != nil {
		return adviceDetail, err
	}
	photos, err := adviceDao.GetPhotosByAid(advice.Aid)
	if err != nil {
		return adviceDetail, err
	}
	adviceDetail.Category = categoryName
	adviceDetail.Title = advice.Title
	adviceDetail.Status = advice.Status
	adviceDetail.CreatedAt = advice.CreatedAt.Format("2006-01-02 15:04:05")
	adviceDetail.Detail = advice.Detail
	adviceDetail.PhoneNumber = advice.PhoneNumber
	adviceDetail.Pictures = photos
	return adviceDetail, nil
}
