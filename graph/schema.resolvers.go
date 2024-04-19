package graph

import (
	"context"
	"github.com/sajagsubedi/ContactLync/graph/model"
	"github.com/sajagsubedi/ContactLync/database"
)

var db = database.ConnectDb()

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input *model.CreateUserInput) (*model.User, error) {
    return db.CreateUser(input),nil
}

// UpdateUser is the resolver for the updateUser field.
func (r *mutationResolver) UpdateUser(ctx context.Context, input *model.UpdateUserInput) (*model.User, error) {    return db.UpdateUser(input),nil
}

// DeleteUser is the resolver for the deleteUser field.
func (r *mutationResolver) DeleteUser(ctx context.Context, id string) (*model.User, error) {
      return db.DeleteUser(id),nil
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
      return db.Users(),nil
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
      return db.User(id),nil
}

// UserByFilter is the resolver for the userByFilter field.
func (r *queryResolver) UserByFilter(ctx context.Context, input *model.FilterInput) ([]*model.User, error) {
    return db.UserByFilter(input),nil}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
