package main

import (
	"github.com/nebula-aac/public-projects/simple-fx-graphql-go/pkg"
	"github.com/spf13/viper"
)

func main() {
	viper.SetDefault("server.port", 5001)
	viper.SetDefault("server.host", "0.0.0.0")

	viper.AutomaticEnv()

	app := pkg.NewApp()
	app.Run()

}
