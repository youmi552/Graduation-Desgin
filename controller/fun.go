package controller

import (
	"GraduationDesign/e"
	"GraduationDesign/models"
)

func Resp(code uint, data interface{}) models.Resp {
	return models.Resp{
		Code: code,
		Msg:  e.GetMsg(code),
		Data: data,
	}
}
