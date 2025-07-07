package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

// TestMainはこのパッケージの他のテストよりも先に実行されます。
// mainパッケージのinit()関数で必要とされる環境変数を設定するためにこれを使用します。
func TestMain(m *testing.M) {
	// main.goのinit()で必要な環境変数を設定します
	os.Setenv("NAME", "test-app-from-testmain")
	os.Setenv("SLACK_API", "") // 実際のスラック通知を防ぎます

	// すべてのテストを実行します
	exitCode := m.Run()

	// テスト実行と同じコードで終了します
	os.Exit(exitCode)
}


// setupRouterはテスト用にルーターをセットアップします。
// main()にあるものの一部を抜き出したもので、実際のサーバーは起動しません。
func setupRouter() *gin.Engine {
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
	return g
}

func TestPing(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	expectedBody := "{}"
	if w.Body.String() != expectedBody {
		t.Errorf("Expected body '%s', got '%s'", expectedBody, w.Body.String())
	}
}