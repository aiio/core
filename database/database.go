package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

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

func MigrateAndComment(db *gorm.DB, comment string, model interface{}) {
	if !db.Migrator().HasTable(model) {
		err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8 comment '"+comment+"'").
			Migrator().CreateTable(model)
		if err != nil {
			panic(err)
		}
	}
}
