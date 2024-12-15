package service

import (
	"github.com/yikuanzz/social/internal/store"
)

type UserService struct {
	userStore *store.UserStore
}

func NewUserService(userStore *store.UserStore) *UserService {
	return &UserService{
		userStore: userStore,
	}
}
