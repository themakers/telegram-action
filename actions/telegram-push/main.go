package main

import (
	"context"
	"flag"
	"github.com/themakers/telegram-action/actions"
	"github.com/themakers/telegram-action/actions/telegram-push/action"
	"log"
)

// https://github.com/sethvargo/go-githubactions
// https://full-stack.blend.com/how-we-write-github-actions-in-go.html
// https://docs.github.com/en/actions/writing-workflows/choosing-what-your-workflow-does/caching-dependencies-to-speed-up-workflows

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var (
		botTokenFlag = flag.String("bot-token", "", "")
		chatIDFlag   = flag.String("chat-id", "", "")
	)
	flag.Parse()

	log.Println("environment:", *botTokenFlag, *chatIDFlag, actions.DumpEnv())

	action.Invoke(ctx, action.Options{
		BotToken: *botTokenFlag,
		ChatID:   *chatIDFlag,
	})
}
