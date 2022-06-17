package database

import (
	"fmt"
	"strings"

	"github.com/aiio/core/config"
	"gorm.io/driver/clickhouse"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// InitGorm 初始化一个gorm.DB
func InitGorm(engine, dsn, tablePrefix string) *gorm.DB {
	var dialector gorm.Dialector
	switch engine {
	case "mysql":
		dialector = mysql.Open(dsn)
	case "sqlserver":
		dialector = sqlserver.Open(dsn)
	case "postgres":
		dialector = postgres.Open(dsn)
	case "clickhouse":
		dialector = clickhouse.Open(dsn)
	case "sqlite":
		dialector = sqlite.Open(dsn)
	default:
		panic("engine error")
	}
	Db, err := gorm.Open(dialector, &gorm.Config{
		SkipDefaultTransaction: true,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   tablePrefix,
			SingularTable: true,
		},
		PrepareStmt:                              true,
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		panic(fmt.Sprintf("InitSqlLite err:%v,dsn:%+v", err, dsn))
	}
	return Db
}

// InitGormEnv 根据环境变量  初始化一个gorm.DB
func InitGormEnv(envEngineName, envDsnName, envPrefixName string) *gorm.DB {
	return InitGormEnv(config.GetEnv(envEngineName), config.GetEnv(envDsnName), config.GetEnv(envPrefixName))
}

// InitGormTag 根据Tag环境变量  初始化一个gorm.DB
func InitGormTag(tag string) *gorm.DB {
	return InitGormEnv(strings.ToUpper(fmt.Sprintf("DB_%v_ENGINE", tag)),
		strings.ToUpper(fmt.Sprintf("DB_%v_DSN", tag)),
		strings.ToUpper(fmt.Sprintf("DB_%v_PREFIX", tag)),
	)
}

// MigrateAndComment 同步表结构
func MigrateAndComment(db *gorm.DB, comment string, model interface{}) {
	if !db.Migrator().HasTable(model) {
		err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8 comment '"+comment+"'").
			Migrator().CreateTable(model)
		if err != nil {
			panic(err)
		}
	}
}

// AutoMigrateAndComment 自动同步表结构
func AutoMigrateAndComment(db *gorm.DB, comment string, model interface{}) {
	if !db.Migrator().HasTable(model) {
		err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8 comment '"+comment+"'").
			Migrator().CreateTable(model)
		if err != nil {
			panic(err)
		}
	} else {
		err := db.AutoMigrate(model)
		if err != nil {
			panic(err)
		}
	}
}
