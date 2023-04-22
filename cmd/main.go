package main

import (
	"context"
	posCreditation "github.com/AlibekDalgat/pos-credition"
	"github.com/AlibekDalgat/pos-credition/pkg/handler"
	"github.com/AlibekDalgat/pos-credition/pkg/repository"
	"github.com/AlibekDalgat/pos-credition/pkg/service"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/subosito/gotenv"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := initConfig(); err != nil {
		logrus.Fatalf("configni inisyalislengende xata boldu: %s", err.Error())
	}
	if err := gotenv.Load(); err != nil {
		logrus.Fatalf("env almaşınağanların tolturğanda xata boldu: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Post:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBname:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		logrus.Fatalf("db'ni inisyalislendirgende xata boldu: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(posCreditation.Server)
	go func() {
		if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil && err != http.ErrServerClosed {
			logrus.Fatalf("htt server işletilgende xata boldu: %s", err.Error())
		}
	}()

	logrus.Println("Todo başlandı")
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Println("Todo tamamlandı")
	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("server tamamlanğan zamanda xata boldu %s", err.Error())
	}
	if err := db.Close(); err != nil {
		logrus.Errorf("db birleşme yapğan zamanda xata boldu", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
