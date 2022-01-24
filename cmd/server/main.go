package main

import (
	"os"

	"github.com/InOuttt/snr/insert/config"
	"github.com/InOuttt/snr/insert/controller"
	"github.com/InOuttt/snr/insert/repository"
	"github.com/InOuttt/snr/insert/service"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/vanng822/go-solr/solr"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			println("Panic: %v ", r)
		}
	}()

	// Setup Configuration
	godotenv.Load(".env")
	cfg := config.LoadDbConfig("MYSQL_")
	println("mysql dsn: %v", cfg.DSN)
	sqlConn, err := config.NewMysqlSession(cfg)
	if err != nil {
		panic("failed inisiate mysql DB")
	}
	defer sqlConn.Db.Close()

	cfg = config.LoadDbConfig("SOLR_")
	println("Solr dsn: %v", cfg.DSN)
	si, err := solr.NewSolrInterface(cfg.DSN, "")
	if err != nil {
		panic("failed connect to solr")
	}

	dirPath := os.Getenv("DIR_PATH")

	// service repo controller
	feedRepo := repository.NewFeedRepository(sqlConn, si)
	feedSvc := service.NewFeedService(feedRepo)
	feedCtrl := controller.NewFeedController(feedSvc, feedRepo)

	feedCtrl.Handle(dirPath)

}
