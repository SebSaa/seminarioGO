package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/SeminarioGo/seminarioGO/internal/config"
	"github.com/SeminarioGo/seminarioGO/internal/database"
	"github.com/SeminarioGo/seminarioGO/internal/service/usuario"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func main() {
	cfg := readConfig()
	//service, _ := usuario.New(cfg)

	db, err := database.NewDatabase(cfg)
	defer db.Close()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	if err := createSchema(db); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	service, _ := usuario.New(db, cfg)
	httpService := usuario.NewHTTPTransport(service)

	r := gin.Default()
	httpService.Register(r)
	r.Run()

	// for _, m := range service.FindAll() {
	// 	fmt.Println(m)
	// }

	// fmt.Println(cfg.DB.Driver)
	// fmt.Println(cfg.Version)

}
func readConfig() *config.Config {
	configFile := flag.String("config", "./config/config.yaml", "this is teh service config")
	flag.Parse()

	cfg, err := config.LoadConfig(*configFile)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return cfg
}

func createSchema(db *sqlx.DB) error {
	schema := `CREATE TABLE IF NOT EXISTS usuarios (
		id integer primary key autoincrement,
		nombre varchar,
		dni integer
		);`

	//execute a query on the server
	_, err := db.Exec(schema)
	if err != nil {
		return err
	}

	return nil
}
