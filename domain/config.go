package domain

type Config struct {
	MongoUrl       string
	DbName         string
	JwtKey         string
	AppEnv         string
	GoogleClientId string
	DisableAuth    bool
}
