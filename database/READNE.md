env

### Mysql
```go
package dao

import (
	"github.com/aiio/core/database"
)

func Demo() {
	// init DB for env tag
	// DB_DEMO_ENGINE=mysql
	// DB_DEMO_DSN=root:123456@(127.0.0.1:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local

	// use https://github.com/go-gorm/gorm
	db := database.InitGormTag("demo")

	// use https://github.com/jmoiron/sqlx
	db := database.InitSqlxTag("demo")

	// use https://github.com/upper/db
	db := database.InitUpperTag("demo")
}
```


### Mysql
```shell
DB_DEFAULT_ENGINE=mysql
DB_DEFAULT_DSN=root:123456@(127.0.0.1:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local
```
