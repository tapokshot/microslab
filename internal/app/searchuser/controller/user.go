package controller

import (
    "context"
    "fmt"
    model "github.com/tapokshot/microslab/internal/app/grpc"
)

type SearchUserController struct {
}

func (s *SearchUserController) Search(
    ctx context.Context,
    in *model.SearchUserRequest) (*model.SearchUserResponse, error) {

    fmt.Print("Start handle request")
    fmt.Print(in)
    fmt.Print("Finish handle request")

    return &model.SearchUserResponse{City: "Moscow"}, nil
}
