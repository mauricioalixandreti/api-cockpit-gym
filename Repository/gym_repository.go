package repository

import (
	"context"
	"fmt"

	Models "API-COCKPIT-GYM/Models"
)

const gymPath = "gyms"

func CreateGym(ctx context.Context, gym Models.Gym) error {
	return Create(ctx, gymPath, gym.ID, gym)
}

func UpdateGym(ctx context.Context, gym Models.Gym) error {
	return Update(ctx, gymPath, gym.ID, gym, "gym.Nome", "gym.Endereço", "gym.Endereço", "gym.Telefone")
}

func DeleteGym(ctx context.Context, id string) error {
	return Delete(ctx, gymPath, id)
}

func FindGymByID(ctx context.Context, id string) (*Models.Gym, error) {
	var gym Models.Gym
	if err := FindByID(ctx, gymPath, id, &gym); err != nil {
		return nil, err
	}
	if gym.ID == "" {
		return nil, fmt.Errorf("gym %s not found", id)
	}
	return &gym, nil
}

func FindAllGyms(ctx context.Context) (map[string]Models.Gym, error) {
	s := make(map[string]Models.Gym)
	if err := FindAll(ctx, gymPath, &s); err != nil {
		return nil, err
	}
	return s, nil
}
