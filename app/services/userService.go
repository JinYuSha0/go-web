package services

import (
	"go-web/app/core"
	"go-web/dao"
	"go-web/models"

	"github.com/kataras/iris"
)

type UserService interface {
	IsExists(account string) (iris.Map, error)
	Register(user models.UserRegister) (iris.Map, error)
	Login(account string, password string) (iris.Map, error)
}

type userService struct {
	DAO dao.UserDAO
}

func NewUserService(userDAO dao.UserDAO) UserService {
	return &userService{
		DAO: userDAO,
	}
}

func (s *userService) Register(user models.UserRegister) (res iris.Map, err error) {
	exists, err := s.DAO.IsExists(user.Account)

	core.CheckError(err)

	if exists {
		res = core.GenErrorRes(1001, "already register", nil)
		return
	}

	err = s.DAO.Register(user.Account, user.Password)

	core.CheckError(err)

	return
}

func (s *userService) Login(account string, password string) (res iris.Map, err error) {
	// s.DAO.Register(account, password)
	return
}

func (s *userService) IsExists(account string) (res iris.Map, err error) {
	var exists bool
	exists, err = s.DAO.IsExists(account)

	res = core.GenSuccessRes(exists)

	return
}
