package telegram

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	tgbotapi "github.com/OvyFlash/telegram-bot-api"
)

func TestTelegram_SendToMessageThread(t *testing.T) {
	t.Parallel()

	requests := make(chan url.Values, 1)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/bottoken/getMe" {
			fmt.Fprint(w, `{"ok":true,"result":{"id":1,"is_bot":true,"first_name":"test","username":"test_bot"}}`)
			return
		}

		if err := r.ParseForm(); err != nil {
			t.Errorf("parse request form: %v", err)
		}
		requests <- r.Form
		fmt.Fprint(w, `{"ok":true,"result":{"message_id":1,"date":0,"chat":{"id":-1001,"type":"supergroup"}}}`)
	}))
	t.Cleanup(server.Close)

	client, err := tgbotapi.NewBotAPIWithOptions(
		"token",
		tgbotapi.WithAPIEndpoint(server.URL+"/bot%s/%s"),
	)
	if err != nil {
		t.Fatalf("create Telegram client: %v", err)
	}

	service := &Telegram{client: client}
	service.AddReceivers(-1001)
	service.SetMessageThreadID(42)

	if sendErr := service.Send(context.Background(), "subject", "message"); sendErr != nil {
		t.Fatalf("send Telegram message: %v", sendErr)
	}

	request := <-requests
	if got := request.Get("message_thread_id"); got != "42" {
		t.Fatalf("message_thread_id = %q, want %q", got, "42")
	}
}
