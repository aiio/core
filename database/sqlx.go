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

// InitSqlx 初始化一个 sqlx.DB
func InitSqlx(engine, dsn string) *sqlx.DB {
	db, err := sqlx.Connect(engine, dsn)
	if err != nil {
		panic(err)
	}
	return db
}

// InitSqlxEnv 根据环境变量 初始化一个 sqlx.DB
func InitSqlxEnv(envEngineName, envDsnName string) *sqlx.DB {
	return InitSqlx(config.GetEnv(envEngineName), config.GetEnv(envDsnName))
}

// InitSqlxTag 根据Tag环境变量 初始化一个 sqlx.DB
func InitSqlxTag(tag string) *sqlx.DB {
	return InitSqlxEnv(strings.ToUpper(fmt.Sprintf("DB_%v_ENGINE", tag)),
		strings.ToUpper(fmt.Sprintf("DB_%v_DSN", tag)),
	)
}
