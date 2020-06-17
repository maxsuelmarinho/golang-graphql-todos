package resolvers

import (
	"context"

	"github.com/maxsuelmarinho/golang-graphql-todos/pkg/api/model"
)

func (r *todoResolver) User(ctx context.Context, obj *model.Todo) (*model.User, error) {
	return &model.User{ID: obj.UserID, Name: "user " + obj.UserID}, nil
}
