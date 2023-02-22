package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)
type Info struct{
	Name string `json:"name"`
	Email string `json:"email"`
	Message string `json:"message"`
}
func show(c *gin.Context) {
	reqBody,found:=c.Get("body")
	if !found{
		fmt.Println("Value not found")
	}
	fmt.Println("Requset body is:",reqBody)
}
func main() {
	fmt.Println("Working with form data using GIN")
	r := gin.Default()
	// attaching a get request to get the requst in the form of 
	// r.GET("/form", show)
	r.POST("/post", func(c *gin.Context) {
		var i Info
		name := c.PostForm("name")
		email := c.PostForm("email")
		message := c.PostForm("message")
		fmt.Println("Name:", name, "Email:", email, "Message:", message)
		i.Name=name
		i.Email=email
		i.Message=message
		// c.JSON(http.StatusOK, Info{
        //     Name: name,
        //     Email: email,
        //     Message: message,
        // })
		c.JSON(http.StatusOK,gin.H{
			"status": "success",
			"messaage": "Form submitted successfully",
		})
		c.Set("body",&i)
	})

	r.Run(":8000")
}
