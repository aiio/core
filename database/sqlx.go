package database

import (
	"fmt"
	"strings"

	"github.com/aiio/core/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

// InitSQLX 初始化一个 sqlx.DB
func InitSQLX(engine, dsn string) *sqlx.DB {
	db, err := sqlx.Connect(engine, dsn)
	if err != nil {
		panic(err)
	}
	return db
}

// InitSQLXEnv 根据环境变量 初始化一个 sqlx.DB
func InitSQLXEnv(envEngineName, envDsnName string) *sqlx.DB {
	return InitSQLX(config.GetEnv(envEngineName), config.GetEnv(envDsnName))
}

// InitSQLXTag 根据Tag环境变量 初始化一个 sqlx.DB
func InitSQLXTag(tag string) *sqlx.DB {
	return InitSQLXEnv(fmt.Sprintf("MYSQL_%v_ENGINE", strings.ToUpper(tag)),
		fmt.Sprintf("MYSQL_%v_DSN", strings.ToUpper(tag)),
	)
}
