package app

import (
	"github.com/gin-gonic/gin"
	"github.com/go-gorp/gorp"
	"github.com/ian.song/golang-gin-api/db"
	"log"
)

type App struct{
	DbMap *gorp.DbMap
}

func (app *App) Init(){
	app.DbMap = db.InitDb()
}

func (app *App) AppendEnv(c *gin.Context){
	log.Print("AppendEnv is executed each request")
	c.Set("dbMap",app.DbMap)
}
