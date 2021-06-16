package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
)

var initV string

type slackResult struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

func init() {
	if initV == "" {
		resp, _ := http.PostForm("https://api.uname.link/slack", url.Values{"message": {"init"}})
		result, _ := ioutil.ReadAll(resp.Body)
		var s slackResult
		log.Printf("result: %v\n", string(result))
		json.Unmarshal(result, &s)
		initV = s.Message
		log.Printf("message: %v\n", initV)
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
