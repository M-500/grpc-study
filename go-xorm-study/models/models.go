package models

import (
	"time"
)

type User struct {
	Id         int       `xorm:"not null pk autoincr comment('主键ID自增长') INT"`
	Phone      string    `xorm:"not null comment('手机号，用于登录，唯一值') CHAR(11)"`
	NickName   string    `xorm:"comment('昵称') VARCHAR(32)"`
	Age        int       `xorm:"comment('年龄') INT"`
	Pwd        string    `xorm:"not null comment('密码(密文存储)') VARCHAR(128)"`
	Email      []byte    `xorm:"comment('邮箱') VARBINARY(32)"`
	Mark       string    `xorm:"comment('备注') VARCHAR(64)"`
	CreateTime time.Time `xorm:"not null comment('创建时间') DATETIME"`
	UpdateTime time.Time `xorm:"comment('更新时间') DATETIME"`
	DeleteTime time.Time `xorm:"comment('删除时间') DATETIME"`
}
