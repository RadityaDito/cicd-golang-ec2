package repositories

import (
    "golang-echo-crud/config"
    "golang-echo-crud/models"
)

func GetUsers() ([]models.User, error) {
    var users []models.User
    err := config.DB.Find(&users).Error
    return users, err
}

func GetUserByID(id uint) (models.User, error) {
    var user models.User
    err := config.DB.First(&user, id).Error
    return user, err
}

func CreateUser(user *models.User) error {
    return config.DB.Create(user).Error
}

func UpdateUser(user *models.User) error {
    return config.DB.Save(user).Error
}

func DeleteUser(id uint) error {
    return config.DB.Delete(&models.User{}, id).Error
}
