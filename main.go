package main

import (
	"os"
	"os/exec"
	"os/signal"

	"github.com/gin-gonic/gin"
)

func main() {

	go func() {
		r := gin.Default()
		r.GET("/", func(c *gin.Context) {
			c.Writer.Write([]byte("test"))
		})
		r.Run()
	}()

	chromePath := "/Applications/Google Chrome.app/Contents/MacOS/Google Chrome"
	cmd := exec.Command(chromePath, "--app=http://127.0.0.1:8080/")
	cmd.Start()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	select {
	case <-quit:
		cmd.Process.Kill()
	}

}
