package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETGreeting(t *testing.T) {
	t.Run("returns Pepper's score", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/greeting", nil)
		response := httptest.NewRecorder()

		greetingController(response, request)
		var greetingReply Greeting
		err := json.Unmarshal(response.Body.Bytes(), &greetingReply)
		if err != nil {
			t.Error(err)
		}
		greetinWanted := Greeting{0, "Hello, World!"}

		if greetingReply.Id != greetinWanted.Id {
			t.Errorf(".id got %q, want %q", greetingReply.Id, greetinWanted.Id)
		}
		if greetingReply.Content != greetinWanted.Content {
			t.Errorf(".content got %q, want %q", greetingReply.Content, greetinWanted.Content)
		}
	})
}
