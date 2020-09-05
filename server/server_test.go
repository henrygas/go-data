package server

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandleGet(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/post/", handleRequest)

	writer := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/post/2", nil)
	mux.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}
	var post Post
	_ = json.Unmarshal(writer.Body.Bytes(), &post)
	if post.Id != 2 {
		t.Error("Cannot retrieve JSON post")
	}
}

func TestHandlePost(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/post/", handleRequest)

	writer := httptest.NewRecorder()
	postJson := strings.NewReader(`{"Content": "Hello World", "Author": "Henry"}`)
	request, _ := http.NewRequest("POST", "/post/", postJson)

	mux.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}
}
