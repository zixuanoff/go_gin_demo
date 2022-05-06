package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

type State struct {
	StartTime   string
	CurrentTime string
	Counts      int
	Version     string
	Owner       string
}

func Recv(c *gin.Context, user *User) {
	if err := c.ShouldBindJSON(user); err != nil {
		log.Fatal(err)
		c.JSON(200, gin.H{"errcode": 400, "description": "Post Data Err"})
	} else {
		c.JSON(http.StatusOK, user)
	}
}

func main() {
	//Logging to a file
	gin.DisableConsoleColor()
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	r := gin.Default()

	InitDB()

	state := State{
		StartTime:   time.Now().String(),
		CurrentTime: time.Now().String(),
		Counts:      0,
		Version:     "0.0",
		Owner:       "admin",
	}

	r.GET("/check", func(c *gin.Context) {
		state.Counts++
		state.CurrentTime = time.Now().String()
		c.JSON(http.StatusOK, state)
	})

	r.POST("/echo", func(c *gin.Context) {
		state.Counts++
		user := new(User)
		Recv(c, user)
	})

	r.POST("/insert", func(c *gin.Context) {
		state.Counts++
		user := new(User)
		Recv(c, user)
		InsertUser(*user)
	})

	//Port
	r.Run(":9090")
	defer DB.Close()
}
