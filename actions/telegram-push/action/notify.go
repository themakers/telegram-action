package action

import (
	"bytes"
	"context"
	_ "embed"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"text/template"
)

//go:embed templates/push.md
var template_push_md string

type notifyOptions struct {
	BotToken string
	ChatID   string

	Repository string
	Branch     string
	CommitSHA  string
	Actor      string
}

func notify(ctx context.Context, opt notifyOptions) {
	var messageBuffer bytes.Buffer

	if tmpl, err := template.New("message").Parse(template_push_md); err != nil {
		panic(err)
	} else if err := tmpl.Execute(&messageBuffer, opt); err != nil {
		panic(err)
	}

	if body, err := json.MarshalIndent(map[string]any{
		"chat_id":    opt.ChatID,
		"text":       messageBuffer.String(),
		"parse_mode": "Markdown",
	}, "", "  "); err != nil {
		panic(err)
	} else {

		log.Println("request:", string(body))

		if req, err := http.NewRequestWithContext(ctx,
			"POST", fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", opt.BotToken), bytes.NewBuffer(body),
		); err != nil {
			panic(err)
		} else {
			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{}

			if resp, err := client.Do(req); err != nil {
				panic(err)
			} else {
				defer resp.Body.Close()

				if data, err := io.ReadAll(resp.Body); err != nil {
					panic(err)
				} else {
					log.Println("response:", resp.StatusCode, string(data))
				}

				if resp.StatusCode != http.StatusOK {
					panic(fmt.Sprintf("Request failed with status: %s", resp.Status))
				}

				log.Println("OK")
			}
		}
	}
}
