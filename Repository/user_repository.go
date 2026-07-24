package repository

import (
	"context"
	"fmt"

	Models "API-COCKPIT-GYM/Models"
)

const userPath = "users"

func CreateUser(ctx context.Context, user Models.User) error {
	return Create(ctx, userPath, user.ID, user)
}

func UpdateUser(ctx context.Context, user Models.User) error {
	return Update(ctx, userPath, user.ID, user)
}

func DeleteUser(ctx context.Context, id string) error {
	return Delete(ctx, userPath, id)
}

func FindUserByID(ctx context.Context, id string) (*Models.User, error) {
	var user Models.User
	if err := FindByID(ctx, userPath, id, &user); err != nil {
		return nil, err
	}
	if user.ID == "" {
		return nil, fmt.Errorf("user %s not found", id)
	}
	return &user, nil
}

func FindAllUsers(ctx context.Context) (map[string]Models.User, error) {
	users := make(map[string]Models.User)
	if err := FindAll(ctx, userPath, &users); err != nil {
		return nil, err
	}
	return users, nil
}
