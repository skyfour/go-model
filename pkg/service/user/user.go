package user

import (
	"fmt"
	"github.com/skyfour/go-model/pkg/model"
	"github.com/skyfour/go-model/pkg/utils"
)

func GetUser(uid string) model.Users {
	var user model.Users
	if err := utils.DB.Where("id = ?", uid).First(&user).Error; err != nil {
		//user not found, return error message

	}
	return user
}

func InsertUser(user model.Users) interface{} {
	t := utils.DB.Create(&user)
	fmt.Printf("%v\n",(*t).Value)
	return (*t).Value
}

func DeleteUser(uid string) {
	utils.DB.Where("id = ?",uid).Delete(model.Users{})
}
