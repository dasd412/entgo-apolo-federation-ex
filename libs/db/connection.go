package db

import (
	"fmt"
)

type (
	connection struct {
		username string
		password string
		hostname string
		dbName   string
		port     string
	}

	Connection interface {
		GetDSN() string
	}

	Database struct {
		connection
	}
)

func (conn *connection) GetDSN() string {
	const format = "%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true&loc=Local"
	return fmt.Sprintf(
		format,
		conn.username,
		conn.password,
		conn.hostname,
		conn.port,
		conn.dbName,
	)
}

func newDatabase(cfg DBConfig) Connection {

	return &Database{
		connection: connection{
			username: cfg.Username,
			password: cfg.Password,
			hostname: cfg.Hostname,
			dbName:   cfg.DBName,
			port:     cfg.Port,
		},
	}
}
