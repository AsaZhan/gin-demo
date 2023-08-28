package impl

import (
	"database/sql"
	"fmt"
	"gin-demo/server/app"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func ConnectToDb(cfg app.AppConfig) (*sql.DB, error) {
	dsc := cfg.DataSourceConfig
	datasourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", dsc.Username, dsc.Password, dsc.IP, dsc.Port, dsc.Database)
	log.Println(datasourceName)
	if db, err := sql.Open("mysql", datasourceName); err != nil {
		log.Fatal(err)
		return nil, err
	} else if err := db.Ping(); err != nil {
		log.Fatal(err)
		return nil, err
	} else {
		log.Println("connected to database!")
		return db, nil
	}
}
