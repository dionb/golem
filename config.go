package crudinator

type Config struct {
	PersistentStore PersistentStoreConfig
	EventSink       EventSinkConfig
	Oauth           OauthConfig
}

type PersistentStoreConfig struct {
	Engine   string //type of db. Eg: "postgres", "mysql",
	Host     string
	Port     string
	Protocol string
	Username string
	Password string
	Schema   string

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
