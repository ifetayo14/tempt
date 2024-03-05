package controller

import (
	"agit/database"
	"agit/helpers"
	"agit/model"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/labstack/gommon/log"
	"time"
)

func Register(request model.RegisterRequest) error {
	db := database.GetDB()

	user := model.User{
		Id:        uuid.New(),
		Username:  request.Username,
		Password:  helpers.HashPass(request.Password),
		CreatedAt: time.Now(),
	}

	fmt.Println(user)

	//if err := checkUniqueUsername(request.Username); err != nil {
	//	return err
	//}

	return db.Create(&user).Error
}

func GetUserByUsername(param string) (*model.User, error) {
	user := model.User{}
	db := database.GetDB()

	err := db.Model(&model.User{}).Where("username = ?", param).Where("deleted_at IS NOT NULL").
		First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func checkUniqueUsername(username string) error {
	exist := ""
	db := database.GetDB()

	err := db.Select("username").Table("user").Where("username = ?", username).Find(&exist).Error
	if err != nil {
		log.Errorf(err.Error())
		return err
	}

	if exist != "" {
		return errors.New(helpers.ERRUSERNAMEEXISTS)
	}

	return nil
}
