package internal

import (
	"github.com/codebyshennan/mockapi/domain"
)

// Read env variables
func ReadEnv() (m *domain.Config, err error) {
	m = &domain.Config{
		MongoUrl:       "mongodb://localhost:27017",
		DbName:         "sandbox-staging",
		JwtKey:         "123",
		AppEnv:         "development",
		GoogleClientId: "257544260761-mjmgllu9gvvbt4vut09urq4ridltl80m.apps.googleusercontent.com",
		DisableAuth:    false,
	}
	// keys := []string{
	// 	"JWT_SIGNING_KEY",
	// 	"DATABASE_URL",
	// 	"DATABASE_NAME",
	// 	"APP_ENV",
	// }

	// for _, key := range keys {
	// 	if k := os.Getenv(key); k != "" {
	// 		m[key] = k
	// 	} else {
	// 		err = errors.New("ERR_INIT")
	// 		return
	// 	}
	// }
	// return
	return
}
