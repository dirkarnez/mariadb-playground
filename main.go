package main

import (
	"github.com/kataras/iris/v12"
	"log"
	"net/http"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// cache = redis.Redis(host='redis', port=6379)
/**

def get_hit_count():
    retries = 5
    while True:
        try:
            return cache.incr('hits')
        except redis.exceptions.ConnectionError as exc:
            if retries == 0:
                raise exc
            retries -= 1
            time.sleep(0.5)


@app.route('/')
def hello():
    count = get_hit_count()
    return 'Hello World! I have been seen {} times.\n'.format(count)

**/

func main() {
	var DbInstance *gorm.DB = nil
	datasource.DbInstance, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       "root:123456@tcp(mariadb:3306)/default?charset=utf8&parseTime=True", // data source name, refer https://github.com/go-sql-driver/mysql#dsn-data-source-name
		DefaultStringSize:         256,                                                                 // add default size for string fields, by default, will use db type `longtext` for fields without size, not a primary key, no index defined and don't have default values
		DisableDatetimePrecision:  true,                                                                // disable datetime precision support, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,                                                                // drop & create index when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,                                                                // use change when rename column, rename rename not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false,                                                               // smart configure based on used version
	}), &gorm.Config{})
	if err != nil {
		log.Fatal("db open err")
	} else {
		fmt.Println("Connected to the database")
	}
	
	app := iris.New()
	app.Get("/", func(ctx iris.Context) {
		ctx.JSON(iris.Map{
			"code":  http.StatusOK,
			"data": "hello world!",
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
		log.Println(err.Error())
	}
}
