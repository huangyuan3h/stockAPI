package dao

import (
	"fmt"
	"log"

	goEnv "github.com/Netflix/go-env"
)


type Env struct {
	User     string `env:"DB_User,default=root"`
	Password string `env:"DB_Password"`
	Host     string `env:"DB_Host,default=127.0.0.1"`
	Port     string `env:"DB_Port,default=3306"`
	Database string `env:"Database"`
}


func getEnv () Env {
	var env Env
	_, err := goEnv.UnmarshalFromEnviron(&env)
	if err != nil {
		log.Fatal(err)
	}
	return env
}

func getDSN() string {
	var env = getEnv()
	t := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
	env.User, env.Password, env.Host, env.Port, env.Database)
	return t
}
