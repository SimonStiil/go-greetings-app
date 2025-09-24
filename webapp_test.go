package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETGreeting(t *testing.T) {
	t.Run("returns Pepper's score", func(t *testing.T) {
		count := getCount() + 1
		request, _ := http.NewRequest(http.MethodGet, "/greeting", nil)
		response := httptest.NewRecorder()

		greetingController(response, request)
		var greetingReply Greeting
		err := json.Unmarshal(response.Body.Bytes(), &greetingReply)
		if err != nil {
			t.Error(err)
		}
		greetinWanted := Greeting{count, "Hello, World!"}

		if greetingReply.Id != greetinWanted.Id {
			t.Errorf(".id got %v, want %v", greetingReply.Id, greetinWanted.Id)
		}
		if greetingReply.Content != greetinWanted.Content {
			t.Errorf(".content got %v, want %v", greetingReply.Content, greetinWanted.Content)
		}
	})
}
