package respository

import (
	"fmt"
	"gorm.io/gorm"
	"ptc/internal/application/command"
	"ptc/internal/model"
)

// 通过用户 telephone 来查询 uid
func GetUID(telephone string) int {
	var account model.UserRegister
	db.Where("telephone = ?", telephone).First(&account)
	return account.UID
}

// 检查用户是否已存在
func CheckAccount(telephone string) bool {
	var account model.UserRegister
	db.Where("telephone = ?", telephone).First(&account)
	return account.UID > 0
}

// 添加新用户
func AddAccount(account *command.Account) error {
	//开启事务
	return db.Transaction(func(tx *gorm.DB) error {
		// 在事务中执行一些 db 操作（从这里开始，您应该使用 'tx' 而不是 'db'）
		var err error
		userRegister := model.UserRegister{
			PassWord:  account.Password,
			Telephone: account.Telephone,
		}
		err = tx.Create(&userRegister).Error
		if err != nil {
			return err
		}

		err = tx.Create(&model.UserRelatedData{
			UserId:         userRegister.UID,
			FollowerNumber: 0,
			FansNumber:     0,
			PostNumber:     0,
			CollectNumber:  0,
			ForwardNumber:  0,
		}).Error

		if err != nil {
			return err
		}

		err = tx.Create(&model.UserDetails{
			UserId:     userRegister.UID,
			NickName:   account.NickName,
			Sex:        account.Sex,
			ProfileUrl: account.ProfileUrl,
		}).Error

		if err != nil {
			fmt.Println(err.Error())
			return err
		}
		// 返回 nil 提交事务
		return nil
	})
}

// 查询用户密码
func QueryAccountByTelephone(telephone string) (string, int, error) {
	var account model.UserRegister
	err := db.Where("telephone = ?", telephone).Find(&account).Error
	return account.PassWord, account.UID, err
}
