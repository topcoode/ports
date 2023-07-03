package main

import (
	"encoding/json"
	"fmt"
	"ports/postgres"

	"github.com/gin-gonic/gin"
)

type Userdata struct {
	Networkfunction string
}

func main() {
	router := *gin.Default()
	router.POST("/DataFrmClient2", DataFrmClient2)
	router.Run(":8082")
}
func DataFrmClient2(c *gin.Context) {
	fmt.Println("the value of c:", c)
	var userdata Userdata
	err := c.ShouldBind(userdata)
	if err != nil {
		fmt.Println("error in binding", err)
	}

	db, err := postgres.Postgresconnection()
	sytx := `INSERT INTO client3 VALUES($1)`
	connection, err := db.Exec(sytx, userdata.Networkfunction)
	if err != nil {
		fmt.Println("error in connection", err)
	}
	fmt.Println(connection)
	data, err := json.Marshal(userdata)
	if err != nil {
		fmt.Println("problem in marshalling", err)
	}
	fmt.Println(data)

}
