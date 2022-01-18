package main

import (
	"log"
	"net/http"

	"github.com/kataras/iris/v12"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dbInstance, dbInstanceErr := gorm.Open(mysql.New(mysql.Config{
		DSN:                       "root:123456@tcp(mariadb:3306)/default?charset=utf8&parseTime=True", // data source name, refer https://github.com/go-sql-driver/mysql#dsn-data-source-name
		DefaultStringSize:         256,                                                                 // add default size for string fields, by default, will use db type `longtext` for fields without size, not a primary key, no index defined and don't have default values
		DisableDatetimePrecision:  true,                                                                // disable datetime precision support, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,                                                                // drop & create index when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,                                                                // use change when rename column, rename rename not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false,                                                               // smart configure based on used version
	}), &gorm.Config{})
	if dbInstanceErr != nil {
		log.Fatal(dbInstanceErr.Error())
	}

	app := iris.New()
	app.Get("/status", func(ctx iris.Context) {
		ctx.JSON(iris.Map{
			"code": http.StatusOK,
			"data": iris.Map{
				"RowsAffected": dbInstance.RowsAffected,
			},
		})
	})

	app.Get("/query/{query:string}", func(ctx iris.Context) {
		query := ctx.Params().GetString("query")
		// dbInstance.

		// // Raw SQL
		// rows, err := db.Raw("select name, age, email from users where name = ?", "jinzhu").Rows()
		// defer rows.Close()
		// for rows.Next() {
		// 	rows.Scan(&name, &age, &email)

		// 	// do something
		// }

		ctx.JSON(iris.Map{
			"code": http.StatusOK,
			"data": query,
		})
	})

	err := app.Run(
		// Start the web server at localhost:5000
		iris.Addr(":5000"),
		// skip err server closed when CTRL/CMD+C pressed:
		iris.WithoutServerError(iris.ErrServerClosed),
		// enables faster json serialization and more:
		iris.WithOptimizations,
	)

	if err != nil {
		log.Fatal(err.Error())
	}
}
