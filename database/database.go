package database

import (
	"fmt"
	"strings"

	"github.com/aiio/core/config"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// InitConn 初始化一个gorm.DB
func InitConn(engine, dsn, tablePrefix string) *gorm.DB {
	var dialector gorm.Dialector
	switch engine {
	case "mysql":
		dialector = mysql.New(mysql.Config{
			DSN:               dsn,
			DefaultStringSize: 191,
		})
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

// InitConnEnv 根据环境变量  初始化一个gorm.DB
func InitConnEnv(envEngineName, envDsnName, envPrefixName string) *gorm.DB {
	return InitConnEnv(config.GetEnv(envEngineName), config.GetEnv(envDsnName), config.GetEnv(envPrefixName))
}

// InitConnTag 根据Tag环境变量  初始化一个gorm.DB
func InitConnTag(tag string) *gorm.DB {
	return InitConnEnv(fmt.Sprintf("MYSQL_%v_ENGINE", strings.ToUpper(tag)),
		fmt.Sprintf("MYSQL_%v_DSN", strings.ToUpper(tag)),
		fmt.Sprintf("MYSQL_%v_PREFIX", strings.ToUpper(tag)),
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
