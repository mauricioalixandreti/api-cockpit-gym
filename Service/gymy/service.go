package gymy

import (
	"context"
	"errors"
	"time"

	Models "API-COCKPIT-GYM/Models"
	repository "API-COCKPIT-GYM/Repository"
)

func CreateGym(ctx context.Context, payload Models.Gym) (Models.Gym, error) {
	if payload.ID == "" {
		return Models.Gym{}, errors.New("id is required")
	}
	if payload.Nome == "" {
		return Models.Gym{}, errors.New("nome is required")
	}
	if payload.Email == "" {
		return Models.Gym{}, errors.New("email is required")
	}
	if payload.Endereço == "" {
		return Models.Gym{}, errors.New("endereço is required")
	}
	if payload.Telefone == "" {
		return Models.Gym{}, errors.New("telefone is required")
	}

	payload.CreatedAt = time.Now()
	payload.UpdatedAt = payload.CreatedAt

	if err := repository.CreateGym(ctx, payload); err != nil {
		return Models.Gym{}, err
	}
	return payload, nil
}

func UpdateGym(ctx context.Context, payload Models.Gym) (Models.Gym, error) {
	if payload.ID == "" {
		return Models.Gym{}, errors.New("id is required")
	}
	if payload.Nome == "" {
		return Models.Gym{}, errors.New("nome is required")
	}

	payload.UpdatedAt = time.Now()

	if err := repository.UpdateGym(ctx, payload); err != nil {
		return Models.Gym{}, err
	}
	return payload, nil
}

func DeleteGym(ctx context.Context, id string) error {
	if id == "" {
		return errors.New("id is required")
	}
	return repository.DeleteGym(ctx, id)
}

func FindGymByID(ctx context.Context, id string) (*Models.Gym, error) {
	if id == "" {
		return nil, errors.New("id is required")
	}
	return repository.FindGymByID(ctx, id)
}

func FindAllGyms(ctx context.Context) (map[string]Models.Gym, error) {
	return repository.FindAllGyms(ctx)
}

func FindGymOwnerEmail(email string) (*Models.Gym, error) {

	// Validar a busca da academia pelo email do proprietario

	if email == "" {
		return nil, errors.New("email obrigatório")
	}

	return repository.FindGymOwnerEmail(email)

}
