package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"example/gin/db"

	_ "github.com/lib/pq"
)

type User struct {
	Username string
	Password string
}

func addUser(ctx *gin.Context) {
	body := User{}
	data, err := ctx.GetRawData()
	if err != nil {
		ctx.AbortWithStatusJSON(400, "User is not defined")
		return
	}
	err = json.Unmarshal(data, &body)
	if err != nil {
		ctx.AbortWithStatusJSON(400, "Bad Input")
		return
	}
	//use Exec whenever we want to insert update or delete
	//Doing Exec(query) will not use a prepared statement, so lesser TCP calls to the SQL server
	_, err = db.Db.Exec("insert into users(username,password) values ($1,$2)", body.Username, body.Password)
	if err != nil {
		fmt.Println(err)
		ctx.AbortWithStatusJSON(400, "Couldn't create the new user.")
	} else {
		ctx.JSON(http.StatusOK, "User is successfully created.")
	}

}

func main() {
	router := gin.Default()

	// fmt.Println(os.Environ())
	db.ConnectDatabase()
	db.CreateTableUser()
	router.POST("/user", addUser)
	router.Run()
}
