package config

import (
	"context"
	"os"
	"strconv"
	"time"
)

// Database configuration
type DbConfig struct {
	DSN      string
	DbName   string
	PoolMin  int
	PoolMax  int
	IdleTime int
	Timeout  time.Duration
}

const (
	EnvDsn      = "DB_DSN"
	EnvDbName   = "DB_NAME"
	EnvPoolMin  = "DB_POOL_MIN"
	EnvPoolMax  = "DB_POOL_MAX"
	EnvIdleTime = "DB_MAX_IDLE_TIME_SECOND"
	EnvTimeout  = "DB_TIMEOUT"
)

// Interface for Database Session
// type DbSession interface {
// 	Open() error
// 	Ping() error
// 	Close() error
// }

func NewCtx(duration time.Duration) (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), duration)
}

// Load Db Configuration from env
func LoadDbConfig(prefix ...string) DbConfig {
	prf := ""
	if len(prefix) > 0 {
		prf = prefix[0]
	}

	envDsn := os.Getenv(prf + EnvDsn)
	envDb := os.Getenv(prf + EnvDbName)
	var err error
	var pMin, pMax, idleTime, timeout int
	pMin, err = strconv.Atoi(os.Getenv(prf + EnvPoolMin))
	if err != nil {
		pMin = 5
	}

	pMax, err = strconv.Atoi(os.Getenv(prf + EnvPoolMax))
	if err != nil {
		pMax = 10
	}

	idleTime, err = strconv.Atoi(os.Getenv(prf + EnvIdleTime))
	if err != nil {
		idleTime = 30
	}

	timeout, err = strconv.Atoi(os.Getenv(prf + EnvTimeout))
	if err != nil {
		timeout = 2
	}

	return DbConfig{
		DSN:      envDsn,
		DbName:   envDb,
		PoolMin:  pMin,
		PoolMax:  pMax,
		IdleTime: idleTime,
		Timeout:  time.Duration(timeout) * time.Second,
	}
}
