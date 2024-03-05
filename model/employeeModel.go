package model

import "time"

type (
	Employee struct {
		Id          string    `json:"id" gorm:"column:id"`
		Name        string    `json:"name" gorm:"column:name"`
		NIP         string    `json:"nip" gorm:"column:nip"`
		POB         string    `json:"pob" gorm:"column:pob"`
		DOB         time.Time `json:"dob" gorm:"column:dob"`
		Age         int       `json:"age" gorm:"column:age"`
		Address     string    `json:"address" gorm:"column:address"`
		Religion    string    `json:"religion" gorm:"column:religion"`
		Gender      string    `json:"gender" gorm:"column:gender"`
		PhoneNumber string    `json:"phone_number" gorm:"column:phone_number"`
		Email       string    `json:"email" gorm:"column:email"`
		CreatedAt   time.Time `json:"created_at" gorm:"column:created_at"`
		UpdatedAt   time.Time `json:"updated_at" gorm:"column:updated_at"`
		DeletedAt   time.Time `json:"deleted_at" gorm:"column:deleted_at"`
	}

	CreateRequest struct {
		Id          string    `json:"id"`
		Name        string    `json:"name"`
		NIP         string    `json:"nip"`
		POB         string    `json:"pob"`
		DOB         time.Time `json:"dob"`
		Age         int       `json:"age"`
		Address     string    `json:"address"`
		Religion    string    `json:"religion"`
		Gender      string    `json:"gender"`
		PhoneNumber string    `json:"phone_number"`
		Email       string    `json:"email"`
	}

	UpdateRequest struct {
		Name        string    `json:"name" gorm:"column:name"`
		NIP         string    `json:"nip" gorm:"column:nip"`
		POB         string    `json:"pob" gorm:"column:pob"`
		DOB         time.Time `json:"dob" gorm:"column:dob"`
		Age         int       `json:"age" gorm:"column:age"`
		Address     string    `json:"address" gorm:"column:address"`
		Religion    string    `json:"religion" gorm:"column:religion"`
		Gender      string    `json:"gender" gorm:"column:gender"`
		PhoneNumber string    `json:"phone_number" gorm:"column:phone_number"`
		Email       string    `json:"email" gorm:"column:email"`
		UpdatedAt   time.Time `gorm:"column:updated_at"`
	}

	Pagination struct {
		Limit  int `json:"limit"`
		Offset int `json:"offset"`
	}
)

func (Employee) TableName() string {
	return "employee"
}
