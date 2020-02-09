package service

import (
    "github.com/tapokshot/microslab/internal/app/microslab/database"
    "github.com/tapokshot/microslab/internal/app/microslab/model"
)

type User struct {
    store database.Store
}

func NewUserService(store database.Store) User {
    return User{store: store}
}

func (u *User) GetUserById(id int) (*model.User, error) {
    result := &model.User{}
    err := u.store.DB.QueryRow("SELECT id, first_name FROM public.user where id = $1", id).Scan(
        &result.ID,
        &result.FirstName,
    )
    return result, err
}

func (u *User) SaveUser(user *model.AddUserRequest) error {
    _, err := u.store.DB.Exec("INSERT INTO public.user (first_name, last_name, phone) VALUES ($1, $2, $3)",
        user.FirstName, user.LastName, user.Phone)
    return err
}
