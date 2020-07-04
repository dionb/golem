package crudinator

import "net/http"

// ValidatorFunc describes the format for validator functions. The annotation used to call the validator is provided to enable implementing variations within the same validator (eg: dates)
type ValidatorFunc func(value interface{}, annotation string) error

type EventSink interface {
	Send(interface{}) error
}

type PersistentStore interface {
	// Connect will use the provided config to connect to the persistent store provider
	Connect(PersistentStoreConfig) error

	// Session will be called before handling each request. It should return a new PersistentStore with the object that should be used for querying properly initialised. This is present to support connection pooling, but is not required to be used. If session returns nil then the original PersistentStore will be used
	Session() PersistentStore

	// Get returns the object identified by the given key from the table with the supplied name, and saves the result into the supplied dst object pointer
	Get(key interface{}, tableName string, dst interface{}) error

	// List returns a list of objects filtered using key:value pairs. The interface passed into List will always be a pointer to a slice of the reource to be listed
	List(tableName string, filters map[string]interface{}, dst interface{}) error

	// Insert will create a new entry using the provided key and value and will return an error if the operation has failed.
	Insert(key interface{}, tableName string, value interface{}) error

	// Insert will create a new entry using the provided key and value and overwrite the existing value if it does exist, and will return an error if the operation has failed.
	Set(key interface{}, tableName string, value interface{}) error

	// Update will provide partial modification of objects in a json-patch style update (TBD)
	Update(key interface{}, tableName string, value interface{}) error

	// Delete will remove the entry associated with the provided key
	Delete(key interface{}, tableName string) error

	// Close will close the underlying connection
	Close() error

	// Raw will return a pointer to the underlying db connection
	Raw() interface{}
}

type AuthProvider interface {
}

type StdGetHandler interface {
	Get(rw http.ResponseWriter, req *http.Request)
}

type CRUDGetHandler interface {
	Get(rw http.ResponseWriter, req *http.Request, ctx Context)
}

type StdListHandler interface {
	List(rw http.ResponseWriter, req *http.Request)
}
