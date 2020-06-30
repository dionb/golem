# crudinator
Make writing CRUD APIs a thing of the past

The goal of this project is to make it so that you only ever have to specify a crud api and not implement it, while allowing extensibility to handle corner cases. Some features of this will be planned for rapid development only, but most are intended to be mature enough to use in a production system. A good way of thinking about this would be that we want to wrap storage solutions with auth* and event sourcing for monitoring/auditing.

Planned features:
 - A highly configurable integration for oidc
 - Event sourcing with Kafka, RabbitMQ, and maybe some Cloud-provider specific solutions
 - Storage engine support for Postgres, Mysql, SQLite, MongoDB, CockroachDB, Reddis, etcd, CouchDB, S3?, maybe more?
 - Configuration via OpenAPI specifications, scraping an existing database, HTTP interface with schema stored in the supporting db.
 - Automated management of schema in the connected storage engine. This will probably be done through integration with existing db-specific libraries
 - Vault integration for pulling credentials for oauth, event sink and db connections
 - I want to investigate other transport mechanisms, in particular grpc, but that is dependent on the complexity of implementing and maintaining the different protocol support

Initially this will require type definitions to be written in go, with annotations for validators. Custom request handlers can be added by implementing request functions for the specific task you wish to modify. 