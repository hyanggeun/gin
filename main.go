package main
import (
"github.com/ian.song/golang-gin-api/app"
"github.com/ian.song/golang-gin-api/routes"
)

func main()  {
	app := &app.App{}
	app.Init()

	router := routes.Init(app)

	router.Run(":8080")
}