package database

import (
	"fmt"
	"strings"

	"github.com/aiio/core/config"
	"github.com/upper/db/v4"
	"github.com/upper/db/v4/adapter/cockroachdb"
	"github.com/upper/db/v4/adapter/mongo"
	"github.com/upper/db/v4/adapter/mssql"
	"github.com/upper/db/v4/adapter/mysql"
	"github.com/upper/db/v4/adapter/postgresql"
	"github.com/upper/db/v4/adapter/ql"
	"github.com/upper/db/v4/adapter/sqlite"
)

// InitUpper 初始化一个 sqlx.DB
func InitUpper(engine, dsn string) db.Session {
	switch engine {
	case "mysql":
		url, err := mysql.ParseURL(dsn)
		if err != nil {
			panic(err)
		}
		sess, err := mysql.Open(url)
		if err != nil {
			panic(err)
		}
		return sess
	case "postgresql":
		url, err := postgresql.ParseURL(dsn)
		if err != nil {
			panic(err)
		}
		sess, err := postgresql.Open(url)
		if err != nil {
			panic(err)
		}
		return sess
	case "mssql":
		url, err := mssql.ParseURL(dsn)
		if err != nil {
			panic(err)
		}
		sess, err := mssql.Open(url)
		if err != nil {
			panic(err)
		}
		return sess
	case "cockroachdb":
		url, err := cockroachdb.ParseURL(dsn)
		if err != nil {
			panic(err)
		}
		sess, err := cockroachdb.Open(url)
		if err != nil {
			panic(err)
		}
		return sess
	case "mongodb":
		url, err := mongo.ParseURL(dsn)
		if err != nil {
			panic(err)
		}
		sess, err := mongo.Open(url)
		if err != nil {
			panic(err)
		}
		return sess
	case "sqlite":
		url, err := sqlite.ParseURL(dsn)
		if err != nil {
			panic(err)
		}
		sess, err := sqlite.Open(url)
		if err != nil {
			panic(err)
		}
		return sess
	case "ql":
		url, err := ql.ParseURL(dsn)
		if err != nil {
			panic(err)
		}
		sess, err := ql.Open(url)
		if err != nil {
			panic(err)
		}
		return sess
	default:
		panic("mysql engine err:" + engine)
	}
}

// InitUpperEnv 根据环境变量 初始化一个 db.Session
func InitUpperEnv(envEngineName, envDsnName string) db.Session {
	return InitUpper(config.GetEnv(envEngineName), config.GetEnv(envDsnName))
}

// InitUpperTag 根据Tag环境变量 初始化一个 db.Session
func InitUpperTag(tag string) db.Session {
	return InitUpperEnv(strings.ToUpper(fmt.Sprintf("DB_%v_ENGINE", tag)),
		strings.ToUpper(fmt.Sprintf("DB_%v_DSN", tag)),
	)
}
