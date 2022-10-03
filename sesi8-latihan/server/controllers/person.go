package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"sesi8-latihan/server/models"
	"sesi8-latihan/server/views"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PersonController struct {
	db *gorm.DB
}

func NewPersonController(db *gorm.DB) *PersonController {
	return &PersonController{
		db: db,
	}
}

// GetPeople godoc
// @Summary    Get all people
// @Decription Get list people
// @Tags       person
// @Accept     json
// @Produce    json
// @Success    200 {object} views.GetAllPeopleSwagger
// @Router     /people [get]
func (p *PersonController) GetPeople(ctx *gin.Context) {
	var people []models.Person

	err := p.db.Find(&people).Error
	if err != nil {
		WriteJsonResponse(ctx, &views.Response{
			Status:  http.StatusInternalServerError,
			Message: "GET_PEOPLE_FAIL",
			Error:   err.Error(),
		})
		return
	}
	WriteJsonResponse(ctx, &views.Response{
		Status:  http.StatusOK,
		Message: "GET_PEOPLE_SUCCESS",
		Payload: people,
	})
}

// CreatePeople godoc
// @Summary    Create people
// @Decription Create new people
// @Tags       person
// @Accept     json
// @Produce    json
// @Param people body models.Person true "Create people"
// @Success    200 {object} views.CreatePeopleSwagger
// @Router     /people [post]
func (p *PersonController) CreatePeople(ctx *gin.Context) {
	body, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		WriteJsonResponse(ctx, &views.Response{
			Status:  http.StatusInternalServerError,
			Message: "CREATE_PEOPLE_FAIL",
			Error:   err.Error(),
		})
		return
	}
	person := models.Person{}
	err = json.Unmarshal(body, &person)
	if err != nil {
		WriteJsonResponse(ctx, &views.Response{
			Status:  http.StatusInternalServerError,
			Message: "CREATE_PEOPLE_FAIL",
			Error:   err.Error(),
		})
		return
	}
	err = p.db.Create(&person).Error
	if err != nil {
		WriteJsonResponse(ctx, &views.Response{
			Status:  http.StatusInternalServerError,
			Message: "CREATE_PEOPLE_FAIL",
			Error:   err.Error(),
		})
		return
	}

	WriteJsonResponse(ctx, &views.Response{
		Status:  http.StatusOK,
		Message: "CREATE_PEOPLE_SUCCESS",
		Payload: person,
	})
}

// UpdatePeople godoc
// @Summary    Update people
// @Decription Update existing people
// @Tags       person
// @Accept     json
// @Produce    json
// @Param people body models.Person true "Update people"
// @Success    200 {object} views.UpdatePeopleSwagger
// @Router     /people [put]
func (p *PersonController) UpdatePeople(ctx *gin.Context) {
	body, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		WriteJsonResponse(ctx, &views.Response{
			Status:  http.StatusInternalServerError,
			Message: "UPDATE_PEOPLE_FAIL",
			Error:   err.Error(),
		})
		return
	}
	person := models.Person{}
	err = json.Unmarshal(body, &person)
	if err != nil {
		WriteJsonResponse(ctx, &views.Response{
			Status:  http.StatusInternalServerError,
			Message: "UPDATE_PEOPLE_FAIL",
			Error:   err.Error(),
		})
		return
	}
	err = p.db.Where("Name = ?", person.Name).Updates(&person).Error
	if err != nil {
		WriteJsonResponse(ctx, &views.Response{
			Status:  http.StatusInternalServerError,
			Message: "UPDATE_PEOPLE_FAIL",
			Error:   err.Error(),
		})
		return
	}

	WriteJsonResponse(ctx, &views.Response{
		Status:  http.StatusOK,
		Message: "UPDATE_PEOPLE_SUCCESS",
		Payload: person,
	})
}

// DeletePeople godoc
// @Summary    Delete people
// @Decription Delete existing people
// @Tags       person
// @Accept     json
// @Produce    json
// @Param people body models.Person true "Delete people"
// @Success    200 {object} views.DeletePeopleSwagger
// @Router     /people [delete]
func (p *PersonController) DeletePeople(ctx *gin.Context) {
	body, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		WriteJsonResponse(ctx, &views.Response{
			Status:  http.StatusInternalServerError,
			Message: "DELETE_PEOPLE_FAIL",
			Error:   err.Error(),
		})
		return
	}
	person := models.Person{}
	err = json.Unmarshal(body, &person)
	if err != nil {
		WriteJsonResponse(ctx, &views.Response{
			Status:  http.StatusInternalServerError,
			Message: "DELETE_PEOPLE_FAIL",
			Error:   err.Error(),
		})
		return
	}
	err = p.db.Where("Name = ?", person.Name).Take(&models.Person{}).Delete(&models.Person{}).Error
	if err != nil {
		WriteJsonResponse(ctx, &views.Response{
			Status:  http.StatusInternalServerError,
			Message: "DELETE_PEOPLE_FAIL",
			Error:   err.Error(),
		})
		return
	}

	WriteJsonResponse(ctx, &views.Response{
		Status:  http.StatusOK,
		Message: "DELETE_PEOPLE_SUCCESS",
	})
}
