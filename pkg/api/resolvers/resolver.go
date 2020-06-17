package resolvers

//go:generate go run github.com/99designs/gqlgen

import (
	"github.com/maxsuelmarinho/golang-graphql-todos/graph/generated"
	"github.com/maxsuelmarinho/golang-graphql-todos/pkg/api/model"
)

type Resolver struct {
	todos []*model.Todo
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Todo returns generated.TodoResolver implementation.
func (r *Resolver) Todo() generated.TodoResolver { return &todoResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
