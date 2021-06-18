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

	uuid "github.com/google/uuid"

	"github.com/gin-gonic/gin"
)

var initV string
var Hostname string
var UUID string
var slackAPI = "https://api.uname.link/slack"
var slackChannel string

type slackResult struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

func postForm(message string, slackChannel string) (result []byte, err error) {
	d := time.Now()
	newMessage := fmt.Sprintf("%s(%s:%s)", message, Hostname, d.String())
	if slackChannel == "" {
		slackChannel = "kawano"
	}
	resp, _ := http.PostForm(slackAPI, url.Values{"message": {newMessage}, "slack_channel": {slackChannel}})
	return ioutil.ReadAll(resp.Body)
}

func genUUID() (uuidString string) {
	uuidObj := uuid.New().String()
	return uuidObj
}

func init() {
	if initV == "" {
		UUID = genUUID()
		Hostname, _ = os.Hostname()
		slackChannel = os.Getenv("SLACK_CHANNEL")

		result, _ := postForm("init:"+UUID, slackChannel)
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
		t := fmt.Sprintf("%s:%s", sig, UUID)
		postForm("signal:"+t, slackChannel)
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
