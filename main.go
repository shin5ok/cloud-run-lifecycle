package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	uuid "github.com/google/uuid"
	"google.golang.org/api/run/v1"

	"github.com/gin-gonic/gin"
)

var (
	initV        string
	Hostname     string
	UUID         string
	doNotify     = os.Getenv("DO_NOTIFY")
	slackChannel = os.Getenv("SLACK_CHANNEL")
	slackAPI     = os.Getenv("SLACK_API")
	appName      = os.Getenv("NAME")
)

type slackResult struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

func postForm(message string, slackChannel string) (result []byte, err error) {
	if slackAPI == "" {
		log.Println("slackAPI is empty...That's why nothing to do.")
		return []byte{}, err
	}
	d := time.Now()
	newMessage := fmt.Sprintf("%s(%s:%s)", message, Hostname, d.String())
	if slackChannel == "" {
		slackChannel = "kawano"
	}
	postData := map[string]string{"text": newMessage, "channel": slackChannel}
	postJsonData, _ := json.Marshal(postData)
	req, err := http.NewRequest("POST", slackAPI, bytes.NewReader(postJsonData))

	client := &http.Client{}
	client.Do(req)

	if err != nil {
		log.Println(err)
		return []byte{}, err
	}
	return postJsonData, nil
}

func genUUID() (uuidString string) {
	uuidObj := uuid.New().String()
	return uuidObj
}

func init() {
	if appName == "" {
		log.Fatal("'NAME' env variable is empty")
	}

	UUID = genUUID()
	start := time.Now()
	Hostname, _ = os.Hostname()

	result, _ := postForm(fmt.Sprintf("init: %s: %s", appName, UUID), slackChannel)
	log.Printf("result: %v\n", string(result))
	var s slackResult
	if err := json.NewDecoder(bytes.NewBuffer(result)).Decode(&s); err != nil {
		log.Println(err.Error())
	}
	initV = s.Message

	log.Printf("message: %v\n", initV)

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		finish := time.Now()
		difftime := finish.Sub(start)
		t := fmt.Sprintf("%s: (%s) %s:%s", appName, difftime, sig, UUID)
		if doNotify != "" {
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

	g.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	g.Run(":" + port)

}

func getMeta() map[string]string {
	r := &run.Service{
		Metadata: &run.ObjectMeta{},
	}
	annotations := r.Metadata.Annotations
	return annotations
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
