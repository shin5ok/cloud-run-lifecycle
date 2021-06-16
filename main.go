package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
)

var initV string

func init() {
	if initV == "" {
		resp, _ := http.PostForm("https://api.uname.link/slack", url.Values{"message": {"init"}})
		result, _ := ioutil.ReadAll(resp.Body)
		initV = string(result)
	}
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		t := fmt.Sprintf("%s", sig)
		http.PostForm("https://api.uname.link/slack", url.Values{"message": {"sigterm:" + t}})
	}()
}

func main() {
	g := gin.Default()

	g.GET("/coldstart", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": initV})
	})

	g.Run()

}
