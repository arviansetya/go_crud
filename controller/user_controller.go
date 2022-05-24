package controller

import (
	"crud/config"
)

type Users struct {
	ID     string `json:"id" form:"id" gorm:"primaryKey"`
	Name   string `json:"name" form:"name"`
	Email  string `json:"email" form:"email"`
	Phone  string `json:"phone" form:"phone"`
	Alamat string `json:"alamat" form:"alamat"`
}

func GetUserById(id string) (Users, error) {
	var user Users
	result := config.DB.Where("id = ?", id).First(&user)
	return user, result.Error
}
func GetUserByAll(alldata string) ([]Users, error) {
	var user []Users
	result := config.DB.Where("id LIKE ? OR name LIKE ?", "%"+alldata+"%", "%"+alldata+"%").Find(&user)

	return user, result.Error
}

func (user *Users) CreateData() error {
	if err := config.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (user *Users) UpdateData(id string) error {
	if err := config.DB.Model(&Users{}).Where("id = ?", id).Updates(user).Error; err != nil {
		return err
	}
	return nil
}

func (user *Users) DeleteData() error {
	if err := config.DB.Delete(user).Error; err != nil {
		return err
	}
	return nil
}
