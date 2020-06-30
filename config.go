package crudinator

type Config struct {
	PersistentStore PersistentStoreConfig
	EventSink       EventSinkConfig
	Oauth           OauthConfig
}

type PersistentStoreConfig struct {
	Host     string
	Port     string
	Protocol string
	Username string
	Password string

	ConnectionString string
}

type EventSinkConfig struct {
	Host     string
	Port     string
	Protocol string
}

type OauthConfig struct {
	Host     string
	Port     string
	Protocol string
}

func ParseConfig() Config {
	return Config{}
}
