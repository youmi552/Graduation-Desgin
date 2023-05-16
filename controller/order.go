package controller

import (
	"GraduationDesign/e"
	"GraduationDesign/models"
	"GraduationDesign/service"
	"GraduationDesign/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var orderService service.OrderService

func CreateOrder(c *gin.Context) {
	id, _ := c.Get("uid")
	uid, err := strconv.Atoi(util.GetInterfaceToString(id))
	if err != nil {
		c.JSON(http.StatusOK, Resp(e.STRCONVERROR, err.Error()))
		return
	}
	var order models.Order
	err = c.ShouldBind(&order)
	if err != nil {
		c.JSON(http.StatusOK, Resp(e.GETPARAMFAIL, nil))
		return
	}
	order.Uid = uid
	err = orderService.CreateOrder(order)
	if err != nil {
		c.JSON(http.StatusOK, Resp(e.CREATEORDERFAIL, err.Error()))
		return
	}
	go goodsService.UpdateGoodsStatus(order.Gid, 2)

	err = cartService.ClearCart(uid)
	c.JSON(http.StatusOK, Resp(e.SUCCESS, nil))

}
func GetOrders(c *gin.Context) {
	id, _ := c.Get("uid")
	uid, err := strconv.Atoi(util.GetInterfaceToString(id))
	if err != nil {
		c.JSON(http.StatusOK, Resp(e.STRCONVERROR, err.Error()))
		return
	}
	orders, err := orderService.GetOrdersByUid(uid)
	if err != nil {
		c.JSON(http.StatusOK, Resp(e.GETORDERSFAIL, err.Error()))
		return
	}
	fmt.Println(orders)
	c.JSON(http.StatusOK, Resp(e.SUCCESS, orders))
}

func GetMySell(c *gin.Context) {
	id, _ := c.Get("uid")
	uid, err := strconv.Atoi(util.GetInterfaceToString(id))
	if err != nil {
		c.JSON(http.StatusOK, Resp(e.STRCONVERROR, err.Error()))
		return
	}
	orders, err := orderService.GetOrdersByUid(uid)
	if err != nil {
		c.JSON(http.StatusOK, Resp(e.GETORDERSFAIL, err.Error()))
		return
	}
	c.JSON(http.StatusOK, Resp(e.SUCCESS, orders))
}

func GetOrderDetail(c *gin.Context) {
	id := c.Param("oid")
	oid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusOK, Resp(e.STRCONVERROR, err.Error()))
		return
	}
	orderDetail, err := orderService.GetOrderDetail(oid)
	if err != nil {
		c.JSON(http.StatusOK, Resp(e.GETORDERDETAILFAIL, err.Error()))
		return
	}
	c.JSON(http.StatusOK, Resp(e.SUCCESS, orderDetail))
}

func UpdateOrder(c *gin.Context) {
	var order models.Order
	err := c.ShouldBind(&order)
	if err != nil {
		c.JSON(http.StatusOK, Resp(e.GETPARAMFAIL, nil))
		return
	}
	err = orderService.UpdateOrder(order)
	if err != nil {
		c.JSON(http.StatusOK, Resp(e.GETORDERDETAILFAIL, err.Error()))
		return
	}
	c.JSON(http.StatusOK, Resp(e.SUCCESS, nil))

}
