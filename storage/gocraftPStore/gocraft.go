package gocraftPStore

import (
	"fmt"

	"github.com/dionb/crudinator"
	"github.com/gocraft/dbr"
)

type GocraftdbrPersistentStore struct {
	conn    *dbr.Connection
	session *dbr.Session
}

func New() *GocraftdbrPersistentStore {
	return &GocraftdbrPersistentStore{}
}

func (pps *GocraftdbrPersistentStore) Connect(conf crudinator.PersistentStoreConfig) error {
	dsn := ""
	if conf.ConnectionString != "" {
		dsn = conf.ConnectionString
	} else {
		if conf.Host == "" {
			conf.Host = "localhost"
		}
		switch conf.Engine {
		case "postgres":
			if conf.Port == "" {
				conf.Port = "5432"
			}
			dsn = fmt.Sprintf(
				"Host=%s;Port=%s;Database=%s;User ID=%s;Password=%s;sslMode=Preferred;",
				conf.Password,
				conf.Username,
				conf.Host,
				conf.Port,
				conf.Schema,
			)
		case "mysql":
			dsn = fmt.Sprintf(
				"Server=%s;Port=%s;Database=%s;Uid=%s;Pwd=%s;SslMode=Preferred;",
				conf.Host,
				conf.Port,
				conf.Schema,
				conf.Username,
				conf.Password,
			)
		case "sqlite":
		}
	}

	conn, err := dbr.Open(conf.Engine, dsn, nil)
	if err != nil {
		return err
	}

	pps.conn = conn
	return nil
}

// func (pps *GocraftdbrPersistentStore) Session() crudinator.PersistentStore {
// 	return &GocraftdbrPersistentStore{
// 		session: pps.conn.NewSession(nil),
// 	}
// }

func (pps *GocraftdbrPersistentStore) Get(key interface{}, tableName string, dst interface{}) error {
	_, err := pps.session.Select("*").From(tableName).Where("id = ?", key).Load(&dst)
	return err
}

func (pps *GocraftdbrPersistentStore) List(tableName string, filters map[string]interface{}, dst interface{}) error {
	q := pps.session.Select("*").From(tableName)
	for k, v := range filters {
		q = q.Where("? = ?", k, v)
	}
	_, err := q.Load(&dst)
	return err
}

func (pps *GocraftdbrPersistentStore) Insert(key interface{}, tableName string, value interface{}) error {
	_, err := pps.session.InsertInto(tableName).Record(value).Exec()
	return err
}

func (pps *GocraftdbrPersistentStore) Set(key interface{}, tableName string, value interface{}) error {
	err := pps.Insert(key, tableName, value)
	if err != nil {
		return pps.Update(key, tableName, value)
	}
	return nil
}

func (pps *GocraftdbrPersistentStore) Update(key interface{}, tableName string, value interface{}) error {

	return nil
}

func (pps *GocraftdbrPersistentStore) Delete(key interface{}, tableName string) error {
	_, err := pps.session.DeleteFrom(tableName).Where("id = ?", key).Exec()
	return err
}

func (pps *GocraftdbrPersistentStore) Close() error {
	return pps.session.Close()
}

func (pps *GocraftdbrPersistentStore) Raw() interface{} {
	return pps.session
}
