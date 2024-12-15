package store

import "github.com/yikuanzz/social/internal/base/data"

type UserStore struct {
	data *data.Data
}

func NewUserStore(data *data.Data) *UserStore {
	return &UserStore{
		data: data,
	}
}
