package main_test

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"cloud-run-lifecycle"
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

func TestPing(t *testing.T) {
	router := main.SetupRouter()

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