package main

import (
	"xixi-movie/config"
	"xixi-movie/router"

	"github.com/gin-gonic/gin"
)


func main(){
	config.LoadEnv()
	r:=gin.Default()
	router.Setup(r)

	r.Run(":8899")
}
