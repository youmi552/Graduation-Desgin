package service

import (
	"GraduationDesign/dao"
	"GraduationDesign/models"
	"time"
)

type NoticeService struct {
}

var noticeDao dao.NoticeDao

// 获取公告
func (NoticeService) GetNotice() (models.Notice, error) {
	notices, err := noticeDao.GetNotice()
	if err != nil {
		return models.Notice{}, err
	}
	return models.Notice{Date: notices[0], Content: notices[1]}, nil
}

// 更新公告
func (NoticeService) UpdateNotice(date string, content string) (models.Notice, error) {
	nowTime := time.Now()
	date2 := nowTime.Format("2006-01-02")
	notice := models.Notice{Date: date2, Content: content}
	err := noticeDao.UpdateNotice(date, notice)
	if err != nil {
		return models.Notice{}, err
	}
	return notice, nil
}
