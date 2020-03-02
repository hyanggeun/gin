package db
import (
	"database/sql"
	"github.com/go-gorp/gorp"
	"github.com/ian.song/golang-gin-api/model"
	_ "github.com/lib/pq" // TODO: for sql.Open
)

func InitDb() *gorp.DbMap{
	db,err := sql.Open("postgres", "user=kakao_ent dbname=postgis_test sslmode=disable")
	if err !=nil{
		panic(err)
	}
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}
	dbmap.AddTableWithName(model.Article{}, "articles").SetKeys(true, "Id")
	dbmap.AddTableWithName(model.User{}, "users").SetKeys(true, "Id")
	dbmap.DropTables()
	err = dbmap.CreateTablesIfNotExists()
	if err !=nil{
		panic(err)
	}
	return dbmap
}
