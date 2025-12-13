package database

import "test/models"

var userList []models.User

func CreateUser(user models.User) {
	userList = append(userList, user)
}

func GetUserByEmail(email string) *models.User {
	for i := range userList {
		if userList[i].Email == email {
			return &userList[i]
		}
	}

	return nil
}

func CheckPasswordCorrect(user models.User, password string) bool {
	if user.Password == password {
		return true
	}
	return false
}

func GetUserCount() int {
	return len(userList)
}
