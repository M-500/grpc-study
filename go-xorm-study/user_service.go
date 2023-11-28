package go_xorm_study

import (
	"context"
	"grpc-study/go-xorm-study/models"
)

//
// @Description
// @Author 代码小学生王木木
// @Date 2023/11/27 17:15
//

type UserService struct {
	ctx        context.Context
	daoUserDao *UserDao
}

func NewUserService(ctx context.Context) *UserService {
	return &UserService{
		ctx:        ctx,
		daoUserDao: NewUserDao(ctx),
	}
}

func (s *UserService) Get(id int) (*models.User, error) {
	return s.daoUserDao.Get(id)
}

func (s *UserService) FindByPhone(phone string) (*models.User, error) {
	return s.daoUserDao.FindByPhone(phone)
}

func (s *UserService) FindAllPager(page, size int) ([]models.User, int64, error) {
	return s.daoUserDao.FindAllPager(page, size)
}

// 更高一层的封装
func (s *UserService) Save(data *models.User, musColumns ...string) error {
	return s.daoUserDao.Save(data, musColumns...)
}
