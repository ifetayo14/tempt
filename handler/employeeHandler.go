package handler

import (
	"agit/controller"
	"agit/helpers"
	"agit/model"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/labstack/gommon/log"
	"math"
	"net/http"
	"strconv"
	"time"
)

func Create(ctx *gin.Context) {
	request := model.CreateRequest{}

	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		log.Errorf("error binding request: ", err)
		ctx.JSON(http.StatusInternalServerError, helpers.JSONResponse{
			Message: err.Error(),
		})
		return
	}

	if err := controller.Create(model.Employee{
		Id:          uuid.New().String(),
		Name:        request.Name,
		NIP:         request.NIP,
		POB:         request.POB,
		DOB:         request.DOB,
		Age:         request.Age,
		Address:     request.Address,
		Religion:    request.Religion,
		Gender:      request.Gender,
		PhoneNumber: request.PhoneNumber,
		Email:       request.Email,
		CreatedAt:   time.Now(),
	}); err != nil {
		log.Errorf("error create employee: ", err)
		ctx.JSON(http.StatusInternalServerError, helpers.JSONResponse{
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, helpers.JSONResponse{
		Message: "OK",
	})
}

func Update(ctx *gin.Context) {
	request := model.UpdateRequest{}
	empId := ctx.Param("id")

	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		log.Errorf("error binding request: ", err)
		ctx.JSON(http.StatusInternalServerError, helpers.JSONResponse{
			Message: err.Error(),
		})
		return
	}

	data, err := controller.Detail(empId)
	if err != nil {
		log.Errorf("error binding request: ", err)
		ctx.JSON(http.StatusInternalServerError, helpers.JSONResponse{
			Message: err.Error(),
		})
		return
	}

	if request.Name != "" {
		data.Name = request.Name
	}
	if request.NIP != "" {
		data.NIP = request.NIP
	}
	if request.POB != "" {
		data.POB = request.POB
	}
	if request.DOB.String() != "" {
		data.DOB = request.DOB
	}
	if request.Age != 0 {
		data.Age = request.Age
	}
	if request.Address != "" {
		data.Address = request.Address
	}
	if request.Religion != "" {
		data.Religion = request.Religion
	}
	if request.Gender != "" {
		data.Gender = request.Gender
	}
	if request.PhoneNumber != "" {
		data.PhoneNumber = request.PhoneNumber
	}
	if request.Email != "" {
		data.Email = request.Email
	}
	data.UpdatedAt = time.Now()

	if err := controller.Update(empId, data); err != nil {
		log.Errorf("error update data: ", err)
		ctx.JSON(http.StatusInternalServerError, helpers.JSONResponse{
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, data)
}

func GetAll(ctx *gin.Context) {
	request := model.Pagination{}

	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		log.Errorf("error binding request: ", err)
		ctx.JSON(http.StatusInternalServerError, helpers.JSONResponse{
			Message: err.Error(),
		})
		return
	}

	data, total, err := controller.GetAll(request)
	if err != nil {
		log.Errorf("error get data: ", err)
		ctx.JSON(http.StatusInternalServerError, helpers.JSONResponse{
			Message: err.Error(),
		})
		return
	}

	page := math.Ceil(float64(total) / float64(request.Limit))
	ctx.Header("Pagination-Rows", strconv.Itoa(int(total)))
	ctx.Header("Pagination-Page", strconv.Itoa(int(page)))
	ctx.Header("Pagination-Limit", strconv.Itoa(request.Limit))

	ctx.JSON(http.StatusOK, data)
}

func Detail(ctx *gin.Context) {
	empId := ctx.Param("id")
	resp, err := controller.Detail(empId)
	if err != nil {
		log.Errorf("error get data: ", err)
		ctx.JSON(http.StatusInternalServerError, helpers.JSONResponse{
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func Delete(ctx *gin.Context) {
	empId := ctx.Param("id")
	err := controller.Delete(empId)
	if err != nil {
		log.Errorf("error get data: ", err)
		ctx.JSON(http.StatusInternalServerError, helpers.JSONResponse{
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, helpers.JSONResponse{
		Message: "OK",
	})
}
