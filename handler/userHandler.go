package handler

import (
	"agit/controller"
	"agit/helpers"
	"agit/model"
	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
	"net/http"
)

func Register(ctx *gin.Context) {
	request := model.RegisterRequest{}

	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		log.Errorf("error binding request: ", err)
		ctx.JSON(http.StatusInternalServerError, helpers.JSONResponse{
			Message: err.Error(),
		})
		return
	}

	if err := controller.Register(request); err != nil {
		log.Errorf("error register user: ", err)
		if err.Error() == helpers.ERRUSERNAMEEXISTS {
			ctx.JSON(http.StatusBadRequest, helpers.JSONResponse{
				Message: helpers.ERRUSERNAMEEXISTS,
			})
			return
		}

		ctx.JSON(http.StatusInternalServerError, helpers.JSONResponse{
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, helpers.JSONResponse{
		Message: "OK",
	})
}
