# use

## database

### env
```dotenv
# demo database env, is mysql engine
DB_DEMO_ENGINE=mysql
DB_DEMO_DSN=root:123456@(127.0.0.1:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local
# demo2 database env, is sqlite engine
DB_DEMO2_ENGINE=sqlite
DB_DEMO2_DSN=sqlite.db
```
### code
```go
package dao

import (
	"github.com/aiio/core/database"
)

func Demo() {
	// init DB for env tag
	
	// mysql 
	// DB_DEMO_ENGINE=mysql
	// DB_DEMO_DSN=root:123456@(127.0.0.1:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local
	
	// sqlite
	// DB_DEMO_ENGINE=sqlite
	// DB_DEMO_DSN=sqlite.db

	// use https://github.com/go-gorm/gorm
	db := database.InitGormTag("demo")

	// use https://github.com/jmoiron/sqlx
	db2 := database.InitSqlxTag("demo2")

	// use https://github.com/upper/db
	db3 := database.InitUpperTag("demo3")
}
```