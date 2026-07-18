package repository

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/db"
	"google.golang.org/api/option"

	Models "API-COCKPIT-GYM/Models"
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

func SalvarUsuario(ctx context.Context, user Models.User) error {

	ref := GetDB().NewRef("usuarios/" + user.Name)

	return ref.Set(ctx, user)
}
