package repository

import (
	"context"
	"entgo.io/contrib/entgql"
	"user/pkg/ent"
)

type (
	userRepository struct {
	}

	UserRepository interface {
		FindOne(
			ctx context.Context,
			client *ent.Client,
			opts ...func(query *ent.UserQuery),
		) (*ent.User, error)
		Paginate(
			ctx context.Context,
			client *ent.Client,
			after *entgql.Cursor[int],
			first *int,
			before *entgql.Cursor[int],
			last *int,
			where *ent.UserWhereInput,
		) (*ent.UserConnection, error)
		CreateOne(
			ctx context.Context,
			client *ent.Client,
			input ent.CreateUserInput,
		) (*ent.User, error)
		UpdateOne(
			ctx context.Context,
			client *ent.Client,
			id int,
			input ent.UpdateUserInput,
		) (*ent.User, error)
		DeleteOne(ctx context.Context, client *ent.Client, id int) error
	}
)

func NewUserRepository() UserRepository {
	return &userRepository{}
}

func (u *userRepository) FindOne(
	ctx context.Context,
	client *ent.Client,
	opts ...func(query *ent.UserQuery),
) (*ent.User, error) {
	query := client.User.
		Query()

	for _, opt := range opts {
		opt(query)
	}

	return query.Only(ctx)
}

func (u *userRepository) Paginate(ctx context.Context, client *ent.Client, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, where *ent.UserWhereInput) (*ent.UserConnection, error) {
	return client.User.Query().
		Paginate(
			ctx,
			after,
			first,
			before,
			last,
			ent.WithUserFilter(where.Filter),
		)
}

func (u *userRepository) CreateOne(
	ctx context.Context,
	client *ent.Client,
	input ent.CreateUserInput,
) (*ent.User, error) {
	return client.User.
		Create().
		SetInput(input).
		Save(ctx)
}

func (u *userRepository) UpdateOne(ctx context.Context, client *ent.Client, id int, input ent.UpdateUserInput) (*ent.User, error) {
	return client.User.
		UpdateOneID(id).
		SetInput(input).
		Save(ctx)
}

func (u *userRepository) DeleteOne(ctx context.Context, client *ent.Client, id int) error {
	return client.User.
		DeleteOneID(id).
		Exec(ctx)
}
