package controller

import (
	"GraduationDesign/e"
	"GraduationDesign/models"
	"GraduationDesign/service"
	"GraduationDesign/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
	"path"
	"strconv"
)

var adviceService service.AdviceService

func AddAdvice(c *gin.Context) {
	id, _ := c.Get("uid")
	uid, err := strconv.Atoi(util.GetInterfaceToString(id))
	if err != nil {
		c.JSON(http.StatusOK, Resp(e.STRCONVERROR, nil))
		return
	}
	var advice models.Advice
	advice.Uid = uid
	err = c.ShouldBind(&advice)
	if err != nil {
		c.JSON(http.StatusOK, Resp(e.INVAILDPARAMS, nil))
	}
	aid, err := adviceService.AddAdvice(advice)
	if err != nil {
		c.JSON(http.StatusOK, Resp(e.ADDADVICEFAIL, err.Error()))
		return
	}
	c.JSON(http.StatusOK, Resp(e.SUCCESS, aid))
}

func UploadPhotos(c *gin.Context) {
	id, _ := c.GetPostForm("aid")
	aid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusOK, Resp(e.STRCONVERROR, err.Error()))
		return
	}
	f, err := c.FormFile("f1")
	if err != nil {
		c.JSON(http.StatusOK, Resp(e.UPLOADPICTUREFAIL, err.Error()))
		return
	}
	dst := path.Join("./photos", f.Filename)
	c.SaveUploadedFile(f, dst)
	serverAddress := viper.GetString("server.address")
	photo := fmt.Sprintf("%s/%s", serverAddress, dst)
	err = adviceService.UploadPhoto(aid, photo)
	if err != nil {
		c.JSON(http.StatusOK, Resp(e.UPLOADPICTUREFAIL, err.Error()))
		return
	}
	c.JSON(http.StatusOK, Resp(e.SUCCESS, photo))
}
