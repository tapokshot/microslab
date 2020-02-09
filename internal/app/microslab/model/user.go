package model

type User struct {
    ID        int
    FirstName string
    LastName  string
    Phone     string
}

type AddUserRequest struct {
    FirstName string `json:"first_name"`
    LastName  string `json:"last_name"`
    Phone     string `json:"phone"`
}
