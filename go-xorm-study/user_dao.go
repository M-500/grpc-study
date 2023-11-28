package go_xorm_study

import (
	"context"
	"grpc-study/go-xorm-study/models"
	"time"
	"xorm.io/xorm"
)

//
// @Description
// @Author 代码小学生王木木
// @Date 2023/11/27 17:15
//

type UserDao struct {
	db  *xorm.Engine
	ctx context.Context
}

func NewUserDao(ctx context.Context) *UserDao {
	return &UserDao{
		db:  GetDb(),
		ctx: ctx,
	}
}

func (dao *UserDao) Get(id int) (*models.User, error) {
	data := &models.User{}
	_, err := dao.db.ID(id).Get(data)
	if err != nil {
		return nil, err
	}
	if data == nil || data.Id == 0 {
		return nil, nil
	}
	return data, nil
}

func (dao *UserDao) FindByPhone(phone string) (*models.User, error) {
	data := &models.User{}
	sess := dao.db.Where("`phone`=?", phone) // 预编译方式，避免直接拼接SQL语句造成SQL注入
	_, err := sess.Get(data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (dao *UserDao) FindAllPager(page, size int) ([]models.User, int64, error) {
	if page < 1 {
		page = 1
	}
	if size < 1 {
		size = 10
	}
	start := (page - 1) * size
	dataList := make([]models.User, 0)
	total, err := dao.db.Desc("id").Limit(size, start).FindAndCount(&dataList)
	return dataList, total, err
}

func (dao *UserDao) Insert(data *models.User) error {
	data.CreateTime = time.Now()
	data.CreateTime = time.Now()
	_, err := dao.db.Insert(data)
	return err
}

func (dao *UserDao) Update(data *models.User, musColumns ...string) error {
	sess := dao.db.ID(data.Id)
	if len(musColumns) > 0 {
		sess.MustCols(musColumns...)
	}
	data.CreateTime = time.Now()
	_, err := sess.Update(data)
	return err
}

// 更高一层的封装
func (dao *UserDao) Save(data *models.User, musColumns ...string) error {
	if data.Id > 0 {
		return dao.Update(data, musColumns...)
	}
	return dao.Insert(data)
}
