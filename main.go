package main

import (
	"context"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

type Notification struct {
	Title     string
	Subtitle  string
	Message   string
	URL       string
	AvatarURL string
}

func NewNotification(ghn *github.Notification) *Notification {
	msg := strings.ReplaceAll(ghn.GetSubject().GetTitle(), "[", `\[`)
	return &Notification{
		Title:     ghn.GetRepository().GetName(),
		Subtitle:  ghn.GetSubject().GetType(),
		Message:   msg,
		URL:       "https://github.com/notifications",
		AvatarURL: ghn.GetRepository().GetOwner().GetAvatarURL(),
	}
}

// Notify sends desktop notification.
func (n Notification) Notify() error {
	tn, err := exec.LookPath("terminal-notifier")
	if err != nil {
		return err
	}
	cmd := exec.Command(tn, "-title", n.Title, "-subtitle", n.Subtitle, "-message", n.Message, "-contentImage", n.AvatarURL, "-open", n.URL)

	return cmd.Run()
}

func main() {
	token := os.Getenv("GH_TOKEN")
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)
	listN, _, err := client.Activity.ListNotifications(ctx, nil)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	for _, n := range listN {
		nt := NewNotification(n)
		err := nt.Notify()
		if err != nil {
			log.Fatal(err)
		}
	}
}
