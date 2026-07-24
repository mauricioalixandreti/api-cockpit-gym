package repository

import (
	"context"
	"errors"
	"log"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/db"
	"google.golang.org/api/option"
)

var dbClient *db.Client

func InitFirebase() {
	ctx := context.Background()
	opt := option.WithCredentialsFile("service-account.json")

	config := &firebase.Config{
		DatabaseURL: "https://gymy-881c7-default-rtdb.firebaseio.com",
	}

	app, err := firebase.NewApp(ctx, config, opt)
	if err != nil {
		log.Fatalf("[Firebase Erro] Falha ao inicializar a aplicação: %v", err)
	}

	dbClient, err = app.Database(ctx)
	if err != nil {
		log.Fatalf("[Firebase Erro] Falha ao obter cliente do Realtime Database: %v", err)
	}

	log.Println("[Firebase] Conectado e sincronizado com sucesso!")
}

func GetDB() *db.Client {
	if dbClient == nil {
		log.Fatalln("[Firebase Erro] Tentativa de uso do banco antes de sua inicialização.")
	}
	return dbClient
}

func Create(ctx context.Context, path string, id string, payload any) error {
	if path == "" {
		return errors.New("path is required")
	}
	if id == "" {
		return errors.New("id is required")
	}

	ref := GetDB().NewRef(path + "/" + id)
	return ref.Set(ctx, payload)
}

func Update(ctx context.Context, path string, id string, payload any) error {
	if path == "" {
		return errors.New("path is required")
	}
	if id == "" {
		return errors.New("id is required")
	}

	mapPayload, ok := payload.(map[string]any)
	if !ok {
		return errors.New("payload must be a map[string]any")
	}

	ref := GetDB().NewRef(path + "/" + id)
	return ref.Update(ctx, mapPayload)
}

func Delete(ctx context.Context, path string, id string) error {
	if path == "" {
		return errors.New("path is required")
	}
	if id == "" {
		return errors.New("id is required")
	}

	ref := GetDB().NewRef(path + "/" + id)
	return ref.Delete(ctx)
}

func FindByID(ctx context.Context, path string, id string, destination any) error {
	if path == "" {
		return errors.New("path is required")
	}
	if id == "" {
		return errors.New("id is required")
	}

	ref := GetDB().NewRef(path + "/" + id)
	return ref.Get(ctx, destination)
}

func FindAll(ctx context.Context, path string, destination any) error {
	if path == "" {
		return errors.New("path is required")
	}

	ref := GetDB().NewRef(path)
	return ref.Get(ctx, destination)
}
