package main
import (
     "fmt"
     "github.com/gin-gonic/gin"
    )
func main(){
	fmt.Println("hello world")
	r:=gin.default()
	r.GET("/",func(c *gin.context){
		c.JSON(200,gin.H{
			"messahe":"hello world",
		})
	})
	r.Run()
}