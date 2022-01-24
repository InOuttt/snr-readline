package config

import (
	"context"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// Implementation for Database Client
type MysqlSession struct {
	Db     *sql.DB
	config DbConfig
}

// New Mysql Session
func NewMysqlSession(config DbConfig) (*MysqlSession, error) {
	db, err := sql.Open("mysql", config.DSN)
	if err != nil {
		print("Unable to Connect DB to %v", err)
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &MysqlSession{
		Db:     db,
		config: config,
	}, nil
}

func (m *MysqlSession) NewCtx() (context.Context, context.CancelFunc) {
	return NewCtx(m.config.Timeout)
}
