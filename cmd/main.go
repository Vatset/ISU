package main

import (
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"isu1.0/comands"
	"isu1.0/repository"
	"isu1.0/service"
	"log"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs:%s", err.Error())
	}
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: viper.GetString("db.password"), /*os.Getenv("DB_PASSWORD"),*/
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})

	if err != nil {
		log.Fatalf("failed to initialize db %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	cli := comands.NewCLi(services)
	cli.Help()

}
func initConfig() error {
	viper.AddConfigPath("configs") //path
	viper.SetConfigName("configs") //file name
	return viper.ReadInConfig()
}
