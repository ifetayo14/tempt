package controller

import (
	"agit/database"
	"agit/model"
	"time"
)

func Create(employee model.Employee) error {
	db := database.GetDB()
	return db.Debug().Create(&employee).Error
}

func Update(employeeId string, request model.Employee) error {
	db := database.GetDB()
	return db.Debug().Model(&model.Employee{}).Where("id = ?", employeeId).Updates(request).Error
}

func GetAll(request model.Pagination) (employee []model.Employee, total int64, err error) {
	db := database.GetDB()

	col := db.Debug().Model(&model.Employee{}).Where("deleted_at IS NULL").Order("created_at desc").Count(&total)

	err = col.Limit(request.Limit).Offset(request.Offset).Find(&employee).Error
	if err != nil {
		return nil, 0, err
	}

	return employee, total, err
}

func Detail(employeeId string) (employee model.Employee, err error) {
	db := database.GetDB()
	err = db.Debug().Model(&model.Employee{}).Where("id = ?", employeeId).First(&employee).Error
	if err != nil {
		return employee, err
	}

	return employee, nil
}

func Delete(employeeId string) error {
	db := database.GetDB()

	return db.Debug().Model(&model.Employee{}).Where("id = ?", employeeId).Update("deleted_at", time.Now()).Error
}
