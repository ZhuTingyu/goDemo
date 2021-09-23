package user

import (
	"demo/database"
	"fmt"
	"github.com/jinzhu/gorm"
)

type Dao struct {
	database.Dao
}

func (dao *Dao) GetUser(id string) (*User, error) {
	fmt.Println(id)
	var user User
	err := dao.GetDB().Table("user").Where("id = ?", id).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}
