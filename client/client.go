package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Userdata struct {
	Networkfunction string
}

func main() {
	router := gin.Default()
	router.POST("/PushingIntoServer", Client)
	router.Run(":8080")
}
func Client(c *gin.Context) {
	var userdata Userdata
	userdata.Networkfunction = "policycontrolfunction"
	Userdata, err := json.Marshal(userdata)
	if err != nil {
		fmt.Println("error on marshalling..........", err)
	}
	fmt.Println("userdata------->", Userdata)
	req, err := http.NewRequest("POST", "http://127.0.0.1:8081/StoringClientData", bytes.NewBuffer(Userdata))
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
