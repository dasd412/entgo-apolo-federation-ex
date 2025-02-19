package db

import (
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type DBConfig struct {
	Username string
	Password string
	Hostname string
	DBName   string
	Port     string
}

func NewClients[T any](
	masterCfg, replicaCfg DBConfig,
	newClientFunc func(*sql.Driver, *sql.Driver) *T,
) *T {
	master := getDB(masterCfg)
	masterDriver, err := OpenDriver(master)
	if err != nil {
		log.Fatalf("failed opening connection to mysql(master): %v", err)
	}

	replica := getDB(replicaCfg)
	replicaDriver, err := OpenDriver(replica)
	if err != nil {
		log.Fatalf("failed opening connection to mysql(replica): %v", err)
	}

	return newClientFunc(masterDriver, replicaDriver)
}

func OpenDriver(conn Connection) (*sql.Driver, error) {
	//fmt.Printf("DSN: %s", conn.GetDSN())
	driver, err := sql.Open(dialect.MySQL, conn.GetDSN())
	if err != nil {
		return nil, err
	}

	return driver, nil
}

func getDB(config DBConfig) Connection {
	return newDatabase(config)
}
