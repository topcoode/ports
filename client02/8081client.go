package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"ports/postgres"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Userdata struct {
	Networkfunction string
}

func main() {
	router := gin.Default()
	router.POST("/StoringClientData", StoringClientData)
	router.POST("/SendingClientData", SendingClientData)

	router.Run(":8081")
}
func StoringClientData(c *gin.Context) {
	var userdata Userdata
	err := c.ShouldBindJSON(userdata)
	if err != nil {
		fmt.Println("error in binding", err)
	}

	db := postgres.Postgresconnection()
	sytx := `INSERT INTO client2 VALUES($1)`
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
func SendingClientData(c *gin.Context) {
	db := postgres.Postgresconnection()
	rows, err := db.Query("SELECT Networkfunction FROM client2")
	if err != nil {
		fmt.Println(err)
		//c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var users []Userdata
	for rows.Next() {
		var user Userdata
		err := rows.Scan(&user.Networkfunction)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		users = append(users, user)
	}
	Userdata, err := json.Marshal(users)
	if err != nil {
		fmt.Println("error on marshalling..........", err)

	}
	c.JSON(http.StatusOK, users)
	req, err := http.NewRequest("POST", "http://127.0.0.1:8082/DataFrmClient2", bytes.NewBuffer(Userdata))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Set the request content type to JSON
	req.Header.Set("Content-Type", "application/json")

	// Send the HTTP request
	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()
}
