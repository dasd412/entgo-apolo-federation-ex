package service

import (
	"context"
	"entgo.io/contrib/entgql"
	"user/internal/repository"
	"user/pkg/ent"
	"user/pkg/ent/user"
)

type (
	userService struct {
		userRepository repository.UserRepository
	}

	UserService interface {
		FindUser(
			ctx context.Context,
			client *ent.Client,
			id int,
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
		CreateUser(
			ctx context.Context,
			client *ent.Client,
			input ent.CreateUserInput,
		) (*ent.User, error)
		UpdateUser(
			ctx context.Context,
			client *ent.Client,
			id int,
			input ent.UpdateUserInput,
		) (*ent.User, error)
		DeleteUser(ctx context.Context, client *ent.Client, id int) (
			bool,
			error,
		)
	}
)

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService{
		userRepository: userRepository,
	}
}

func (s *userService) FindUser(
	ctx context.Context,
	client *ent.Client,
	id int,
) (*ent.User, error) {
	return s.userRepository.FindOne(
		ctx, client, func(query *ent.UserQuery) {
			query.Where(user.ID(id))
		},
	)
}

func (s *userService) Paginate(ctx context.Context, client *ent.Client, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, where *ent.UserWhereInput) (*ent.UserConnection, error) {
	return s.userRepository.Paginate(ctx, client, after, first, before, last, where)
}

func (s *userService) CreateUser(
	ctx context.Context,
	client *ent.Client,
	input ent.CreateUserInput,
) (*ent.User, error) {
	return s.userRepository.CreateOne(
		ctx, client, input,
	)
}

func (s *userService) UpdateUser(
	ctx context.Context,
	client *ent.Client,
	id int,
	input ent.UpdateUserInput,
) (*ent.User, error) {
	return s.userRepository.UpdateOne(
		ctx, client, id, input,
	)
}

func (s *userService) DeleteUser(
	ctx context.Context,
	client *ent.Client,
	id int,
) (bool, error) {
	var success = false

	err := s.userRepository.DeleteOne(ctx, client, id)

	if err == nil {
		success = true
	}

	return success, err
}
