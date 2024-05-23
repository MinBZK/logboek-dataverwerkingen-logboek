package main

import (
	"flag"
	"log"
	"os"
	"strings"

	"github.com/MinBZK/logboek-dataverwerkingen-logboek/server/pkg/server"
	"github.com/MinBZK/logboek-dataverwerkingen-logboek/server/pkg/storage"
)

type options struct {
	listenAddress    string
	storageType      string
	sqlitePath       string
	cassandraServers string
}

func parseFlags() *options {
	o := &options{}
	flag.StringVar(&o.listenAddress, "listen-address", envValue("LISTEN_ADDRESS", "127.0.0.1:9000"), "Listen address")
	flag.StringVar(&o.storageType, "storage-type", envValue("STORAGE_TYPE", "sqlite"), "Storage backend to use")
	flag.StringVar(&o.sqlitePath, "storage-sqlite-path", envValue("STORAGE_SQLITE_PATH", "logboek.db"), "Path to SQLite database file")
	flag.StringVar(&o.cassandraServers, "storage-cassandra-servers", envValue("STORAGE_CASSANDRA_SERVERS", "127.0.0.1:9042"), "List of Cassandra servers")

	flag.Parse()

	return o
}

func main() {
	o := parseFlags()

	ds := newDatastore(o)

	srv, err := server.New(ds)
	if err != nil {
		log.Fatalf("Creating server: %v", err)
	}

	log.Println("Starting server...")
	if err := srv.Start(o.listenAddress); err != nil {
		log.Fatal(err)
	}
}

func newDatastore(o *options) storage.Store {
	var (
		store storage.Store
		err   error
	)
	switch o.storageType {
	case "sqlite":
		store, err = storage.NewSqlite(o.sqlitePath)
		if err != nil {
			log.Fatalf("Initializing Sqlite storage: %v", err)
		}
	case "cassandra":
		servers := strings.Split(o.cassandraServers, ",")
		store, err = storage.NewCassandra(servers...)
		if err != nil {
			log.Fatalf("Initializing Cassandra storage: %v", err)
		}
	default:
		log.Fatalf("Unsupported datatstore: %s", o.storageType)
	}

	return store
}

func envValue(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
