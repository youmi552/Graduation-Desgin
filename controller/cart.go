package controller

import (
	"GraduationDesign/e"
	"GraduationDesign/service"
	"GraduationDesign/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var cartService service.CartService

func GetUserCart(c *gin.Context) {
	id, _ := c.Get("uid")
	uid, err := strconv.Atoi(util.GetInterfaceToString(id))
	if err != nil {
		c.JSON(http.StatusOK, Resp(e.STRCONVERROR, err.Error()))
		return
	}
	goods, err := cartService.GetUserCart(uid)
	if err != nil {
		c.JSON(http.StatusOK, Resp(e.GETCARTFAIL, err.Error()))
		return
	}
	c.JSON(http.StatusOK, Resp(e.SUCCESS, goods))
}

func AddGoodInCart(c *gin.Context) {
	id, _ := c.Get("uid")
	uid, err := strconv.Atoi(util.GetInterfaceToString(id))
	if err != nil {
		c.JSON(http.StatusOK, Resp(e.STRCONVERROR, err.Error()))
		return
	}
	gd := c.Param("gid")
	gid, err := strconv.Atoi(gd)
	if err != nil {
		c.JSON(http.StatusOK, Resp(e.STRCONVERROR, err.Error()))
		return
	}
	err = cartService.AddGoodInCart(uid, gid)
	if err != nil {
		c.JSON(http.StatusOK, Resp(e.ADDGOODINCARTFAIL, err.Error()))
		return
	}
	c.JSON(http.StatusOK, Resp(e.SUCCESS, nil))
}

func RemoveGoodInCart(c *gin.Context) {
	id, _ := c.Get("uid")
	uid, err := strconv.Atoi(util.GetInterfaceToString(id))
	if err != nil {
		c.JSON(http.StatusOK, Resp(e.STRCONVERROR, err.Error()))
		return
	}
	gd := c.Param("gid")
	gid, err := strconv.Atoi(gd)
	if err != nil {
		c.JSON(http.StatusOK, Resp(e.STRCONVERROR, err.Error()))
		return
	}
	err = cartService.RemoveGoodInCart(uid, gid)
	if err != nil {
		c.JSON(http.StatusOK, Resp(e.REMOVEGOODINCARTFAIL, err.Error()))
		return
	}
	c.JSON(http.StatusOK, Resp(e.SUCCESS, nil))
}
