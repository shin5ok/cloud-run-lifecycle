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
var slackDefaultAPI = "https://api.uname.link/slack"

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
	var appName = os.Getenv("NAME")
	if appName == "" {
		log.Fatal("'NAME' env variable is empty")
	}

	UUID = genUUID()
	start := time.Now()
	slackChannel := os.Getenv("SLACK_CHANNEL")
	slackNotify := os.Getenv("SLACK_NOTIFY")
	if slackAPI := os.Getenv("SLACK_API"); slackAPI == "" {
		slackAPI = slackDefaultAPI
	}
	Hostname, _ = os.Hostname()

	result, _ := postForm(fmt.Sprintf("init: %s: %s", appName, UUID), slackChannel)
	log.Printf("result: %v\n", string(result))
	var s slackResult
	json.Unmarshal(result, &s)
	initV = s.Message

	log.Printf("message: %v\n", initV)

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		finish := time.Now()
		difftime := finish.Sub(start)
		t := fmt.Sprintf("%s: (%s) %s:%s", appName, difftime, sig, UUID)
		if slackNotify {
			postForm(t, slackChannel)
		}
		log.Println(t)
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

	g.GET("/fuka", fuka)

	g.Run()

}

func fuka(c *gin.Context) {
	for _, v := range []int{10, 6, 23, 26, 39} {
		// log.Printf("fuka %d\n", v)
		slowfibo(v)
	}
}

func slowfibo(n int) int {
	if n < 2 {
		return n
	}
	return slowfibo(n-2) + slowfibo(n-1)
}
