package handler

import (
	"agit/controller"
	"agit/helpers"
	"agit/model"
	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
	"net/http"
)

func Login(ctx *gin.Context) {
	request := model.Request{}
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		log.Errorf("error binding request: ", err)
		ctx.JSON(http.StatusInternalServerError, helpers.JSONResponse{
			Message: err.Error(),
		})
		return
	}

	userData, err := controller.GetUserByUsername(request.Username)
	if err != nil {
		log.Errorf("error getUser: ", err)
		ctx.JSON(http.StatusInternalServerError, helpers.JSONResponse{
			Message: err.Error(),
		})

		return
	}

	if userData == nil {
		ctx.JSON(http.StatusBadRequest, helpers.JSONResponse{
			Message: "user is not exist",
		})
		return
	}

	comparePass := helpers.ComparePass([]byte(userData.Password), []byte(request.Password))
	if !comparePass {
		ctx.JSON(http.StatusUnauthorized, helpers.JSONResponse{
			Message: "invalid email or password",
		})
		return
	}

	token := helpers.GenerateToken(userData.Id.String(), userData.Username)

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
