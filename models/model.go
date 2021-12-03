package models

import (
	"encoding/json"
	"fmt"
	"go-gin-api/config"
)

type Model interface {
	getCollectionName() string
}

func getModelAttributes(model Model) map[string]interface{} {
	json_value, _ := json.Marshal(model)
	output := make(map[string]interface{})
	json.Unmarshal(json_value, &output)
	fmt.Println(output)
	return output
}

func AddToCollection(model Model, connection *config.FirestoreConnection) {
	connection.FirestoreAdd(model.getCollectionName(), getModelAttributes(model))
}

func FindAll(model Model, connection *config.FirestoreConnection) []config.FirestoreData {
	return connection.FirestoreGetAll(model.getCollectionName())
}
