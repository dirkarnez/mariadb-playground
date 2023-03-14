package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	isRunInDocker bool = false
	databaseName string
)

func main() {
	flag.BoolVar(&isRunInDocker, "docker", false, "Is run in docker")
	flag.StringVar(&databaseName, "database", "", "Database Name")
	flag.Parse()

	

	if len(databaseName) < 1 {
		log.Fatal("Database name should be specified")
	}

	var DSN string
	if isRunInDocker {
		DSN = "root:@tcp(mariadb:3306)/?charset=utf8&parseTime=True"
	} else {
		DSN = "root:@tcp(localhost:3306)/?charset=utf8&parseTime=True"
	}

	db, err := sql.Open("mysql", DSN)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec(fmt.Sprintf("DROP DATABASE IF EXISTS %s", databaseName))
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(fmt.Sprintf("CREATE DATABASE %s", databaseName))
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(fmt.Sprintf("USE %s", databaseName))
	if err != nil {
		log.Fatal(err)
	}

	// dbInstance, dbInstanceErr := gorm.Open(mysql.New(mysql.Config{
	// 	DSN:                       "root:123456@tcp(mariadb:3306)/default?charset=utf8&parseTime=True", // data source name, refer https://github.com/go-sql-driver/mysql#dsn-data-source-name
	// 	DefaultStringSize:         256,                                                                 // add default size for string fields, by default, will use db type `longtext` for fields without size, not a primary key, no index defined and don't have default values
	// 	DisableDatetimePrecision:  true,                                                                // disable datetime precision support, which not supported before MySQL 5.6
	// 	DontSupportRenameIndex:    true,                                                                // drop & create index when rename index, rename index not supported before MySQL 5.7, MariaDB
	// 	DontSupportRenameColumn:   true,                                                                // use change when rename column, rename rename not supported before MySQL 8, MariaDB
	// 	SkipInitializeWithVersion: false,                                                               // smart configure based on used version
	// }), &gorm.Config{})
	// if dbInstanceErr != nil {
	// 	log.Fatal(dbInstanceErr.Error())
	// }

	dbInstance, dbInstanceErr := gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		DefaultStringSize:         256,   // add default size for string fields, by default, will use db type `longtext` for fields without size, not a primary key, no index defined and don't have default values
		DisableDatetimePrecision:  true,  // disable datetime precision support, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,  // drop & create index when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,  // use change when rename column, rename rename not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false, // smart configure based on used version
	}), &gorm.Config{})
	if dbInstanceErr != nil {
		log.Fatal(dbInstanceErr.Error())
	}

	if dbInstance != nil {
		fmt.Println("Connected!")
	}
	// app := iris.New()
	// app.Get("/status", func(ctx iris.Context) {
	// 	ctx.JSON(iris.Map{
	// 		"code": http.StatusOK,
	// 		"data": iris.Map{
	// 			"RowsAffected": dbInstance.RowsAffected,
	// 		},
	// 	})
	// })

	// app.Get("/query", func(ctx iris.Context) {
	// 	query := "INSERT INTO 'lab1_schema1'.'Employee' ('EmpNo', 'FirstName', 'LastName', 'PhoneNo', 'Email') VALUES ('1', 'Tom', 'Chan', '23456789', 'tomchan@gmail.com')" //ctx.Params().GetString("query")

	// 	var result [][]string
	// 	rows, _ := dbInstance.Raw(query).Rows()
	// 	defer rows.Close()

	// 	cols, _ := rows.Columns()
	// 	pointers := make([]interface{}, len(cols))
	// 	container := make([]string, len(cols))

	// 	for i, _ := range pointers {
	// 		pointers[i] = &container[i]
	// 	}
	// 	for rows.Next() {
	// 		rows.Scan(pointers...)
	// 		result = append(result, container)
	// 	}

	// 	ctx.JSON(iris.Map{
	// 		"code": http.StatusOK,
	// 		"data": result,
	// 	})
	// })

	// err := app.Run(
	// 	// Start the web server at localhost:5000
	// 	iris.Addr(":5000"),
	// 	// skip err server closed when CTRL/CMD+C pressed:
	// 	iris.WithoutServerError(iris.ErrServerClosed),
	// 	// enables faster json serialization and more:
	// 	iris.WithOptimizations,
	// )

	// if err != nil {
	// 	log.Fatal(err.Error())
	// }

	fmt.Println("done")
	fmt.Scanln()
}
