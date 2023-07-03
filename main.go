package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"ports/postgres"

	"github.com/gin-gonic/gin"
)

type Userdata struct {
	networkfunction string
}

func main() {
	router := gin.Default()
	router.GET("/database", Database)
	router.Run(":8080")
}
func Database(c *gin.Context) {
	var users []Userdata
	db, err := postgres.Postgresconnection()
	fmt.Println("err in db connection", err)
	rows, err := db.Query("SELECT networkfunction FROM client2;")
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()
	fmt.Println("query data :", users)
	for rows.Next() {
		var user Userdata
		err := rows.Scan(&user.networkfunction)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			fmt.Println("error in scan", err)
			return
		}
		users = append(users, user)
	}
	Userdata, err := json.Marshal(users)
	if err != nil {
		fmt.Println("error on marshalling..........", err)
	}
	fmt.Println("userdata json:", Userdata)
}
