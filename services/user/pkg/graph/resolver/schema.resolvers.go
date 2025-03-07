package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.66

import (
	"auth"
	"context"
	"errorwrapper"
	"user/pkg/ent"
	"user/pkg/graph/gen"
	"user/pkg/graph/gen/graphqlmodel"
)

// Signup is the resolver for the signup field.
func (r *mutationResolver) Signup(ctx context.Context, input ent.CreateUserInput) (*graphqlmodel.AuthPayload, error) {
	entClient := ent.FromContext(ctx)
	return r.userService.Signup(ctx, entClient, input)
}

// Login is the resolver for the login field.
func (r *mutationResolver) Login(ctx context.Context, email string, password string) (*graphqlmodel.AuthPayload, error) {
	entClient := ent.FromContext(ctx)
	return r.userService.Login(ctx, entClient, email, password)
}

// RefreshToken is the resolver for the refreshToken field.
func (r *mutationResolver) RefreshToken(ctx context.Context, refreshToken string) (*graphqlmodel.AuthPayload, error) {
	return r.userService.RefreshToken(ctx, r.entClient, refreshToken)
}

// UpdateUser is the resolver for the updateUser field.
func (r *mutationResolver) UpdateUser(ctx context.Context, id int, input ent.UpdateUserInput) (*ent.User, error) {
	passport := auth.PassportFromContext(ctx)
	if !passport.UserAuthority.IsAdmin() && !passport.UserAuthority.IsUser() {
		return nil, &errorwrapper.HTTPError{
			StatusCode: 401,
			Message:    "UpdateUser should have role admin or user",
		}
	}
	if passport.UserId != id {
		return nil, &errorwrapper.HTTPError{
			StatusCode: 403,
			Message:    "User ID does not match",
		}
	}
	entClient := ent.FromContext(ctx)
	return r.userService.UpdateUser(ctx, entClient, id, input)
}

// DeleteUser is the resolver for the deleteUser field.
func (r *mutationResolver) DeleteUser(ctx context.Context, id int) (bool, error) {
	passport := auth.PassportFromContext(ctx)
	if !passport.UserAuthority.IsAdmin() {
		return false, &errorwrapper.HTTPError{
			StatusCode: 401,
			Message:    "DeleteUser should have role admin",
		}
	}
	entClient := ent.FromContext(ctx)
	return r.userService.DeleteUser(ctx, entClient, id)
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id int) (*ent.User, error) {
	return r.userService.FindUser(ctx, r.entClient, id)
}

// Mutation returns gen.MutationResolver implementation.
func (r *Resolver) Mutation() gen.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
