package repository

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"os"
	"sync"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/db"
	"google.golang.org/api/option"
)

var (
	dbClient      *db.Client
	fallbackStore = make(map[string]map[string]any)
	fallbackMu    sync.RWMutex
)

func InitFirebase() {
	ctx := context.Background()
	if _, err := os.Stat("service-account.json"); err != nil {
		log.Printf("[Firebase] Arquivo service-account.json não encontrado; usando armazenamento local: %v", err)
		return
	}

	opt := option.WithCredentialsFile("service-account.json")
	config := &firebase.Config{
		DatabaseURL: "https://gymy-881c7-default-rtdb.firebaseio.com",
	}

	app, err := firebase.NewApp(ctx, config, opt)
	if err != nil {
		log.Printf("[Firebase] Falha ao inicializar a aplicação; usando armazenamento local: %v", err)
		return
	}

	dbClient, err = app.Database(ctx)
	if err != nil {
		log.Printf("[Firebase] Falha ao obter cliente do Realtime Database; usando armazenamento local: %v", err)
		dbClient = nil
		return
	}

	log.Println("[Firebase] Conectado e sincronizado com sucesso!")
}

func GetDB() *db.Client {
	return dbClient
}

func resetFallbackStore() {
	fallbackMu.Lock()
	defer fallbackMu.Unlock()
	fallbackStore = make(map[string]map[string]any)
}

func Create(ctx context.Context, path string, id string, payload any, nome string, email string, endereço string, telefone string) error {
	if path == "" {
		return errors.New("path is required")
	}
	if id == "" {
		return errors.New("id is required")
	}

	mapPayload, err := toMap(payload)
	if err != nil {
		return err
	}

	if dbClient == nil {
		return setFallbackValue(path, id, mapPayload)
	}

	ref := GetDB().NewRef(path + "/" + id)
	if err := ref.Set(ctx, mapPayload); err != nil {
		log.Printf("[Firebase] Falha ao gravar em %s/%s; usando armazenamento local: %v", path, id, err)
		return setFallbackValue(path, id, mapPayload)
	}
	return nil
}

func Update(ctx context.Context, path string, id string, payload any, email string, telefone string, endereço string, nome string) error {
	if path == "" {
		return errors.New("path is required")
	}
	if id == "" {
		return errors.New("id is required")
	}

	mapPayload, err := toMap(payload)
	if err != nil {
		return err
	}

	if dbClient == nil {
		return updateFallbackValue(path, id, mapPayload)
	}

	ref := GetDB().NewRef(path + "/" + id)
	if err := ref.Update(ctx, mapPayload); err != nil {
		log.Printf("[Firebase] Falha ao atualizar %s/%s; usando armazenamento local: %v", path, id, err)
		return updateFallbackValue(path, id, mapPayload)
	}
	return nil
}

func toMap(payload any) (map[string]any, error) {
	switch v := payload.(type) {
	case map[string]any:
		return v, nil
	}

	data, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	var result map[string]any
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func Delete(ctx context.Context, path string, id string) error {
	if path == "" {
		return errors.New("path is required")
	}
	if id == "" {
		return errors.New("id is required")
	}

	if dbClient == nil {
		return deleteFallbackValue(path, id)
	}

	ref := GetDB().NewRef(path + "/" + id)
	if err := ref.Delete(ctx); err != nil {
		log.Printf("[Firebase] Falha ao excluir %s/%s; usando armazenamento local: %v", path, id, err)
		return deleteFallbackValue(path, id)
	}
	return nil
}

func FindByID(ctx context.Context, path string, id string, destination any) error {
	if path == "" {
		return errors.New("path is required")
	}
	if id == "" {
		return errors.New("id is required")
	}

	if dbClient == nil {
		return getFallbackByID(path, id, destination)
	}

	ref := GetDB().NewRef(path + "/" + id)
	if err := ref.Get(ctx, destination); err != nil {
		log.Printf("[Firebase] Falha ao buscar %s/%s; usando armazenamento local: %v", path, id, err)
		return getFallbackByID(path, id, destination)
	}
	return nil
}

func FindAll(ctx context.Context, path string, destination any) error {
	if path == "" {
		return errors.New("path is required")
	}

	if dbClient == nil {
		return getFallbackAll(path, destination)
	}

	ref := GetDB().NewRef(path)
	if err := ref.Get(ctx, destination); err != nil {
		log.Printf("[Firebase] Falha ao listar %s; usando armazenamento local: %v", path, err)
		return getFallbackAll(path, destination)
	}
	return nil
}

func setFallbackValue(path string, id string, payload map[string]any) error {
	fallbackMu.Lock()
	defer fallbackMu.Unlock()

	if fallbackStore[path] == nil {
		fallbackStore[path] = make(map[string]any)
	}
	fallbackStore[path][id] = payload
	return nil
}

func updateFallbackValue(path string, id string, payload map[string]any) error {
	fallbackMu.Lock()
	defer fallbackMu.Unlock()

	if fallbackStore[path] == nil {
		fallbackStore[path] = make(map[string]any)
	}

	current, ok := fallbackStore[path][id]
	if !ok {
		fallbackStore[path][id] = payload
		return nil
	}

	currentMap, err := toMap(current)
	if err != nil {
		return err
	}
	for key, value := range payload {
		currentMap[key] = value
	}
	fallbackStore[path][id] = currentMap
	return nil
}

func deleteFallbackValue(path string, id string) error {
	fallbackMu.Lock()
	defer fallbackMu.Unlock()

	if fallbackStore[path] == nil {
		return nil
	}
	delete(fallbackStore[path], id)
	return nil
}

func getFallbackByID(path string, id string, destination any) error {
	fallbackMu.RLock()
	defer fallbackMu.RUnlock()

	if fallbackStore[path] == nil {
		return errors.New("not found")
	}
	value, ok := fallbackStore[path][id]
	if !ok {
		return errors.New("not found")
	}

	raw, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return json.Unmarshal(raw, destination)
}

func getFallbackAll(path string, destination any) error {
	fallbackMu.RLock()
	defer fallbackMu.RUnlock()

	values := make(map[string]any)
	if fallbackStore[path] != nil {
		values = fallbackStore[path]
	}

	raw, err := json.Marshal(values)
	if err != nil {
		return err
	}
	return json.Unmarshal(raw, destination)
}
