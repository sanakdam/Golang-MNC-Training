package controllers

import (
	"Golang-MNC-Training/Kode-Assignment-2/server/models"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type OrderController struct {
	db *gorm.DB
}

func NewOrderController(db *gorm.DB) *OrderController {
	return &OrderController{
		db: db,
	}
}

func (o *OrderController) GetOrders(ctx *gin.Context) {
	var order []models.Order

	err := o.db.Preload("Items").Find(&order).Error
	if err != nil {
		WriteJsonResponse(ctx, &Response{
			Status:  http.StatusInternalServerError,
			Message: "GET_ORDER_FAIL",
			Error:   err.Error(),
		})
		return
	}
	WriteJsonResponse(ctx, &Response{
		Status:  http.StatusOK,
		Message: "GET_ORDER_SUCCESS",
		Payload: order,
	})
}

func (o *OrderController) CreateOrder(ctx *gin.Context) {
	body, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		WriteJsonResponse(ctx, &Response{
			Status:  http.StatusInternalServerError,
			Message: "CREATE_ORDER_FAIL",
			Error:   err.Error(),
		})
		return
	}
	order := models.Order{}
	err = json.Unmarshal(body, &order)
	if err != nil {
		WriteJsonResponse(ctx, &Response{
			Status:  http.StatusInternalServerError,
			Message: "CREATE_ORDER_FAIL",
			Error:   err.Error(),
		})
		return
	}

	order.Prepare()
	err = order.Validate()
	if err != nil {
		WriteJsonResponse(ctx, &Response{
			Status:  http.StatusInternalServerError,
			Message: "CREATE_ORDER_FAIL",
			Error:   err.Error(),
		})
		return
	}

	err = o.db.Create(&order).Error
	if err != nil {
		WriteJsonResponse(ctx, &Response{
			Status:  http.StatusInternalServerError,
			Message: "CREATE_ORDER_FAIL",
			Error:   err.Error(),
		})
		return
	}

	WriteJsonResponse(ctx, &Response{
		Status:  http.StatusOK,
		Message: "CREATE_ORDER_SUCCESS",
		Payload: order,
	})
}

func (o *OrderController) UpdateOrder(ctx *gin.Context) {
	oid, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		WriteJsonResponse(ctx, &Response{
			Status:  http.StatusInternalServerError,
			Message: "UPDATE_ORDER_FAIL",
			Error:   err.Error(),
		})
		return
	}

	order := models.Order{}
	err = o.db.Where("id = ?", oid).Take(&order).Error
	if err != nil {
		WriteJsonResponse(ctx, &Response{
			Status:  http.StatusInternalServerError,
			Message: "UPDATE_ORDER_FAIL",
			Error:   err.Error(),
		})
		return
	}

	body, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		WriteJsonResponse(ctx, &Response{
			Status:  http.StatusInternalServerError,
			Message: "UPDATE_ORDER_FAIL",
			Error:   err.Error(),
		})
		return
	}

	orderUpdate := models.Order{}
	err = json.Unmarshal(body, &orderUpdate)
	if err != nil {
		WriteJsonResponse(ctx, &Response{
			Status:  http.StatusInternalServerError,
			Message: "UPDATE_ORDER_FAIL",
			Error:   err.Error(),
		})
		return
	}

	orderUpdate.ID = order.ID
	orderUpdate.Prepare()
	err = orderUpdate.Validate()
	if err != nil {
		WriteJsonResponse(ctx, &Response{
			Status:  http.StatusInternalServerError,
			Message: "UPDATE_ORDER_FAIL",
			Error:   err.Error(),
		})
		return
	}

	err = o.db.Where("id = ?", order.ID).Updates(&orderUpdate).Error
	if err != nil {
		WriteJsonResponse(ctx, &Response{
			Status:  http.StatusInternalServerError,
			Message: "UPDATE_ORDER_FAIL",
			Error:   err.Error(),
		})
		return
	}

	WriteJsonResponse(ctx, &Response{
		Status:  http.StatusOK,
		Message: "UPDATE_ORDER_SUCCESS",
		Payload: orderUpdate,
	})
}

func (o *OrderController) DeleteOrder(ctx *gin.Context) {
	oid, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		WriteJsonResponse(ctx, &Response{
			Status:  http.StatusInternalServerError,
			Message: "DELETE_ORDER_FAIL",
			Error:   err.Error(),
		})
		return
	}

	// Check if the post exist
	order := models.Order{}
	err = o.db.Where("id = ?", oid).Take(&order).Error
	if err != nil {
		WriteJsonResponse(ctx, &Response{
			Status:  http.StatusInternalServerError,
			Message: "DELETE_ORDER_FAIL",
			Error:   err.Error(),
		})
		return
	}

	err = o.db.Where("id = ?", order.ID).Take(&models.Order{}).Delete(&models.Order{}).Error
	if err != nil {
		WriteJsonResponse(ctx, &Response{
			Status:  http.StatusInternalServerError,
			Message: "DELETE_ORDER_FAIL",
			Error:   err.Error(),
		})
		return
	}

	WriteJsonResponse(ctx, &Response{
		Status:  http.StatusOK,
		Message: "DELETE_ORDER_SUCCESS",
	})
}
