package database

import "github.com/yesleymiranda/go-account/src/users"

// DeclareModels declare all models for auto migrations
func DeclareModels() []interface{} {
	models := make([]interface{}, 0)
	models = append(models, &users.User{})
	return models
}
