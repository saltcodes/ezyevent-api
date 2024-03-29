package queries

import (
	"ezyevent-api/internal/model"
	"fmt"
	"gorm.io/gorm/clause"
	"strings"
)

// CreateUser takes a model struct and saves it
func CreateUser(user *model.User) error {
	return (*db).Create(user).Error
}

// ListUser Queries the users tab;e and stored them in user list array
func ListUser(userList *[]model.User) error {
	return (*db).Find(&userList).Error
}

// GetUser Queries and return a single user
func GetUser(id string, user *model.User) error {
	return (*db).First(&user, "id=?", id).Error
}

// UpdateUser update users
func UpdateUser(id string, user *model.User) error {

	if strings.Compare(id, user.Id) != 0 {
		return fmt.Errorf("invalid id")
	}

	return (*db).Save(user).Error
}

// DeleteUser delete user
func DeleteUser(id string) error {
	var user model.User

	err := (*db).First(&user, "id=?", id).Error
	if user.Id == "" {
		return err
	}
	return (*db).Select(clause.Associations, "Comment").Delete(&user).Error
}
