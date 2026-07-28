package user

import (
	"context"
	"errors"
	"time"

	Models "API-COCKPIT-GYM/Models"
	repository "API-COCKPIT-GYM/Repository"
)

func CreateUser(ctx context.Context, payload Models.User) (Models.User, error) {
	if payload.ID == "" {
		return Models.User{}, errors.New("id is required")
	}
	if payload.Nome == "" {
		return Models.User{}, errors.New("nome is required")
	}
	if payload.Email == "" {
		return Models.User{}, errors.New("email is required")
	}

	payload.CreatedAt = time.Now()
	payload.UpdatedAt = payload.CreatedAt

	if err := repository.CreateUser(ctx, payload); err != nil {
		return Models.User{}, err
	}
	return payload, nil
}

func UpdateUser(ctx context.Context, payload Models.User) (Models.User, error) {
	if payload.ID == "" {
		return Models.User{}, errors.New("id is required")
	}

	payload.UpdatedAt = time.Now()

	if err := repository.UpdateUser(ctx, payload); err != nil {
		return Models.User{}, err
	}
	return payload, nil
}

func DeleteUser(ctx context.Context, id string) error {
	if id == "" {
		return errors.New("id is required")
	}
	return repository.DeleteUser(ctx, id)
}

func FindUserByID(ctx context.Context, id string) (*Models.User, error) {
	if id == "" {
		return nil, errors.New("id is required")
	}
	return repository.FindUserByID(ctx, id)
}

func FindAllUsers(ctx context.Context) (map[string]Models.User, error) {
	return repository.FindAllUsers(ctx)
}
