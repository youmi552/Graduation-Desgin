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

var goodsService service.GoodsService
var historyService service.HistoryService

// 获得所有商品展示信息
func GetGoodsData(c *gin.Context) {
	goods, err := goodsService.GetGoodsData()
	if err != nil {
		c.JSON(http.StatusOK, Resp(e.GETGOODSDATAFAIL, err))
		return
	}
	c.JSON(http.StatusOK, Resp(e.SUCCESS, goods))
}

// 分页获得所有商品展示信息
func GetGoodsDataByPage(c *gin.Context) {
	var pageParams models.PageParams
	err := c.ShouldBind(&pageParams)
	if err != nil {
		c.JSON(http.StatusOK, Resp(e.GETPARAMFAIL, nil))
		return
	}

	goods, err := goodsService.GetGoodsByPage(pageParams.PageNumber, pageParams.PageSize, pageParams.Keyword)
	if err != nil {
		c.JSON(http.StatusOK, Resp(e.GETGOODSDATAFAIL, err))
		return
	}
	c.JSON(http.StatusOK, Resp(e.SUCCESS, goods))
}

// GetGoodsDataByCategory
func GetGoodsDataByCategory(c *gin.Context) {
	var categoryParams models.CategoryParams
	err := c.ShouldBind(&categoryParams)
	if err != nil {
		c.JSON(http.StatusOK, Resp(e.GETPARAMFAIL, nil))
		return
	}

	goods, err := goodsService.GetGoodsByCategory(categoryParams.PageNumber, categoryParams.PageSize, categoryParams.Cid)
	if err != nil {
		c.JSON(http.StatusOK, Resp(e.GETGOODSDATAFAIL, err))
		return
	}
	c.JSON(http.StatusOK, Resp(e.SUCCESS, goods))
}

// 上传商品
func UploadGoods(c *gin.Context) {
	id, _ := c.Get("uid")
	uid, err := strconv.Atoi(util.GetInterfaceToString(id))
	if err != nil {
		c.JSON(http.StatusOK, Resp(e.STRCONVERROR, nil))
		return
	}
	var goods models.Goods
	goods.Uid = uid
	err = c.ShouldBind(&goods)
	if err != nil {
		c.JSON(http.StatusOK, Resp(e.INVAILDPARAMS, err))
		return
	}
	goodsInfo, err := goodsService.UploadGoods(goods)
	if err != nil {
		c.JSON(http.StatusOK, Resp(e.UPLOADGOODSFAIL, err.Error()))
		return
	}
	c.JSON(http.StatusOK, Resp(e.SUCCESS, goodsInfo))
}

// 商品上传图片
func UploadPictures(c *gin.Context) {
	gid, _ := c.GetPostForm("gid")
	gid2, err := strconv.Atoi(gid)
	fmt.Println(gid, gid2)
	if err != nil {
		c.JSON(http.StatusOK, Resp(e.STRCONVERROR, err.Error()))
		return
	}

	f, err := c.FormFile("f1")
	if err != nil {
		c.JSON(http.StatusOK, Resp(e.UPLOADPICTUREFAIL, err.Error()))
		return
	}
	dst := path.Join("./pictures", f.Filename)
	c.SaveUploadedFile(f, dst)
	serverAddress := viper.GetString("server.address")
	picture := fmt.Sprintf("%s/%s", serverAddress, dst)
	err = goodsService.UploadPicture(gid2, picture)
	if err != nil {
		c.JSON(http.StatusOK, Resp(e.UPLOADPICTUREFAIL, err.Error()))
		return
	}
	c.JSON(http.StatusOK, Resp(e.SUCCESS, picture))
}

// 获得商品详细信息
func GetGoodsInfo(c *gin.Context) {
	id, _ := c.Get("uid")
	uid, err := strconv.Atoi(util.GetInterfaceToString(id))
	gid := c.Param("gid")
	gid2, err := strconv.Atoi(gid)
	if err != nil {
		c.JSON(http.StatusOK, Resp(e.STRCONVERROR, err.Error()))
		return
	}
	goodsInfo, err := goodsService.GetGoodsInfo(gid2, uid)
	if err != nil {
		c.JSON(http.StatusOK, Resp(e.GETGOODSINFOFAIL, err.Error()))
		return
	}
	go historyService.AddHistory(uid, gid2)
	c.JSON(http.StatusOK, Resp(e.SUCCESS, goodsInfo))
}

// 获得某一用户的商品数据
func GetGoodsByUid(c *gin.Context) {
	var pageParams models.PageParams
	err := c.ShouldBind(&pageParams)
	if err != nil {
		c.JSON(http.StatusOK, Resp(e.GETPARAMFAIL, nil))
		return
	}
	uid := c.Param("uid")
	uid2, err := strconv.Atoi(uid)
	if err != nil {
		c.JSON(http.StatusOK, Resp(e.STRCONVERROR, err.Error()))
		return
	}
	goods, err := goodsService.GetGoodsDataByUid(uid2, pageParams.PageNumber, pageParams.PageSize)
	if err != nil {
		c.JSON(http.StatusOK, Resp(e.GETGOODSDATAFAIL, err))
		return
	}
	c.JSON(http.StatusOK, Resp(e.SUCCESS, goods))
}

// 收藏/取消收藏商品
func Collection(c *gin.Context) {
	id, _ := c.Get("uid")
	uid, err := strconv.Atoi(util.GetInterfaceToString(id))
	if err != nil {
		c.JSON(http.StatusOK, Resp(e.STRCONVERROR, nil))
		return
	}
	var collectionParams models.CollectionParams
	err = c.ShouldBind(&collectionParams)
	if err != nil {
		c.JSON(http.StatusOK, Resp(e.GETPARAMFAIL, nil))
		return
	}
	if collectionParams.IsCollection {
		collection, err := goodsService.CancelCollection(uid, collectionParams.Gid)
		if err != nil {
			c.JSON(http.StatusOK, Resp(e.CANCELCOLLECTIONFAIL, err.Error()))
			return
		}
		c.JSON(200, Resp(e.SUCCESS, collection))
	} else {
		collection, err := goodsService.Collection(uid, collectionParams.Gid)
		if err != nil {
			c.JSON(http.StatusOK, Resp(e.COLLECTIONFAIL, err.Error()))
			return
		}
		c.JSON(200, Resp(e.SUCCESS, collection))
	}

}
func GetCollection(c *gin.Context) {
	id, _ := c.Get("uid")
	uid, err := strconv.Atoi(util.GetInterfaceToString(id))
	if err != nil {
		c.JSON(http.StatusOK, Resp(e.STRCONVERROR, err.Error()))
		return
	}
	goods, err := goodsService.GetUserCollection(uid)
	if err != nil {
		c.JSON(http.StatusOK, Resp(e.GETCOLLECTIONFAIL, err.Error()))
		return
	}
	c.JSON(http.StatusOK, Resp(e.SUCCESS, goods))
}
