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
	"time"

	"github.com/gin-gonic/gin"
)

var initV string
var Hostname string

type slackResult struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

func postForm(message string) (result []byte, err error) {
	d := time.Now()
	newmessage := fmt.Sprintf("%s(%s:%s)", message, Hostname, d.String())
	resp, _ := http.PostForm("https://api.uname.link/slack", url.Values{"message": {newmessage}})
	// result, err := ioutil.ReadAll(resp.Body)
	return ioutil.ReadAll(resp.Body)
}

func init() {
	if initV == "" {
		Hostname, _ = os.Hostname()
		result, _ := postForm("init")
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
		postForm("signal:" + t)
	}()
}

func main() {
	g := gin.Default()

	g.GET("/coldstart", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": initV})
	})

	g.GET("/envall", func(c *gin.Context) {
		c.JSON(200, gin.H{"env": os.Environ()})
	})

	g.Run()

}
